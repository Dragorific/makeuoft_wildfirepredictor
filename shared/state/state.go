// Copyright (c) 2019 Computing Infrastructure Research Centre (CIRC), McMaster
// University. All rights reserved.

// Package state provides the application state for the microservices.
package state

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/dragorific/makeuoft_wildfirepredictor/config"
	"github.com/olivere/elastic/v7"
	"github.com/streadway/amqp"

	"github.com/sendgrid/sendgrid-go"
	log "github.com/sirupsen/logrus"
)

// State contains the application state.
type State struct {
	Hash       string           // Hash is the Git commit hash
	Log        *log.Logger      // Log is a structured event logger
	Config     *config.Config   // Config is the module configuration.
	AMQP1      *amqp.Channel    // AMQP1 is the RabbitMQ client state channel 1
	AMQP2      *amqp.Channel    // AMQP2 is the RabbitMQ client state channel 2
	AMQPReady1 chan bool        // AMQPReady1 is a channel that indicates when AMQP1 is ready for I/O
	AMQPReady2 chan bool        // AMQPReady2 is a channel that indicates when AMQP2 is ready for I/O
	Elastic    *elastic.Client  // Elastic is the Elasticsearch client state
	ElasticCtx context.Context  // ElasticCtx is the Elasticsearch client context
	AuthReady  bool             // AuthReady is if the "auth" index exists
	SendGrid   *sendgrid.Client // SendGrid is the email client
}

// Provision initalizes a State. It will initalize the module configuration,
// RabbitMQ, and Elasticsearch client states. It will return the initalized
// state or an error.
func Provision(serviceName string, versionHash string) (*State, error) {
	// empty state
	var s State

	// generate hash in state
	s.Hash = "build "
	if versionHash == "" {
		s.Hash += "unknown"
	} else {
		s.Hash += versionHash
	}

	// generate Log in State
	s.Log = log.New()

	s.Log.Infof("starting %s service, %s", serviceName, s.Hash)

	// generate Config in State
	if err := provisionConfig(&s); err != nil {
		s.Log.Error("failed to generate Config")
		return &s, err
	}

	// set minimum log levels
	if s.Config.Common.LoggingLevel == "debug" {
		s.Log.SetLevel(log.DebugLevel)
	}
	if s.Config.Common.LoggingLevel == "info" {
		s.Log.SetLevel(log.InfoLevel)
	}
	if s.Config.Common.LoggingLevel == "warn" {
		s.Log.SetLevel(log.WarnLevel)
	}
	if s.Config.Common.LoggingLevel == "error" {
		s.Log.SetLevel(log.ErrorLevel)
	}
	if s.Config.Common.LoggingLevel == "fatal" {
		s.Log.SetLevel(log.FatalLevel)
	}

	// validate and conditionally adjust RabbitMQ + Elasticsearch hostnames
	fixHostnames(&s)

	// sleep to allow networked containers to start
	delay := s.Config.Common.InitDelay.Duration
	s.Log.Infof("sleeping %s to allow RabbitMQ and Elasticsearch to start", delay)
	time.Sleep(delay)

	// generate AMQPReady1 and AMQPReady2 in state
	s.AMQPReady1 = make(chan bool, 1)
	s.AMQPReady2 = make(chan bool, 1)

	// generate AMQP1 + AMQP2 in State
	if err := provisionRabbitMq(&s); err != nil {
		s.Log.Error("failed to establish RabbitMQ connection")
		return &s, err
	}

	// generate Elastic in State
	if err := provisionElasticsearch(&s); err != nil {
		s.Log.Error("failed to establish Elasticsearch connection")
		return &s, err
	}

	// generate AuthReady in State
	if err := checkAuthReady(&s); err != nil {
		s.Log.Error("failed to check if Elasticsearch 'auth' index exists")
		return &s, err
	}

	// generate SendGrid in State
	s.SendGrid = sendgrid.NewSendClient(s.Config.Common.SendGridToken)

	s.Log.Infof("%s service ready", serviceName)
	return &s, nil
}

// provisionConfig generates the Config entry in State.
func provisionConfig(s *State) error {
	// generate Config
	s.Log.Info("fetching and parsing configuration file")
	config, err := config.GetConfig("config.toml")
	if err != nil {
		return err
	}
	s.Config = config
	return nil
}

// fixHostnames will validate that all Docker specific hostnames are being
// resolved. If the domain cannot be resolved, all of the hostnames will be
// swapped to loopback address of "127.0.0.1".
func fixHostnames(s *State) {
	// check if AmqpHost is valid
	s.Log.Info("attempting to resolve DNS for amqp_host")
	_, err := net.LookupIP(s.Config.Common.AmqpHost)
	if err != nil {
		// not running in Docker or invalid hostname, set to loopback
		s.Log.Warn("service not running in Docker or invalid amqp_host, setting to 127.0.0.1")
		s.Config.Common.AmqpHost = "127.0.0.1"
	}

	// check if ElasticsearchHost is valid
	s.Log.Info("attempting to resolve DNS for elasticsearch_host")
	_, err = net.LookupIP(s.Config.Common.ElasticsearchHost)
	if err != nil {
		// not running in Docker or invalid hostname, set to loopback
		s.Log.Warn("service not running in Docker or invalid elasticsearch_host, setting to 127.0.0.1")
		s.Config.Common.ElasticsearchHost = "127.0.0.1"
	}
}

