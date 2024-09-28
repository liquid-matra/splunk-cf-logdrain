package main

import (
	"fmt"
	"os"
	"os/signal"
	"splunk-cf-logdrain/handlers"

	"github.com/labstack/echo/v4"

	"net/http"
	_ "net/http/pprof"
)

var commit = "deadbeaf"
var release = "v1.2.2"
var buildVersion = release + "-" + commit

func main() {
	e := make(chan *echo.Echo, 1)
	os.Exit(realMain(e))
}

func realMain(echoChan chan<- *echo.Echo) int {
	cfg := NewConfiguration()

	// Echo framework
	e := echo.New()

	// Middleware
	healthHandler := handlers.HealthHandler{}
	e.GET("/health", healthHandler.Handler())
	e.GET("/api/version", handlers.VersionHandler(buildVersion))

	syslogHandler, err := handlers.NewSyslogHandler(cfg.Token, cfg.SyslogEndpoint)
	if err != nil {
		fmt.Printf("syslogHandler: %v\n", err)
		return 8
	}
	e.POST("/syslog/drain/:token", syslogHandler.Handler())

	setupPprof()
	setupInterrupts()

	echoChan <- e
	exitCode := 0
	if err := e.Start(":" + cfg.ListenPort); err != nil {
		fmt.Printf("error: %v\n", err)
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

func setupPprof() {
	go func() {
		err := http.ListenAndServe("localhost:6060", nil)
		if err != nil {
		}
	}()
}
