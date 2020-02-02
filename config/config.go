// Copyright (c) 2019 Computing Infrastructure Research Centre (CIRC), McMaster
// University. All rights reserved.

// Package config contains the TOML configuration required by all modules.
package config

import (
	"errors"
	"io/ioutil"
	"time"

	"github.com/BurntSushi/toml"
)

// Config represents the values gathered from the "config.toml" file.
type Config struct {
	Common  CommonConfig  `toml:"common"`  // common section of config
	API     APIConfig     `toml:"api"`     // api section of config
	Webhook WebhookConfig `toml:"webhook"` // webhook section of config
}

// CommonConfig is the common section of "config.toml".
type CommonConfig struct {
	InitDelay Duration `toml:"init_delay"` // InitDelay is the microservice start-up delay

	AmqpHost string `toml:"amqp_host"` // AMQP hostname or IP address
	AmqpPort int    `toml:"amqp_port"` // AMQP port
	AmqpUser string `toml:"amqp_user"` // AMQP username
	AmqpPass string `toml:"amqp_pass"` // AMQP password

	ElasticsearchHost string `toml:"elasticsearch_host"` // Elasticsearch hostname or IP address
	ElasticsearchPort int    `toml:"elasticsearch_port"` // Elasticsearch port

	LoggingLevel string `toml:"logging_level"` // LoggingLevel is minimum logging level (debug,info,warn,error,fatal)

	SendGridToken  string `toml:"sendgrid_token"`  // SendGridToken is the SendGrid API token for email authentication
	SendGridEmail  string `toml:"sendgrid_email"`  // SendGridEmail is the address emails are sent from
	SendGridName   string `toml:"sendgrid_name"`   // SendGridName is the name of the account sending emails
	SendGridDomain string `toml:"sendgrid_domain"` // SendGridDomain is the domain used for password resets

	UserRegistration bool `toml:"user_registration"` // UserRegistration indicates of registration link is shown on login
	UserActivated    bool `toml:"user_activated"`    // UserActivated indicates if user is automatically activated after registration
}

// APIConfig is the api section of "config.toml".
type APIConfig struct {
	HTTPPort          int      `toml:"http_port"`          // HTTP port
	SecretLength      int      `toml:"secret_length"`      // length of authentication secret
	AuthExpiry        Duration `toml:"auth_expiry"`        // lifetime of authentication token
	AuthRenewAge      Duration `toml:"auth_renew_age"`     // age of token before renewal
	MiddlewareDisable bool     `toml:"middleware_disable"` // disable middleware
	AssetPath         string   `toml:"asset_path"`         // AssetPath is path to static assets
	HTTPSEnabled      bool     `toml:"https_enabled"`      // HTTPSEnabled is indicator if site is accessible over HTTPS (enables cookie hardening)
}

// WebhookConfig is the webhook section of "config.toml".
type WebhookConfig struct {
	HTTPPort int `toml:"http_port"` // HTTP port
}

// GetConfig accepts a file name and returns a Config structure of the parsed
// settings or an error.
func GetConfig(fileName string) (*Config, error) {
	// read file contents
	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		// cannot read file
		return nil, err
	}
	// attempt to parse file contents into config
	var config Config
	_, err = toml.Decode(string(contents), &config)
	if err != nil {
		// error encountered decoding file
		return nil, err
	}
	// validate configuration
	err = validateConfiguration(&config)
	if err != nil {
		// configuration not valid
		return nil, err
	}
	return &config, nil
}

// validateConfiguration validates the Configuration struct, returning an error
// if invalid. Modify this as fields are added/removed.
func validateConfiguration(config *Config) error {
	var zeroDuration Duration

	// common configuration
	if config.Common.InitDelay == zeroDuration {
		return errors.New("Config file [common]: init_delay must be longer than 0 seconds")
	}
	if config.Common.AmqpHost == "" {
		return errors.New("Config file [common]: amqp_host is not specified")
	}
	if config.Common.AmqpPort == 0 {
		return errors.New("Config file [common]: amqp_port is not specified")
	}
	if config.Common.AmqpUser == "" {
		return errors.New("Config file [common]: amqp_user is not specified")
	}
	if config.Common.AmqpPass == "" {
		return errors.New("Config file [common]: amqp_pass is not specified")
	}
	if config.Common.ElasticsearchHost == "" {
		return errors.New("Config file [common]: elasticsearch_host is not specified")
	}
	if config.Common.ElasticsearchPort == 0 {
		return errors.New("Config file [common]: elasticsearch_port is not specified")
	}
	if config.Common.LoggingLevel != "debug" &&
		config.Common.LoggingLevel != "info" &&
		config.Common.LoggingLevel != "warn" &&
		config.Common.LoggingLevel != "error" &&
		config.Common.LoggingLevel != "fatal" {
		return errors.New("Config file [common]: logging_level is not specified")
	}
	if config.Common.SendGridToken == "" {
		return errors.New("Config file [common]: sendgrid_token is not specified")
	}
	if config.Common.SendGridEmail == "" {
		return errors.New("Config file [common]: sendgrid_email is not specified")
	}
	if config.Common.SendGridName == "" {
		return errors.New("Config file [common]: sendgrid_name is not specified")
	}
	if config.Common.SendGridDomain == "" {
		return errors.New("Config file [common]: sendgrid_domain is not specified")
	}

	// api configuration
	if config.API.HTTPPort == 0 {
		return errors.New("Config file [api]: http_port is not specified")
	}
	if config.API.SecretLength == 0 {
		return errors.New("Config file [api]: secret_length is not specified")
	}
	if config.API.AuthExpiry == zeroDuration {
		return errors.New("Config file [api]: auth_expiry must be longer than 0 seconds")
	}
	if config.API.AuthRenewAge == zeroDuration {
		return errors.New("Config file [api]: auth_renew_age must be longer than 0 seconds")
	}
	if config.API.AssetPath == "" {
		return errors.New("Config file [api]: asset_path is not specified")
	}

	// webhook configuration
	if config.Webhook.HTTPPort == 0 {
		return errors.New("Config file [webhook]: http_port is not specified")
	}

	return nil
}

// Duration represents a setting that is a time duration. Do not modify.
type Duration struct {
	time.Duration
}

// UnmarshalText is used by the TOML parser to convert a duration string into a
// standard Golang parsed time duration. Do not modify.
func (d *Duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}
