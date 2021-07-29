package golog

import (
	"context"
	"encoding/json"
	"io"
)

type key string

const loggerKey key = "logger"

// WithLogger creates a new logger and attaches it to the Context.
func WithLogger(ctx context.Context, w io.Writer, fields Fields) context.Context {
	l := &Logger{
		enc:    json.NewEncoder(w),
		fields: fields,
	}

	return context.WithValue(ctx, loggerKey, l)
}

// GetLogger extracts the logger from a context.
// It returns a new empty logger if ctx doesn't contain a logger.
func GetLogger(ctx context.Context) *Logger {
	l, ok := ctx.Value(loggerKey).(*Logger)
	if ok {
		return l
	}

	return &Logger{}
}

// Debug outputs a debug log message using the logger contained in ctx.
// If ctx doesn't contain a logger, it uses a new logger.
func Debug(ctx context.Context, msg string, req *HTTPRequest, fields Fields) {
	l := GetLogger(ctx)
	l.output(levelDebug, msg, req, fields)
}

// Info outputs an info log message using the logger contained in ctx.
// If ctx doesn't contain a logger, it uses a new logger.
func Info(ctx context.Context, msg string, req *HTTPRequest, fields Fields) {
	l := GetLogger(ctx)
	l.output(levelInfo, msg, req, fields)
}

// Notice outputs a notice log message using the logger contained in ctx.
// If ctx doesn't contain a logger, it uses a new logger.
func Notice(ctx context.Context, msg string, req *HTTPRequest, fields Fields) {
	l := GetLogger(ctx)
	l.output(levelNotice, msg, req, fields)
}

// Warning outputs a warning log message using the logger contained in ctx.
// If ctx doesn't contain a logger, it uses a new logger.
func Warning(ctx context.Context, msg string, req *HTTPRequest, fields Fields) {
	l := GetLogger(ctx)
	l.output(levelWarning, msg, req, fields)
}

// Error outputs an error log message using the logger contained in ctx.
// If ctx doesn't contain a logger, it uses a new logger.
func Error(ctx context.Context, msg string, req *HTTPRequest, fields Fields) {
	l := GetLogger(ctx)
	l.output(levelError, msg, req, fields)
}

// Critical outputs a critical log message using the logger contained in ctx.
// If ctx doesn't contain a logger, it uses a new logger.
func Critical(ctx context.Context, msg string, req *HTTPRequest, fields Fields) {
	l := GetLogger(ctx)
	l.output(levelCritical, msg, req, fields)
}

// Alert outputs an alert log message using the logger contained in ctx.
// If ctx doesn't contain a logger, it uses a new logger.
func Alert(ctx context.Context, msg string, req *HTTPRequest, fields Fields) {
	l := GetLogger(ctx)
	l.output(levelAlert, msg, req, fields)
}

// Emergency outputs an emergency log message using the logger contained in ctx.
// If ctx doesn't contain a logger, it uses a new logger.
func Emergency(ctx context.Context, msg string, req *HTTPRequest, fields Fields) {
	l := GetLogger(ctx)
	l.output(levelEmergency, msg, req, fields)
}
