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
	// read remote logging destintion from environment
	destination := os.Getenv("FISKALY_REMOTE_LOGGING_DESTINATION")
	if destination == "" {
		panic("env var FISKALY_REMOTE_LOGGING_DESTINATION not defined.")
	}

	// open UDP socket
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
