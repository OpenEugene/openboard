// Package dbg allows for outputting information that can help with debugging
// the application.
package dbg

import (
	"io"
	"log"
	"sync/atomic"
)

// dbg is thread-safe.
type dbg struct {
	log *log.Logger
	// Turn on and off dbg by setting 1 and 0, respectively.
	toggle uint32
}

func new() *dbg {
	return &dbg{}
}

func (d *dbg) logln(as ...interface{}) {
	tog := atomic.LoadUint32(&d.toggle)
	if tog == 1 {
		d.log.Println(as...)
	}
}

func (d *dbg) logf(format string, as ...interface{}) {
	tog := atomic.LoadUint32((&d.toggle))
	if tog == 1 {
		d.log.Printf(format+"\n", as...)
	}
}

var debug = new()

// Log outputs information to help with application debugging.
func Log(as ...interface{}) {
	debug.logln(as...)
}

// Logf outputs debugging information and is able to interpret formatting verbs.
func Logf(format string, as ...interface{}) {
	debug.logf(format, as...)
}

// SetDebugOut allows for choosing where debug information will be written to.
func SetDebugOut(out io.Writer) {
	if out != nil {
		atomic.StoreUint32(&debug.toggle, 1)
		debug.log = log.New(out, "", 0)
		return
	}

	atomic.StoreUint32(&debug.toggle, 0)
}
