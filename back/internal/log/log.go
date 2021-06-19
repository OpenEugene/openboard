package log

import (
	"io"
	"log"
)

type Output struct {
	Out    io.Writer
	Prefix string
}

type Config struct {
	Err Output
	Inf Output
}

type Log struct {
	err *log.Logger
	inf *log.Logger
}

func New(c Config) *Log {
	return &Log{
		inf: log.New(c.Inf.Out, c.Inf.Prefix, 0),
		err: log.New(c.Err.Out, c.Err.Prefix, 0),
	}
}

func (log *Log) Info(format string, as ...interface{}) {
	log.inf.Printf(format+"\n", as...)
}

func (log *Log) Error(format string, as ...interface{}) {
	log.err.Printf(format+"\n", as...)
}
