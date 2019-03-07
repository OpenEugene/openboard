package main

import (
	"github.com/champagneabuelo/openboard/back/authsvc"
	"github.com/champagneabuelo/openboard/back/internal/grpcsrv"
	"github.com/champagneabuelo/openboard/back/usersvc"
)

type grpcSrv struct {
	s *grpcsrv.GRPCSrv

	port string
}

func newGRPCSrv(port string) (*grpcSrv, error) {
	auth, err := authsvc.New()
	if err != nil {
		return nil, err
	}

	user, err := usersvc.New()
	if err != nil {
		return nil, err
	}

	gs, err := grpcsrv.New()
	if err != nil {
		return nil, err
	}

	if err := gs.RegisterServices(auth, user); err != nil {
		return nil, err
	}

	s := grpcSrv{
		s:    gs,
		port: port,
	}

	return &s, nil
}

func (s *grpcSrv) Serve() error {
	return s.s.Serve(s.port)
}

func (s *grpcSrv) Stop() error {
	s.s.GracefulStop()
	return nil
}
