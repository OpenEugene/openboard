package log

import (
	"io"
	"log"
)

type Log struct {
	err, inf *log.Logger
}

func (log *Log) Info(format string, as ...interface{}) {
	log.inf.Printf(format+"\n", as...)
}

func (log *Log) Error(format string, as ...interface{}) {
	log.err.Printf(format+"\n", as...)
}

type Outputs struct {
	Err, Inf io.Writer
}

func New(h Outputs) *Log {
	return &Log{
		inf: log.New(h.Inf, "[info] ", log.Ldate|log.Ltime),
		err: log.New(h.Err, "[error] ", log.Ldate|log.Ltime),
	}
}
