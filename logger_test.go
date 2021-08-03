package golog

import (
	"context"
	"os"
)

func ExampleSimple() {
	l := NewLogger(os.Stdout, nil)

	l.Debug("hello world", nil, nil)

	//Output:
	// {"logging.googleapis.com/sourceLocation":{"file":"/home/christoph/Dev/go/github.com/fiskaly/golog/logger_test.go","line":11},"message":"hello world","severity":"DEBUG"}
}

func ExampleAdvanced() {
	l := NewLogger(os.Stdout, Fields{
		"request_id": "12345",
		"dynamic":    func() string { return "interesting" },
	})

	l.Info("hello world", nil, nil)

	//Output:
	// {"dynamic":"interesting","logging.googleapis.com/sourceLocation":{"file":"/home/christoph/Dev/go/github.com/fiskaly/golog/logger_test.go","line":23},"message":"hello world","request_id":"12345","severity":"INFO"}
}

func ExampleContext() {
	ctx := WithLogger(context.Background(), os.Stdout, Fields{
		"test": 12345,
	})

	Notice(ctx, "hello world", nil, nil)

	//Output:
	// {"logging.googleapis.com/sourceLocation":{"file":"/home/christoph/Dev/go/github.com/fiskaly/golog/logger_test.go","line":34},"message":"hello world","severity":"NOTICE","test":12345}
}

func ExampleDefault() {
	// empty context
	Warning(context.Background(), "empty context", nil, nil)

	// it works even with a nil context
	Error(nil, "nil context", nil, nil)

	//Output:
	// {"logging.googleapis.com/sourceLocation":{"file":"/home/christoph/Dev/go/github.com/fiskaly/golog/logger_test.go","line":42},"message":"empty context","severity":"WARNING"}
	// {"logging.googleapis.com/sourceLocation":{"file":"/home/christoph/Dev/go/github.com/fiskaly/golog/logger_test.go","line":45},"message":"nil context","severity":"ERROR"}
}
