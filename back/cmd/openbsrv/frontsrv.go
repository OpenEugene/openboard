package main

import (
	"context"
	"net/http"

	"github.com/codemodus/alfred"
)

type frontSrv struct {
	s *http.Server
}

func newFrontSrv(dir, port string) (*frontSrv, error) {
	s := frontSrv{
		s: &http.Server{
			Addr:    port,
			Handler: alfred.New(dir),
		},
	}

	return &s, nil
}

func (s *frontSrv) Serve() error {
	return s.s.ListenAndServe()
}

func (s *frontSrv) Stop() error {
	// TODO: setup context
	return s.s.Shutdown(context.Background())
}
