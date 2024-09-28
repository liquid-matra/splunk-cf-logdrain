package main

type Configuration struct {
	ListenPort     string
	EnvPrefix      string
	TransportUrl   string
	SyslogEndpoint string
	Token          string
}

func NewConfiguration() Configuration {
	return Configuration{
		ListenPort:     "8080",
		EnvPrefix:      "splunk-cf-logdrain",
		TransportUrl:   "",
		SyslogEndpoint: "localhost:5140",
		Token:          "",
	}
}
