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
	log     *log.Logger
	atomVal atomic.Value
	toggle  bool
}

var dbg = new()

func new() *dbgLog {
	var dl dbgLog

	dl.atomVal.Store(dl.toggle)
	return &dl
}

func (dl *dbgLog) logln(as ...interface{}) {
	dbg.log.Println(as...)
}

func (dl *dbgLog) logf(format string, as ...interface{}) {
	dbg.log.Printf(format+"\n", as...)
}

// Log outputs information to help with application debugging.
func Log(as ...interface{}) {
	toggle := dbg.atomVal.Load().(bool)
	if toggle {
		dbg.logln(as...)
	}
}

// Logf outputs debugging information and is able to interpret formatting verbs.
func Logf(format string, as ...interface{}) {
	toggle := dbg.atomVal.Load().(bool)
	if toggle {
		dbg.logf(format, as...)
	}
}

// SetDebugOut allows for choosing where debug information will be written to.
func SetDebugOut(out io.Writer) {
	if out != nil {
		dbg.atomVal.Store(true)
		dbg.log = log.New(out, "", 0)
		return
	}

	dbg.atomVal.Store(false)
}
