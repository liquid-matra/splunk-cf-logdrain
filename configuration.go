package main

import "os"

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
		ListenPort:     "2020",
		FluentBitPort:  "8080",
		EnvPrefix:      "splunk-cf-logdrain",
		TransportUrl:   "",
		SyslogEndpoint: "localhost:5140",
		Token:          "",
	}
	cfg.ReadFromEnv()
	return cfg
}

func (cfg *Configuration) ReadFromEnv() {
	variableName := "LISTEN_PORT"
	content, isSet := os.LookupEnv(variableName)
	if isSet {
		cfg.ListenPort = content
	}

	variableName = "TOKEN"
	content, isSet = os.LookupEnv(variableName)
	if isSet {
		cfg.Token = content
	}

	variableName = "SYSLOG_ENDPOINT"
	content, isSet = os.LookupEnv(variableName)
	if isSet {
		cfg.SyslogEndpoint = content
	}

	variableName = "FLUENTBIT_PORT"
	content, isSet = os.LookupEnv(variableName)
	if isSet {
		cfg.SyslogEndpoint = content
	}

}
