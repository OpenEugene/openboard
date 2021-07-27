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
	atomVal atomic.Value
}

func new() *dbg {
	var d dbg
	var logr *log.Logger
	d.atomVal.Store(logr)
	return &d
}

func (d *dbg) logln(as ...interface{}) {
	logr := d.atomVal.Load().(*log.Logger)
	if logr != nil {
		logr.Println(as...)
	}
}

func (d *dbg) logf(format string, as ...interface{}) {
	logr := d.atomVal.Load().(*log.Logger)
	if logr != nil {
		logr.Printf(format+"\n", as...)
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
	var logr *log.Logger

	if out != nil {
		logr = log.New(out, "", 0)
		debug.atomVal.Store(logr)
		return
	}

	debug.atomVal.Store(logr)
}
