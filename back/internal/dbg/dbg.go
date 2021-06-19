
// Package dbg allows for outputting information that can help with debugging
// the application.
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

// Log outputs information to help with application debugging.
func Log(text string) {
	dbgLoad := atomicV.Load().(*dbgLog)
	if dbgLoad.out == nil {
		return
	}
	dbgLoad.println(text)
}

// Logf outputs debugging information and is able to interpret formatting verbs.
func Logf(format string, as ...interface{}) {
	dbgLoad := atomicV.Load().(*dbgLog)
	if dbgLoad.out == nil {
		return
	}
	dbgLoad.printf(format, as...)
}

// SetDebugOut allows for choosing where debug information will be written to.
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
