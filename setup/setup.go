package setup

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/olivere/elastic/v7"

	log "github.com/sirupsen/logrus"
)

//State holds all the shared main components
type State struct {
	Log     *log.Logger     //Log is the logger for each golang container
	Elastic *elastic.Client //Elastic is the elasticsearch client
	Ctx     context.Context //Ctx is the elasticsearch context
}

//GetMainState sets up all main components for golang and returns a state struct that contains all the relevent info
func GetMainState(container string) *State {
	var s State

	s.Log = log.New()
	s.Log.Infof("starting %s service", container)
	s.Log.SetLevel(log.DebugLevel)

	s.Log.Info("Setting up elasticsearch...")
	host := "elasticsearch"

	_, err := net.LookupIP(host)
	if err != nil {
		s.Log.Warn("Unable to find elasticsearch through docker network, setting ip to localhost")
		host = "127.0.0.1"
	}

	s.Log.Info("Sleeping for 40 seconds to let elasticsearch container get started")
	time.Sleep(40 * time.Second)

	s.Ctx = context.Background()

	esURL := fmt.Sprintf("http://%s:%d",
		host, 9200)

	client, err := elastic.NewSimpleClient(elastic.SetURL(esURL))
	if err != nil {
		s.Log.Error("Failed to connect to elasticsearch")
	}

	_, _, err = client.Ping(esURL).Do(s.Ctx)
	if err != nil {
		s.Log.Error("Unable to ping elasticsearch client")
	}
	s.Elastic = client

	s.Log.Info("Success")

	return &s
}
