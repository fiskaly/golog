package golog

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"runtime"
)

type entry struct {
	fields      Fields
	severity    level
	message     string
	httpRequest *HTTPRequest
	location    location
}

func newEntry(severity level, message string, req *HTTPRequest, fields Fields) entry {
	if fields == nil {
		fields = Fields{}
	}

	_, file, line, _ := runtime.Caller(3)

	file = filepath.Base(file)

	return entry{
		fields:   fields,
		severity: severity,
		message:  message,
		location: location{
			File: file,
			Line: line,
		},
		httpRequest: req,
	}
}

func (e entry) MarshalJSON() ([]byte, error) {
	e.fields["severity"] = e.severity
	e.fields["message"] = e.message
	e.fields["logging.googleapis.com/sourceLocation"] = e.location
	if e.httpRequest != nil {
		e.fields["httpRequest"] = e.httpRequest
	}

	return json.Marshal(e.fields)
}

// level describes the severity level of a log entry.
type level string

const (
	// Debug or trace information.
	levelDebug = "DEBUG"

	// Routine information, such as ongoing status or performance.
	levelInfo = "INFO"

	// Normal but significant events, such as start up, shut down, or a configuration change.
	levelNotice = "NOTICE"

	// Warning events might cause problems.
	levelWarning = "WARNING"

	// Error events are likely to cause problems.
	levelError = "ERROR"

	// Critical events cause more severe problems or outages.
	levelCritical = "CRITICAL"

	// A person must take an action immediately.
	levelAlert = "ALERT"

	// One or more systems are unusable.
	levelEmergency = "EMERGENCY"
)

type location struct {
	File     string `json:"file,omitempty"`
	Line     int    `json:"line,omitempty"`
	Function string `json:"function,omitempty"`
}

// An HTTPRequest contains information about an HTTP request and the response.
type HTTPRequest struct {
	Method    string `json:"requestMethod,omitempty"`
	URL       string `json:"requestUrl,omitempty"`
	Status    int    `json:"status,omitempty"`
	UserAgent string `json:"userAgent,omitempty"`
	RemoteIP  string `json:"remoteIp,omitempty"`
	ServerIP  string `json:"serverIp,omitempty"`
	Referer   string `json:"referer,omitempty"`
	Latency   string `json:"latency,omitempty"`
}

// FromStdHTTPRequest extracts all data from http.Request to the custom HTTPRequest type.
func FromStdHTTPRequest(request *http.Request) *HTTPRequest {
	return &HTTPRequest{
		Method:    request.Method,
		URL:       request.URL.String(),
		UserAgent: request.UserAgent(),
		RemoteIP:  request.RemoteAddr,
		Referer:   request.Referer(),
	}
}

// Fields contains additional key:value pairs to add to the log output.
// Any type which can be marshaled to JSON can be used as a value.
// Additionally, a `func() string` is allowed for dynamic values.
type Fields map[string]interface{}
