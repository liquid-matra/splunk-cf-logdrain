package main

import (
	"fmt"
	"os"
)

type Configuration struct {
	ListenPort     string
	FluentBitPort  string
	EnvPrefix      string
	TransportUrl   string
	SyslogEndpoint string
	Token          string
}

func NewConfiguration() Configuration {
	cfg := Configuration{
		ListenPort:     "8080",
		FluentBitPort:  "2020",
		EnvPrefix:      "splunk-cf-logdrain",
		TransportUrl:   "",
		SyslogEndpoint: "localhost:5140",
		Token:          "",
	}
	cfg.ReadFromEnv()
	return cfg
}

func (cfg *Configuration) Print() {
	fmt.Printf("Configuration -> ListenPort: %s\tFluentbitport: %s\t\n", cfg.ListenPort, cfg.FluentBitPort)
}

func (cfg *Configuration) ReadFromEnv() {
	variableName := "LISTEN_PORT"
	content, isSet := os.LookupEnv(variableName)
	if isSet && os.Getenv(variableName) != "" {
		cfg.ListenPort = content
	}

	variableName = "TOKEN"
	content, isSet = os.LookupEnv(variableName)
	if isSet && os.Getenv(variableName) != "" {
		cfg.Token = content
	}

	variableName = "SYSLOG_ENDPOINT"
	content, isSet = os.LookupEnv(variableName)
	if isSet && os.Getenv(variableName) != "" {
		cfg.SyslogEndpoint = content
	}

	variableName = "FLUENTBIT_PORT"
	content, isSet = os.LookupEnv(variableName)
	if isSet && os.Getenv(variableName) != "" {
		cfg.FluentBitPort = content
	}

}
