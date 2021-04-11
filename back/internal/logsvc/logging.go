package logsvc

import (
	"io"
	"log"
)

type LineLogger interface {
	Info(format string, as ...interface{})
	Error(format string, as ...interface{})
}

type ServerLog struct {
	err, inf *log.Logger
}

func (log *ServerLog) Info(format string, as ...interface{}) {
	log.inf.Printf(format+"\n", as...)
}

func (log *ServerLog) Error(format string, as ...interface{}) {
	log.err.Printf(format+"\n", as...)
}

type Handle struct {
	Err, Inf io.Writer
}

func NewServerLog(h Handle) *ServerLog {
	return &ServerLog{
		inf: log.New(h.Inf, "[info] ", log.Ldate|log.Ltime|log.Lshortfile),
		err: log.New(h.Err, "[error] ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}
