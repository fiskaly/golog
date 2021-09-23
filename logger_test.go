package golog

import (
	"context"
	"os"
)

func ExampleSimple() {
	l := NewLogger(os.Stdout, nil)

	l.Debug("hello world", nil, nil)

	//Output:
	// {"logging.googleapis.com/sourceLocation":{"file":"logger_test.go","line":11},"message":"hello world","severity":"DEBUG"}
}

func ExampleAdvanced() {
	l := NewLogger(os.Stdout, Fields{
		"request_id": "12345",
		"dynamic":    func() interface{} { return "interesting" },
		"number":     func() interface{} { return 1234 },
	})

	l.Info("hello world", nil, nil)

	//Output:
	// {"dynamic":"interesting","logging.googleapis.com/sourceLocation":{"file":"logger_test.go","line":24},"message":"hello world","number":1234,"request_id":"12345","severity":"INFO"}
}

func ExampleContext() {
	ctx := WithLogger(context.Background(), os.Stdout, Fields{
		"test": 12345,
	})

	Notice(ctx, "hello world", nil, nil)

	//Output:
	// {"logging.googleapis.com/sourceLocation":{"file":"logger_test.go","line":35},"message":"hello world","severity":"NOTICE","test":12345}
}

func ExampleDefault() {
	// empty context
	Warning(context.Background(), "empty context", nil, nil)

	// it works even with a nil context
	Error(nil, "nil context", nil, nil)

	//Output:
	// {"logging.googleapis.com/sourceLocation":{"file":"logger_test.go","line":43},"message":"empty context","severity":"WARNING"}
	// {"logging.googleapis.com/sourceLocation":{"file":"logger_test.go","line":46},"message":"nil context","severity":"ERROR"}
}

func ExampleAddFields() {
	ctx := WithLogger(context.Background(), os.Stdout, nil)

	Critical(ctx, "hello world", nil, nil)

	AddFields(ctx, Fields{
		"test": 12345,
	})

	Alert(ctx, "hello world", nil, nil)

	//Output:
	// {"logging.googleapis.com/sourceLocation":{"file":"logger_test.go","line":56},"message":"hello world","severity":"CRITICAL"}
	// {"logging.googleapis.com/sourceLocation":{"file":"logger_test.go","line":62},"message":"hello world","severity":"ALERT","test":12345}
}
