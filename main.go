package main

import (
	"log/slog"
	"os"
	"os/signal"
	"splunk-cf-logdrain/handlers"

	"net/http"
	_ "net/http/pprof"
)

var commit = "deadbeaf"
var release = "v1.2.2"
var buildVersion = release + "-" + commit

func main() {
	e := make(chan *http.ServeMux, 1)
	os.Exit(realMain(e))
}

func realMain(serverChan chan<- *http.ServeMux) int {
	cfg := NewConfiguration()

	// http server and routes
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handlers.HealthHandler)
	mux.HandleFunc("GET /api/version", handlers.VersionHandler(buildVersion))
	mux.HandleFunc("POST /syslog/drain/"+cfg.Token, handlers.SyslogHandler(cfg.Token, cfg.SyslogEndpoint))

	setupInterrupts()

	serverChan <- mux
	exitCode := 0
	if err := http.ListenAndServe(":"+cfg.ListenPort, mux); err != nil {
		slog.Error("unable to run http server", "description", err.Error())
		exitCode = 6
	}
	return exitCode
}

func setupInterrupts() {
	// Setup a channel to receive a signal
	done := make(chan os.Signal, 1)

	// Notify this channel when a SIGINT is received
	signal.Notify(done, os.Interrupt)

	// Fire off a goroutine to loop until that channel receives a signal.
	// When a signal is received simply exit the program
	go func() {
		for range done {
			os.Exit(0)
		}
	}()
}
