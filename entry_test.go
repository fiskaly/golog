package golog

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFromStdHTTPRequest(t *testing.T) {
	var buf bytes.Buffer

	fields := Fields{
		"test": "test",
	}

	request := httptest.NewRequest(http.MethodGet, "/something", nil)
	logger := NewLogger(&buf, Fields{})

	logger.Info("something", FromStdHTTPRequest(request), fields)

	content, err := io.ReadAll(&buf)
	require.NoError(t, err)

	var logMessage entry
	err = json.Unmarshal(content, &logMessage)
	require.NoError(t, err)

	require.NotNil(t, logMessage.HTTPRequest)
	require.Equal(t, http.MethodGet, logMessage.HTTPRequest.Method)
	require.Equal(t, request.UserAgent(), logMessage.HTTPRequest.UserAgent)
}
