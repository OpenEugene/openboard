package main

import (
	"fmt"
	"os"
	"sync"
)

type server interface {
	Serve() error
	Stop() error
}

type serverMgmt struct {
	ss []server
}

func newServerMgmt(ss ...server) *serverMgmt {
	return &serverMgmt{
		ss: ss,
	}
}

func (m *serverMgmt) serve() error {
	var wg sync.WaitGroup
	wg.Add(len(m.ss))

	for _, s := range m.ss {
		go func(s server) {
			defer wg.Done()

			// TODO: gather returned errors
			if err := s.Serve(); err != nil {
				fmt.Fprintln(os.Stderr, "server error:", err)
			}
		}(s)
	}

	wg.Wait()

	return nil
}

func (m *serverMgmt) stop() error {
	for _, s := range m.ss {
		go func(s server) {
			// TODO: gather returned errors
			if err := s.Stop(); err != nil {
				fmt.Fprintln(os.Stderr, "stop error:", err)
			}
		}(s)
	}

	return nil
}
