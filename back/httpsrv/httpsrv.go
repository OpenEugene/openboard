package httpsrv

// HTTPSrv ...
type HTTPSrv struct{}

// New ...
func New() (*HTTPSrv, error) {
	return &HTTPSrv{}, nil
}

// Serve ...
func (s *HTTPSrv) Serve(rpcPort, httpPort string) error {
	return nil
}

// Stop ...
func (s *HTTPSrv) Stop() error {
	return nil
}
