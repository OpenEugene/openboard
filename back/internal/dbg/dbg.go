package dbg

import (
	"io"
	"log"
	"sync/atomic"
)

var atomicV atomic.Value

// DbgLog is thread-safe.
type DbgLog struct {
	log *log.Logger
}

func New(out io.Writer) *DbgLog {
	var dl DbgLog

	if out == nil {
		atomicV.Store(&dl)
		return &dl
	}

	dl = DbgLog{
		log: log.New(out, "[debug] ", log.Ldate|log.Ltime),
	}
	atomicV.Store(&dl)
	return &dl
}

func (l *DbgLog) Log(format string, as ...interface{}) {
	load := atomicV.Load().(*DbgLog)

	if load.log == nil {
		return
	}

	load.log.Printf(format+"\n", as...)
}

func (l *DbgLog) Off() {
	l.log = nil
	atomicV.Store(l)
}

func (l *DbgLog) On(out io.Writer) {
	l.log = log.New(out, "[debug] ", log.Ldate|log.Ltime)
	atomicV.Store(l)
}
