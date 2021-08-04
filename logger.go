package golog

import (
	"encoding/json"
	"io"
)

// A Logger produces structured log output in the format defined by Google for GCP logs.
type Logger struct {
	encoder *json.Encoder
	fields  Fields
}

// NewLogger creates a new logger which outputs to the given `io.Writer`.
// It allows setting fields which are included in every output log entry.
func NewLogger(w io.Writer, fields Fields) *Logger {
	return &Logger{
		encoder: json.NewEncoder(w),
		fields:  fields,
	}
}

// output creates a new log entry and prints it to the logger's output.
func (l *Logger) output(severity LogLevel, msg string, req *HTTPRequest, fields Fields) {
	entry := newEntry(severity, msg, req, fields)

	for k, v := range l.fields {
		f, ok := v.(func() string)
		if ok {
			v = f()
		}

		entry.fields[k] = v
	}

	l.encoder.Encode(entry)
}

// Debug outputs a debug log message.
func (l *Logger) Debug(msg string, req *HTTPRequest, fields Fields) {
	l.output(LevelDebug, msg, req, fields)
}

// Info outputs an info log message.
func (l *Logger) Info(msg string, req *HTTPRequest, fields Fields) {
	l.output(LevelInfo, msg, req, fields)
}

// Notice outputs a notice log message.
func (l *Logger) Notice(msg string, req *HTTPRequest, fields Fields) {
	l.output(LevelNotice, msg, req, fields)
}

// Warning outputs a warning log message.
func (l *Logger) Warning(msg string, req *HTTPRequest, fields Fields) {
	l.output(LevelWarning, msg, req, fields)
}

// Error outputs an error log message.
func (l *Logger) Error(msg string, req *HTTPRequest, fields Fields) {
	l.output(LevelError, msg, req, fields)
}

// Critical outputs a critical log message.
func (l *Logger) Critical(msg string, req *HTTPRequest, fields Fields) {
	l.output(LevelCritical, msg, req, fields)
}

// Alert outputs an alert log message.
func (l *Logger) Alert(msg string, req *HTTPRequest, fields Fields) {
	l.output(LevelAlert, msg, req, fields)
}

// Emergency outputs an emergency log message.
func (l *Logger) Emergency(msg string, req *HTTPRequest, fields Fields) {
	l.output(LevelEmergency, msg, req, fields)
}
