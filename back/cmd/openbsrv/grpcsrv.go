package main

import (
	"database/sql"

	"github.com/OpenEugene/openboard/back/internal/authsvc"
	"github.com/OpenEugene/openboard/back/internal/grpcsrv"
	"github.com/OpenEugene/openboard/back/internal/postsvc"
	"github.com/OpenEugene/openboard/back/internal/usersvc"
)

type grpcSrv struct {
	s *grpcsrv.GRPCSrv

	port string
	svcs []interface{}
}

func newGRPCSrv(port string, db *sql.DB, drvr string) (*grpcSrv, error) {
	auth, err := authsvc.New()
	if err != nil {
		return nil, err
	}

	user, err := usersvc.New()
	if err != nil {
		return nil, err
	}

	post, err := postsvc.New(db, drvr, 123456)
	if err != nil {
		return nil, err
	}

	svcs := []interface{}{
		auth, user, post,
	}

	gs, err := grpcsrv.New()
	if err != nil {
		return nil, err
	}

	if err := registerServices(gs, svcs...); err != nil {
		return nil, err
	}

	s := grpcSrv{
		s:    gs,
		port: port,
		svcs: svcs,
	}

	return &s, nil
}

func (s *grpcSrv) services() []interface{} {
	return s.svcs
}

func (s *grpcSrv) Serve() error {
	return s.s.Serve(s.port)
}

func (s *grpcSrv) Stop() error {
	s.s.GracefulStop()
	return nil
}

func registerServices(srv *grpcsrv.GRPCSrv, svcs ...interface{}) error {
	for _, svc := range svcs {
		if s, ok := svc.(grpcsrv.Registerable); ok {
			if err := srv.RegisterServices(s); err != nil {
				return err
			}
		}
	}
	return nil
}
