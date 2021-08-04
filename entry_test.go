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
	buffer := bytes.NewBuffer([]byte{})

	request := httptest.NewRequest(http.MethodGet, "/something", nil)
	logger := NewLogger(buffer, Fields{})

	logger.Info("something", FromStdHTTPRequest(request), nil)

	contents, err := io.ReadAll(buffer)
	require.NoError(t, err)

	var logMessage entry
	err = json.Unmarshal(contents, &logMessage)
	require.NoError(t, err)

	require.Equal(t, http.MethodGet, logMessage.httpRequest.Method)
}
