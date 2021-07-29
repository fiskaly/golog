package golog

import (
	"context"
	"os"
)

func ExampleSimple() {
	l := NewLogger(os.Stdout, nil)

	l.Warning("hello world", nil, nil)

	//Output:
	// {"logging.googleapis.com/sourceLocation":{"file":"/home/christoph/Dev/go/github.com/fiskaly/golog/logger_test.go","line":11},"message":"hello world","severity":"WARNING"}
}

func ExampleAdvanced() {
	l := NewLogger(os.Stdout, Fields{
		"request_id": "12345",
		"dynamic":    func() string { return "interesting" },
	})

	l.Critical("hello world", nil, nil)

	//Output:
	// {"dynamic":"interesting","logging.googleapis.com/sourceLocation":{"file":"/home/christoph/Dev/go/github.com/fiskaly/golog/logger_test.go","line":23},"message":"hello world","request_id":"12345","severity":"CRITICAL"}
}

func ExampleContext() {
	ctx := WithLogger(context.Background(), os.Stdout, nil)

	Debug(ctx, "hello world", nil, nil)

	//Output:
	// {"logging.googleapis.com/sourceLocation":{"file":"/home/christoph/Dev/go/github.com/fiskaly/golog/logger_test.go","line":32},"message":"hello world","severity":"DEBUG"}
}
