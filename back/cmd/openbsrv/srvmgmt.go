package main

import (
	"sync"
	"time"
)

type server interface {
	Serve() error
	Stop() error
}

type lineLogger interface {
	Info(format string, as ...interface{})
	Error(format string, as ...interface{})
}

type serverMgmt struct {
	ss  []server
	log lineLogger
}

func newServerMgmt(log lineLogger, ss ...server) *serverMgmt {
	return &serverMgmt{
		ss:  ss,
		log: log,
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
				m.log.Error("server error: %v", err)
			}
		}(s)

		time.Sleep(time.Millisecond * 200)
	}

	wg.Wait()

	return nil
}

func (m *serverMgmt) stop() error {
	for _, s := range m.ss {
		go func(s server) {
			// TODO: gather returned errors
			if err := s.Stop(); err != nil {
				m.log.Error("stop error: %v", err)
			}
		}(s)
	}

	return nil
}
