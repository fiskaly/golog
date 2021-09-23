# golog
--
    import "github.com/fiskaly/golog"

Package golog implements structured logging for Google Cloud Platform as
described in https://cloud.google.com/logging/docs/structured-logging. It allows
attaching a logger to a context to print context-specific information to the log
output.

## Usage

#### func  AddFields

```go
func AddFields(ctx context.Context, newFields Fields)
```
AddFields adds new fields to the logger contained in ctx. Existing fields might
be overwritten.

#### func  Alert

```go
func Alert(ctx context.Context, msg string, req *HTTPRequest, fields Fields)
```
Alert outputs an alert log message using the logger contained in ctx. If ctx
doesn't contain a logger, it uses a new logger.

#### func  Critical

```go
func Critical(ctx context.Context, msg string, req *HTTPRequest, fields Fields)
```
Critical outputs a critical log message using the logger contained in ctx. If
ctx doesn't contain a logger, it uses a new logger.

#### func  Debug

```go
func Debug(ctx context.Context, msg string, req *HTTPRequest, fields Fields)
```
Debug outputs a debug log message using the logger contained in ctx. If ctx
doesn't contain a logger, it uses a new logger.

#### func  Emergency

```go
func Emergency(ctx context.Context, msg string, req *HTTPRequest, fields Fields)
```
Emergency outputs an emergency log message using the logger contained in ctx. If
ctx doesn't contain a logger, it uses a new logger.

#### func  Error

```go
func Error(ctx context.Context, msg string, req *HTTPRequest, fields Fields)
```
Error outputs an error log message using the logger contained in ctx. If ctx
doesn't contain a logger, it uses a new logger.

#### func  Info

```go
func Info(ctx context.Context, msg string, req *HTTPRequest, fields Fields)
```
Info outputs an info log message using the logger contained in ctx. If ctx
doesn't contain a logger, it uses a new logger.

#### func  Notice

```go
func Notice(ctx context.Context, msg string, req *HTTPRequest, fields Fields)
```
Notice outputs a notice log message using the logger contained in ctx. If ctx
doesn't contain a logger, it uses a new logger.

#### func  Warning

```go
func Warning(ctx context.Context, msg string, req *HTTPRequest, fields Fields)
```
Warning outputs a warning log message using the logger contained in ctx. If ctx
doesn't contain a logger, it uses a new logger.

#### func  WithLogger

```go
func WithLogger(ctx context.Context, w io.Writer, fields Fields) context.Context
```
WithLogger creates a new logger and attaches it to the Context.

#### type Fields

```go
type Fields map[string]interface{}
```

Fields contains additional key:value pairs to add to the log output. Any type
which can be marshaled to JSON can be used as a value. Additionally, a `func()
string` is allowed for dynamic values.

#### type HTTPRequest

```go
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
```

An HTTPRequest contains information about an HTTP request and the response.

#### func  FromStdHTTPRequest

```go
func FromStdHTTPRequest(request *http.Request) *HTTPRequest
```
FromStdHTTPRequest extracts all data from http.Request to the custom HTTPRequest
type.

#### type Logger

```go
type Logger struct {
}
```

A Logger produces structured log output in the format defined by Google for GCP
logs.

#### func  GetLogger

```go
func GetLogger(ctx context.Context) *Logger
```
GetLogger extracts the logger from a context. It returns a new empty logger if
ctx doesn't contain a logger.

#### func  NewLogger

```go
func NewLogger(w io.Writer, fields Fields) *Logger
```
NewLogger creates a new logger which outputs to the given `io.Writer`. It allows
setting fields which are included in every output log entry.

#### func (*Logger) AddFields

```go
func (l *Logger) AddFields(newFields Fields)
```
AddFields adds new fields to the logger. Existing fields might be overwritten.

#### func (*Logger) Alert

```go
func (l *Logger) Alert(msg string, req *HTTPRequest, fields Fields)
```
Alert outputs an alert log message.

#### func (*Logger) Critical

```go
func (l *Logger) Critical(msg string, req *HTTPRequest, fields Fields)
```
Critical outputs a critical log message.

#### func (*Logger) Debug

```go
func (l *Logger) Debug(msg string, req *HTTPRequest, fields Fields)
```
Debug outputs a debug log message.

#### func (*Logger) Emergency

```go
func (l *Logger) Emergency(msg string, req *HTTPRequest, fields Fields)
```
Emergency outputs an emergency log message.

#### func (*Logger) Error

```go
func (l *Logger) Error(msg string, req *HTTPRequest, fields Fields)
```
Error outputs an error log message.

#### func (*Logger) Info

```go
func (l *Logger) Info(msg string, req *HTTPRequest, fields Fields)
```
Info outputs an info log message.

#### func (*Logger) Notice

```go
func (l *Logger) Notice(msg string, req *HTTPRequest, fields Fields)
```
Notice outputs a notice log message.

#### func (*Logger) Warning

```go
func (l *Logger) Warning(msg string, req *HTTPRequest, fields Fields)
```
Warning outputs a warning log message.
