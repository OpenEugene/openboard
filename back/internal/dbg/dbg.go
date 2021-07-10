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
	log     *log.Logger
	atomVal atomic.Value
	toggle  bool
}

func new() *dbg {
	var d dbg

	d.atomVal.Store(d.toggle)
	return &d
}

func (d *dbg) logln(as ...interface{}) {
	toggle := d.atomVal.Load().(bool)
	if toggle {
		d.log.Println(as...)
	}
}

func (d *dbg) logf(format string, as ...interface{}) {
	toggle := d.atomVal.Load().(bool)
	if toggle {
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
		debug.atomVal.Store(true)
		debug.log = log.New(out, "", 0)
		return
	}

	debug.atomVal.Store(false)
}
