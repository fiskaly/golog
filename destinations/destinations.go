package destinations

import (
	"io"
	"net"
	"os"
	"strings"
)

var (
	connection io.WriteCloser
)

func NewRemote() (io.WriteCloser, error) {
	protocol, destination := readLoggingDestination()

	return establishConnection(protocol, destination)
}

func readLoggingDestination() (string, string) {
	fullDestination := os.Getenv("FISKALY_REMOTE_LOGGING_DESTINATION")
	if fullDestination == "" {
		panic("env var FISKALY_REMOTE_LOGGING_DESTINATION not defined.")
	}
	if strings.Count(fullDestination, ":") != 2 {
		panic("env var FISKALY_REMOTE_LOGGING_DESTINATION has an invalid format. Valid format: 'protocol:host:port'.")
	}

	splitDestination := strings.SplitN(fullDestination, ":", 2)
	protocol := splitDestination[0]
	if protocol != "tcp" && protocol != "udp" {
		panic("protocol must be either tcp or udp.")
	}
	destination := splitDestination[1]

	return protocol, destination
}

func establishConnection(protocol, destination string) (io.WriteCloser, error) {
	if connection == nil {
		if protocol == "tcp" {
			resolved, err := net.ResolveTCPAddr(protocol, destination)
			if err != nil {
				return nil, err
			}

			connection, err = net.DialTCP(protocol, nil, resolved)
			if err != nil {
				return nil, err
			}
		}

		if protocol == "udp" {
			resolved, err := net.ResolveUDPAddr(protocol, destination)
			if err != nil {
				return nil, err
			}

			connection, err = net.DialUDP(protocol, nil, resolved)
			if err != nil {
				return nil, err
			}
		}
	}
	return connection, nil
}
