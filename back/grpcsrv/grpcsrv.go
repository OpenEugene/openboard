package grpcsrv

import (
	"fmt"

	"google.golang.org/grpc"
)

// Registerable describes services able to be registered with a grpc.Server.
type Registerable interface {
	RegisterWithGRPCServer(*grpc.Server) error
}

// GRPCSrv wraps a grpc.Server for convenience.
type GRPCSrv struct {
	*grpc.Server
}

// New returns a pointer to a basic GRPCSrv instance or an error.
func New() (*GRPCSrv, error) {
	opts := []grpc.ServerOption{}

	s := GRPCSrv{
		Server: grpc.NewServer(opts...),
	}

	return &s, nil
}

// RegisterServices registers services with the underlying grpc.Server.
func (s *GRPCSrv) RegisterServices(rs ...Registerable) error {
	for _, r := range rs {
		if err := r.RegisterWithGRPCServer(s.Server); err != nil {
			return fmt.Errorf("cannot register service: %s", err)
		}
	}

	return nil
}

// Serve sets up a tcp listener on the provided port and serves the underlying
// grpc.Server instance.
func (s *GRPCSrv) Serve(port string) error {
	we := func(err error) error {
		return fmt.Errorf("cannot serve: %s", err)
	}

	l, err := tcpListener(port)
	if err != nil {
		return we(err)
	}

	if err = s.Server.Serve(l); err != nil {
		// TODO: if err "contains closed" err, return nil

		return we(err)
	}

	return nil
}
