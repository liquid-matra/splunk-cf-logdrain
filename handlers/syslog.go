package handlers

import (
	"fmt"
	"io"
	"log/syslog"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	v4syslog "github.com/leodido/go-syslog/v4"
	"github.com/leodido/go-syslog/v4/rfc5424"
)

type SyslogHandler struct {
	debug  bool
	token  string
	writer *syslog.Writer
	parser v4syslog.Machine
}

func NewSyslogHandler(token, syslogAddr string) (*SyslogHandler, error) {
	if token == "" {
		return nil, fmt.Errorf("missing TOKEN value")
	}
	handler := &SyslogHandler{}
	handler.token = token

	parser := rfc5424.NewParser()

	if os.Getenv("DEBUG") == "true" {
		handler.debug = true
	}
	writer, err := syslog.Dial("tcp", syslogAddr,
		syslog.LOG_WARNING|syslog.LOG_DAEMON, "splunk-cf-logdrain")
	if err != nil {
		return nil, fmt.Errorf("syslog: %w", err)
	}

	handler.writer = writer
	handler.parser = parser
	return handler, nil
}

func (h *SyslogHandler) Handler() echo.HandlerFunc {
	return func(c echo.Context) error {
		t := c.Param("token")
		if h.token != t {
			return c.String(http.StatusUnauthorized, "")
		}
		b, _ := io.ReadAll(c.Request().Body)
		_, err := h.parser.Parse(b)
		if err != nil {
			return err
		}
		_, _ = h.writer.Write(b)
		return c.String(http.StatusOK, "")
	}
}
