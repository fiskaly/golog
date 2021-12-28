package destinations

import (
	"io"
	"net"
	"os"
)

var (
	connection io.Writer
)

func NewRemote() (io.Writer, error) {
	destination := readLoggingDestination()

	return getConnection(destination)
}

func readLoggingDestination() string {
	destination := os.Getenv("FISKALY_REMOTE_LOGGING_DESTINATION")
	if destination == "" {
		panic("env var FISKALY_REMOTE_LOGGING_DESTINATION not defined.")
	}
	return destination
}

func getConnection(destination string) (io.Writer, error) {
	if connection == nil {
		resolved, err := net.ResolveUDPAddr("udp", destination)
		if err != nil {
			return nil, err
		}

		connection, err = net.DialUDP("udp", nil, resolved)
		if err != nil {
			return nil, err
		}
	}
	return connection, nil
}
