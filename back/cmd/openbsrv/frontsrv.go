package main

import (
	"context"
	"net/http"

	"github.com/codemodus/alfred"
)

type frontSrv struct {
	*http.Server
}

func newFrontSrv(dir, port string) (*frontSrv, error) {
	s := frontSrv{
		Server: &http.Server{
			Addr:    port,
			Handler: alfred.New(dir),
		},
	}

	return &s, nil
}

func (s *frontSrv) Serve() error {
	return s.Server.ListenAndServe()
}

func (s *frontSrv) Stop() error {
	// TODO: setup context
	return s.Shutdown(context.Background())
}
