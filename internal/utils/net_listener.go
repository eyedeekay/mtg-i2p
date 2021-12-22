package utils

import (
	"fmt"
	"net"
	"strings"

	sam "github.com/eyedeekay/sam3/helper"
)

type Listener struct {
	net.Listener
}

func (l Listener) Accept() (net.Conn, error) {
	conn, err := l.Listener.Accept()
	if err != nil {
		return nil, err // nolint: wrapcheck
	}

	return conn, nil
}

func NewListener(bindTo string, bufferSize int) (net.Listener, error) {
	if strings.HasSuffix(bindTo, ".i2p") {
		base, err := sam.I2PListener(bindTo, "127.0.0.1:7656", bindTo)
		if err != nil {
			return nil, fmt.Errorf("cannot build a base I2P listener: %w", err)
		}
		return Listener{
			Listener: base,
		}, nil
	}
	base, err := net.Listen("tcp", bindTo)
	if err != nil {
		return nil, fmt.Errorf("cannot build a base listener: %w", err)
	}

	return Listener{
		Listener: base,
	}, nil
}
