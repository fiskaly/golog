package golog

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"runtime"
)

// {"httpRequest":{"requestMethod":"GET","requestUrl":"/something","remoteIp":"192.0.2.1:1234"},"logging.googleapis.com/sourceLocation":{"file":"entry_test.go","line":24},"message":"something","severity":"INFO"}

type entry struct {
	Fields      Fields       `json:"fields"`
	Severity    LogLevel     `json:"severity"`
	Message     string       `json:"message"`
	HTTPRequest *HTTPRequest `json:"httpRequest"`
	Location    location     `json:"logging.googleapis.com/sourceLocation"`
}

func newEntry(severity LogLevel, message string, req *HTTPRequest, fields Fields) entry {
	if fields == nil {
		fields = Fields{}
	}

	_, file, line, _ := runtime.Caller(3)

	file = filepath.Base(file)

	return entry{
		Fields:   fields,
		Severity: severity,
		Message:  message,
		Location: location{
			File: file,
			Line: line,
		},
		HTTPRequest: req,
	}
}

func (e entry) MarshalJSON() ([]byte, error) {
	if e.Fields == nil {
		e.Fields = make(Fields)
	}

	e.Fields["severity"] = e.Severity
	e.Fields["message"] = e.Message
	e.Fields["logging.googleapis.com/sourceLocation"] = e.Location
	if e.HTTPRequest != nil {
		e.Fields["httpRequest"] = e.HTTPRequest
	}

	return json.Marshal(e.Fields)
}

// LogLevel describes the severity LogLevel of a log entry.
type LogLevel string

const (
	// Debug or trace information.
	LevelDebug = "DEBUG"

	// Routine information, such as ongoing status or performance.
	LevelInfo = "INFO"

	// Normal but significant events, such as start up, shut down, or a configuration change.
	LevelNotice = "NOTICE"

	// Warning events might cause problems.
	LevelWarning = "WARNING"

	// Error events are likely to cause problems.
	LevelError = "ERROR"

	// Critical events cause more severe problems or outages.
	LevelCritical = "CRITICAL"

	// A person must take an action immediately.
	LevelAlert = "ALERT"

	// One or more systems are unusable.
	LevelEmergency = "EMERGENCY"
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
