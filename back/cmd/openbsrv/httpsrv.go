package main

import (
	"github.com/codemodus/hedrs"

	"github.com/OpenEugene/openboard/back/internal/httpsrv"
)

type httpSrv struct {
	s *httpsrv.HTTPSrv

	rpcPort  string
	httpPort string
}

func newHTTPSrv(rpcPort, httpPort string, origins []string) (*httpSrv, error) {
	hs, err := httpsrv.New(hedrs.DefaultOrigins)
	if err != nil {
		return nil, err
	}

	s := httpSrv{
		s:        hs,
		rpcPort:  rpcPort,
		httpPort: httpPort,
	}

	return &s, nil
}

// Serve ...
func (s *httpSrv) Serve() error {
	return s.s.Serve(s.rpcPort, s.httpPort)
}

// Stop ...
func (s *httpSrv) Stop() error {
	return s.s.Stop()
}
