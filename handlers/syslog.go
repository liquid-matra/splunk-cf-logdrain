package handlers

import (
	"io"
	"log/slog"
	"log/syslog"
	"net/http"
	"os"

	v4syslog "github.com/leodido/go-syslog/v4"
	"github.com/leodido/go-syslog/v4/rfc5424"
)

type Syslog struct {
	debug  bool
	token  string
	writer *syslog.Writer
	parser v4syslog.Machine
}

func SyslogHandler(token, syslogAddr string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("please provide the correct token"))
			return
		}

		handler := &Syslog{}
		handler.token = token

		parser := rfc5424.NewParser()

		if os.Getenv("DEBUG") == "true" {
			handler.debug = true
		}
		writer, err := syslog.Dial("tcp", syslogAddr,
			syslog.LOG_WARNING|syslog.LOG_DAEMON, "splunk-cf-logdrain")
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			slog.Error("unable to contact syslog server", "host", syslogAddr)
			w.Write([]byte{})
			return
		}

		handler.writer = writer
		handler.parser = parser

		b, _ := io.ReadAll(r.Body)
		_, err = parser.Parse(b)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			slog.Error("unable to parse syslog message", "description", err.Error(), "syslog-message", r.Body)
			w.Write([]byte{})
			return
		}
		_, _ = writer.Write(b)
		w.WriteHeader(http.StatusOK)
	}
}
