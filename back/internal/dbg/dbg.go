package dbg

import (
	"io"
	"log"
)

type DbgLog struct {
	log *log.Logger
}

func (l *DbgLog) Log(format string, as ...interface{}) {
	if l.log == nil {
		return
	}

	l.log.Printf(format+"\n", as...)
}

func New(out io.Writer) *DbgLog {
	if out == nil {
		return &DbgLog{}
	}

	return &DbgLog{
		log: log.New(out, "[debug] ", log.Ldate|log.Ltime),
	}
}
