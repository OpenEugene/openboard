// Package dbg allows for outputting information that can help with debugging
// the application.
package dbg

import (
	"io"
	"log"
	"sync"
)

// dbg is thread-safe.
type dbg struct {
	mutex sync.Mutex
	log   *log.Logger
}

func new() *dbg {
	return &dbg{}
}

func (d *dbg) logln(as ...interface{}) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if d.log != nil {
		d.log.Println(as...)
	}
}

func (d *dbg) logf(format string, as ...interface{}) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if d.log != nil {
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
	debug.mutex.Lock()
	defer debug.mutex.Unlock()

	if out != nil {
		debug.log = log.New(out, "", 0)
		return
	}

	debug.log = nil
}
