package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"splunk-cf-logdrain/handlers"
	"testing"
)

type Resource struct {
	ResourceType        string                 `json:"resourceType"`
	ID                  string                 `json:"id"`
	ApplicationName     string                 `json:"applicationName,omitempty"`
	EventID             string                 `json:"eventId"`
	Category            string                 `json:"category,omitempty"`
	Component           string                 `json:"component,omitempty"`
	TransactionID       string                 `json:"transactionId"`
	ServiceName         string                 `json:"serviceName,omitempty"`
	ApplicationInstance string                 `json:"applicationInstance,omitempty"`
	ApplicationVersion  string                 `json:"applicationVersion,omitempty"`
	OriginatingUser     string                 `json:"originatingUser,omitempty"`
	ServerName          string                 `json:"serverName,omitempty"`
	LogTime             string                 `json:"logTime"`
	Severity            string                 `json:"severity"`
	TraceID             string                 `json:"traceId,omitempty"`
	SpanID              string                 `json:"spanId,omitempty"`
	LogData             LogData                `json:"logData"`
	Custom              json.RawMessage        `json:"custom,omitempty"`
	Meta                map[string]interface{} `json:"-"`
	Error               error                  `json:"-"`
}
type LogData struct {
	Message string `json:"message"`
}

type mockProducer struct {
	t *testing.T
	q chan Resource
}

func (m *mockProducer) SetMetrics() {
	//TODO implement me
	panic("implement me")
}

func (m *mockProducer) DeadLetter(_ Resource) error {
	return nil
}

func (m *mockProducer) Push(_ []byte) error {
	return nil
}

func (m *mockProducer) Start() (chan bool, error) {
	d := make(chan bool)
	return d, nil
}

func (m *mockProducer) Output() <-chan Resource {
	if m.q == nil {
		m.q = make(chan Resource)
	}
	return m.q
}

/*func setup(t *testing.T) (*echo.Echo, func()) {
	e := echo.New()
	syslogHandler, err := handlers.NewSyslogHandler("t0ken", &mockProducer{t: t})
	assert.Nilf(t, err, "Expected NewSyslogHandler() to succeed")
	ironHandler, err := handlers.NewIronIOHandler("t0ken", &mockProducer{t: t})
	assert.Nilf(t, err, "Expected NewIronIOHandler() to succeed")

	e.POST("/syslog/drain/:token", syslogHandler.Handler(nil))
	e.POST("/ironio/drain/:token", ironHandler.Handler(nil))

	return e, func() {
		_ = e.Close()
	}
}*/

func setup(t *testing.T) *httptest.Server {
	var payload = `Starting Application on 50676a99-dce0-418a-6b25-1e3d with PID 8 (/home/vcap/app/BOOT-INF/classes started by vcap in /home/vcap/app)`
	var appVersion = `1.0-f53a57a`
	var transactionID = `eea9f72c-09b6-4d56-905b-b518fc4dc5b7`

	var rawMessage = `<14>1 2018-09-07T15:39:21.132433+00:00 suite-phs.staging.msa-eustaging 7215cbaa-464d-4856-967c-fd839b0ff7b2 [APP/PROC/WEB/0] - - {"app":"msa-eustaging","val":{"message":"` + payload + `"},"ver":"` + appVersion + `","evt":null,"sev":"INFO","cmp":"CPH","trns":"` + transactionID + `","usr":null,"srv":"msa-eustaging.eu-west.philips-healthsuite.com","service":"msa","inst":"50676a99-dce0-418a-6b25-1e3d","cat":"Tracelog","time":"2018-09-07T15:39:21Z"}`
	body := bytes.NewBufferString(rawMessage)

	rec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/syslog/drain/t0ken", body)
	if err != nil {
		t.Fatal(err)
	}
	sysLogHandlerFunc := handlers.SyslogHandler("t0ken", "localhost:5555")
	sysLogHandlerFunc(rec, req)

	mux := httptest.NewServer(http.HandlerFunc(sysLogHandlerFunc))

	return mux
}

func TestInvalidToken(t *testing.T) {
	req, _ := http.NewRequest(http.MethodPost, "/syslog/drain/t00ken", nil)
	rec := httptest.NewRecorder()

	sysLogHandlerFunc := handlers.SyslogHandler("", "localhost:5555")
	sysLogHandlerFunc(rec, req)

	res := rec.Result()
	if res.StatusCode != http.StatusUnauthorized {
		t.Errorf("expected status code %d, got %d", http.StatusUnauthorized, res.StatusCode)
	}
}

func TestSyslogHandler(t *testing.T) {
	var payload = `Starting Application on 50676a99-dce0-418a-6b25-1e3d with PID 8 (/home/vcap/app/BOOT-INF/classes started by vcap in /home/vcap/app)`
	var appVersion = `1.0-f53a57a`
	var transactionID = `eea9f72c-09b6-4d56-905b-b518fc4dc5b7`

	var rawMessage = `<14>1 2018-09-07T15:39:21.132433+00:00 suite-phs.staging.msa-eustaging 7215cbaa-464d-4856-967c-fd839b0ff7b2 [APP/PROC/WEB/0] - - {"app":"msa-eustaging","val":{"message":"` + payload + `"},"ver":"` + appVersion + `","evt":null,"sev":"INFO","cmp":"CPH","trns":"` + transactionID + `","usr":null,"srv":"msa-eustaging.eu-west.philips-healthsuite.com","service":"msa","inst":"50676a99-dce0-418a-6b25-1e3d","cat":"Tracelog","time":"2018-09-07T15:39:21Z"}`
	body := bytes.NewBufferString(rawMessage)

	rec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/syslog/drain/t0ken", body)
	if err != nil {
		t.Fatal(err)
	}
	sysLogHandlerFunc := handlers.SyslogHandler("t00ken", "localhost:5555")
	sysLogHandlerFunc(rec, req)

	// mux := httptest.NewServer(http.HandlerFunc(sysLogHandlerFunc))
	// mux.Close()
}

/*func TestSyslogHandler(t *testing.T) {
	e, teardown := setup(t)
	defer teardown()

	var payload = `Starting Application on 50676a99-dce0-418a-6b25-1e3d with PID 8 (/home/vcap/app/BOOT-INF/classes started by vcap in /home/vcap/app)`
	var appVersion = `1.0-f53a57a`
	var transactionID = `eea9f72c-09b6-4d56-905b-b518fc4dc5b7`

	var rawMessage = `<14>1 2018-09-07T15:39:21.132433+00:00 suite-phs.staging.msa-eustaging 7215cbaa-464d-4856-967c-fd839b0ff7b2 [APP/PROC/WEB/0] - - {"app":"msa-eustaging","val":{"message":"` + payload + `"},"ver":"` + appVersion + `","evt":null,"sev":"INFO","cmp":"CPH","trns":"` + transactionID + `","usr":null,"srv":"msa-eustaging.eu-west.philips-healthsuite.com","service":"msa","inst":"50676a99-dce0-418a-6b25-1e3d","cat":"Tracelog","time":"2018-09-07T15:39:21Z"}`
	body := bytes.NewBufferString(rawMessage)

	req := httptest.NewRequest(echo.POST, "/syslog/drain/t0ken", body)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}
*/
