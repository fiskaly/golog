package golog

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
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
func (l *Logger) output(severity level, msg string, req *HTTPRequest, fields Fields) {
	countLog()

	entry := newEntry(severity, msg, req, fields)

	if entry.fields == nil {
		entry.fields = make(Fields)
	}

	for k, v := range l.fields {
		f, ok := v.(func() interface{})
		if ok {
			v = f()
		}

		entry.fields[k] = v
	}

	err := l.encoder.Encode(entry)
	if err != nil {
		log.Println(err)
		fmt.Println(err)
	}
}

// AddFields adds new fields to the logger.
// Existing fields might be overwritten.
func (l *Logger) AddFields(newFields Fields) {
	if l.fields == nil {
		l.fields = make(Fields)
	}

	for k, v := range newFields {
		l.fields[k] = v
	}
}

// Debug outputs a debug log message.
func (l *Logger) Debug(msg string, req *HTTPRequest, fields Fields) {
	l.output(levelDebug, msg, req, fields)
}

// Info outputs an info log message.
func (l *Logger) Info(msg string, req *HTTPRequest, fields Fields) {
	l.output(levelInfo, msg, req, fields)
}

// Notice outputs a notice log message.
func (l *Logger) Notice(msg string, req *HTTPRequest, fields Fields) {
	l.output(levelNotice, msg, req, fields)
}

// Warning outputs a warning log message.
func (l *Logger) Warning(msg string, req *HTTPRequest, fields Fields) {
	l.output(levelWarning, msg, req, fields)
}

// Error outputs an error log message.
func (l *Logger) Error(msg string, req *HTTPRequest, fields Fields) {
	l.output(levelError, msg, req, fields)
}

// Critical outputs a critical log message.
func (l *Logger) Critical(msg string, req *HTTPRequest, fields Fields) {
	l.output(levelCritical, msg, req, fields)
}

// Alert outputs an alert log message.
func (l *Logger) Alert(msg string, req *HTTPRequest, fields Fields) {
	l.output(levelAlert, msg, req, fields)
}

// Emergency outputs an emergency log message.
func (l *Logger) Emergency(msg string, req *HTTPRequest, fields Fields) {
	l.output(levelEmergency, msg, req, fields)
}
