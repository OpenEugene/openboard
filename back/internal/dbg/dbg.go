package dbg

import (
	"io"
	"log"
	"sync/atomic"
)

// dbgLog is thread-safe.
type dbgLog struct {
	log *log.Logger
	out io.Writer
}

var atomicV atomic.Value
var dbg = new()

func new() *dbgLog {
	var dl dbgLog
	atomicV.Store(&dl)
	return &dl
}

func (dl *dbgLog) println(text string) {
	dbg.log.Println(text)
}

func (dl *dbgLog) printf(format string, as ...interface{}) {
	dbg.log.Printf(format+"\n", as...)
}

func Log(text string) {
	dbgLoad := atomicV.Load().(*dbgLog)
	if dbgLoad.out == nil {
		return
	}
	dbgLoad.println(text)
}

func Logf(format string, as ...interface{}) {
	dbgLoad := atomicV.Load().(*dbgLog)
	if dbgLoad.out == nil {
		return
	}
	dbgLoad.printf(format, as...)
}

func SetDebugOut(out io.Writer) {
	dbgLoad := atomicV.Load().(*dbgLog)
	dbgLoad.out = out

	if out != nil {
		dbgLoad.log = log.New(out, "", 0)
	} else {
		dbgLoad.log = log.Default()
	}

	atomicV.Store(dbgLoad)
}
