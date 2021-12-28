package destinations

import (
	"io"
	"net"
	"os"
)

var (
	connection io.WriteCloser
)

func NewRemote() (io.WriteCloser, error) {
	destination := readLoggingDestination()

	return establishConnection(destination)
}

func readLoggingDestination() string {
	destination := os.Getenv("FISKALY_REMOTE_LOGGING_DESTINATION")
	if destination == "" {
		panic("env var FISKALY_REMOTE_LOGGING_DESTINATION not defined.")
	}
	return destination
}

func establishConnection(destination string) (io.WriteCloser, error) {
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
