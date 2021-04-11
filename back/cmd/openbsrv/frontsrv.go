package main

import (
	"context"
	"net/http"

	"github.com/codemodus/alfred"
	"github.com/codemodus/chain/v2"
	"github.com/codemodus/hedrs"

	"github.com/OpenEugene/openboard/back/internal/logsvc"
)

type frontSrv struct {
	s *http.Server

	log logsvc.LineLogger
}

func newFrontSrv(log logsvc.LineLogger, port, dir string, origins []string) (*frontSrv, error) {
	origins = append(hedrs.DefaultOrigins, origins...)
	corsOrigins := hedrs.CORSOrigins(hedrs.NewAllowed(origins...))
	corsMethods := hedrs.CORSMethods(hedrs.NewValues(hedrs.AllMethods...))
	corsHeaders := hedrs.CORSHeaders(hedrs.NewValues(hedrs.DefaultHeaders...))

	cmn := chain.New(
		corsOrigins,
		corsMethods,
		corsHeaders,
	)

	s := frontSrv{
		s: &http.Server{
			Addr:    port,
			Handler: cmn.End(alfred.New(dir)),
		},
		log: log,
	}

	return &s, nil
}

func (s *frontSrv) Serve() error {
	s.log.Info("starting frontend server on port %s", s.s.Addr)
	if err := s.s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *frontSrv) Stop() error {
	// TODO: setup context
	return s.s.Shutdown(context.Background())
}
