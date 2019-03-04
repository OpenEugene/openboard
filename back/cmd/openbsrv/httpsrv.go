package main

import "github.com/champagneabuelo/openboard/back/httpsrv"

type httpSrv struct {
	s *httpsrv.HTTPSrv

	rpcPort  string
	httpPort string
}

func newHTTPSrv(rpcPort, httpPort string) (*httpSrv, error) {
	hs, err := httpsrv.New()
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