// provisionRabbitMq generates the AMQP1 + AMQP2 entries in State. It also
// registers a helper to automatically recovery the AMQP connection on failure.
func provisionRabbitMq(s *State) error {
	// generate AMQP state
	s.Log.Info("initializing RabbitMQ client state")
	amqpURL := fmt.Sprintf("amqp://%s:%s@%s:%d/", s.Config.Common.AmqpUser, s.Config.Common.AmqpPass,
		s.Config.Common.AmqpHost, s.Config.Common.AmqpPort)

	// connect to AMQP broker
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return err
	}
	// create AMQP channel 1
	ch1, err := conn.Channel()
	if err != nil {
		return err
	}
	s.AMQP1 = ch1

	// create AMQP channel 2
	ch2, err := conn.Channel()
	if err != nil {
		return err
	}
	s.AMQP2 = ch2

	// automatically reprovision RabbitMQ if connection fails after successful
	// initial startup
	go func(s *State, conn *amqp.Connection) {
		// channel is triggered upon closure
		<-conn.NotifyClose(make(chan *amqp.Error))
		s.Log.Warn("RabbitMQ client has diconnected; reprovisioning ...")
		amqpRecovery(s, conn)
	}(s, conn)

	// AMQP connections 1 and 2 are ready for I/O
	s.AMQPReady1 <- true
	s.AMQPReady2 <- true
	return nil
}

// amqpRecovery is called when an active AMQP connection fails. This will retry
// the connection every 15 seconds until the connection is re-established.
func amqpRecovery(s *State, conn *amqp.Connection) {
	// prevent memory leaks
	if conn != nil {
		conn.Close()
	}
	// recursively call amqpRecovery() until there is no error
	err := provisionRabbitMq(s)
	if err != nil {
		s.Log.Warn("RabbitMQ client failed to reprovision; retrying in 15s ...")
		time.Sleep(15 * time.Second)
		amqpRecovery(s, conn)
		return
	}
	s.Log.Info("RabbitMQ client successfully reprovisioned")
	return
}

// provisionElasticsearch generates the Elastic and ElasticCtx entries in State.
func provisionElasticsearch(s *State) error {
	// generate ELasticsearch state
	s.Log.Info("initializing Elasticsearch client state")

	// generate context
	s.ElasticCtx = context.Background()

	esURL := fmt.Sprintf("http://%s:%d",
		s.Config.Common.ElasticsearchHost, s.Config.Common.ElasticsearchPort)

	// connect to Elasticsearch
	client, err := elastic.NewSimpleClient(elastic.SetURL(esURL))
	if err != nil {
		return err
	}
	// ping to ensure connectivity
	_, _, err = client.Ping(esURL).Do(s.ElasticCtx)
	if err != nil {
		return err
	}
	s.Elastic = client
	return nil
}

// checkAuthReady generates the AuthReady entry in State. It will check if the
// "auth" index is created in Elasticsearch, setting the flag to true or false.
func checkAuthReady(s *State) error {
	// don't use our elasticsearch library to prevent cyclic import (for
	// everything else please use it)
	s.Log.Info("checking if elasticsearch 'auth' index exists")
	exists, err := s.Elastic.IndexExists("auth").Do(s.ElasticCtx)
	if err != nil {
		return err
	}
	s.AuthReady = exists
	if !s.AuthReady {
		s.Log.Warn("elasticsearch 'auth' index does not exist, system not initialized")
	}
	return nil
}

// BindAMQP uses the AMQP connection established in State, sending AMQP messages
// on the output channel provided. It will declare the exchange, create a queue,
// and route messages with the provided key(s) to the queue. The consumerTag is
// used for tracking purposes (RabbitMQ web interface). If the exchange or queue
// cannot be declared, or if the messages cannot be routed or consumed, an error
// will be returned.
func BindAMQP(s *State, amqpCh *amqp.Channel, out chan amqp.Delivery, exchange, queue string, routingKeys []string) error {
	// declare exchange
	err := amqpCh.ExchangeDeclare(
		exchange, // name
		"topic",  // type
		true,     // durable
		false,    // auto deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		return err
	}
	// declare queue
	q, err := amqpCh.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}
	// bind all routing keys to queue (via exchange)
	for _, key := range routingKeys {
		err = amqpCh.QueueBind(
			q.Name,   // name
			key,      // routing key
			exchange, // exchange
			false,    // no wait
			nil,      // arguments
		)
		if err != nil {
			return err
		}
		s.Log.Infof("listening to %s messages on %s", key, exchange)
	}
	// generate channel to consume from
	ch, err := amqpCh.Consume(
		q.Name, // name
		"",     // consumer tag
		true,   // auto ack message
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // arguments
	)
	if err != nil {
		return err
	}
	// consume RabbitMQ and send deliveries to output channel
	go func(input <-chan amqp.Delivery, output chan amqp.Delivery) {
		for d := range input {
			output <- d
		}
	}(ch, out)

	// no error
	return nil
}
