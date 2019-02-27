package grpcsrv

import (
	"fmt"
	"net"
)

func tcpListener(port string) (*net.TCPListener, error) {
	we := func(err error) error {
		return fmt.Errorf("cannot create tcp listener: %s", err)
	}

	a, err := net.ResolveTCPAddr("tcp", port)
	if err != nil {
		return nil, we(err)
	}

	l, err := net.ListenTCP("tcp", a)
	if err != nil {
		return nil, we(err)
	}

	return l, nil
}
