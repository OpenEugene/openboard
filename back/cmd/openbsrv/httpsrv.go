package main

import (
	"github.com/codemodus/hedrs"

	"github.com/OpenEugene/openboard/back/internal/httpsrv"
	"github.com/OpenEugene/openboard/back/internal/logsvc"
)

type httpSrv struct {
	s *httpsrv.HTTPSrv

	rpcPort  string
	httpPort string

	log logsvc.LineLogger
}

func newHTTPSrv(log logsvc.LineLogger, rpcPort, httpPort string, origins []string) (*httpSrv, error) {
	hs, err := httpsrv.New(hedrs.DefaultOrigins)
	if err != nil {
		return nil, err
	}

	s := httpSrv{
		s:        hs,
		rpcPort:  rpcPort,
		httpPort: httpPort,
		log:      log,
	}

	return &s, nil
}

// Serve ...
func (s *httpSrv) Serve() error {
	s.log.Info("starting HTTP server on port %s", s.httpPort)
	return s.s.Serve(s.rpcPort, s.httpPort)
}

// Stop ...
func (s *httpSrv) Stop() error {
	return s.s.Stop()
}
