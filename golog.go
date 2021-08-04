// Package golog implements structured logging for Google Cloud Platform as described in https://cloud.google.com/logging/docs/structured-logging.
// It allows attaching a logger to a context to print context-specific information to the log output.
package golog

import (
	"context"
	"io"
	"os"
)

type key string

const loggerKey key = "logger"

// WithLogger creates a new logger and attaches it to the Context.
func WithLogger(ctx context.Context, w io.Writer, fields Fields) context.Context {
	l := NewLogger(w, fields)

	return context.WithValue(ctx, loggerKey, l)
}

// GetLogger extracts the logger from a context.
// It returns a new empty logger if ctx doesn't contain a logger.
func GetLogger(ctx context.Context) *Logger {
	if ctx == nil {
		return NewLogger(os.Stdout, nil)
	}

	l, ok := ctx.Value(loggerKey).(*Logger)
	if !ok {
		return NewLogger(os.Stdout, nil)
	}

	return l
}

// Debug outputs a debug log message using the logger contained in ctx.
// If ctx doesn't contain a logger, it uses a new logger.
func Debug(ctx context.Context, msg string, req *HTTPRequest, fields Fields) {
	l := GetLogger(ctx)
	l.output(LevelDebug, msg, req, fields)
}

// Info outputs an info log message using the logger contained in ctx.
// If ctx doesn't contain a logger, it uses a new logger.
func Info(ctx context.Context, msg string, req *HTTPRequest, fields Fields) {
	l := GetLogger(ctx)
	l.output(LevelInfo, msg, req, fields)
}

// Notice outputs a notice log message using the logger contained in ctx.
// If ctx doesn't contain a logger, it uses a new logger.
func Notice(ctx context.Context, msg string, req *HTTPRequest, fields Fields) {
	l := GetLogger(ctx)
	l.output(LevelNotice, msg, req, fields)
}

// Warning outputs a warning log message using the logger contained in ctx.
// If ctx doesn't contain a logger, it uses a new logger.
func Warning(ctx context.Context, msg string, req *HTTPRequest, fields Fields) {
	l := GetLogger(ctx)
	l.output(LevelWarning, msg, req, fields)
}

// Error outputs an error log message using the logger contained in ctx.
// If ctx doesn't contain a logger, it uses a new logger.
func Error(ctx context.Context, msg string, req *HTTPRequest, fields Fields) {
	l := GetLogger(ctx)
	l.output(LevelError, msg, req, fields)
}

// Critical outputs a critical log message using the logger contained in ctx.
// If ctx doesn't contain a logger, it uses a new logger.
func Critical(ctx context.Context, msg string, req *HTTPRequest, fields Fields) {
	l := GetLogger(ctx)
	l.output(LevelCritical, msg, req, fields)
}

// Alert outputs an alert log message using the logger contained in ctx.
// If ctx doesn't contain a logger, it uses a new logger.
func Alert(ctx context.Context, msg string, req *HTTPRequest, fields Fields) {
	l := GetLogger(ctx)
	l.output(LevelAlert, msg, req, fields)
}

// Emergency outputs an emergency log message using the logger contained in ctx.
// If ctx doesn't contain a logger, it uses a new logger.
func Emergency(ctx context.Context, msg string, req *HTTPRequest, fields Fields) {
	l := GetLogger(ctx)
	l.output(LevelEmergency, msg, req, fields)
}
