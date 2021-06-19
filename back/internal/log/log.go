// Package log allows for outputting information regarding application errors
// and events important to know for proper operation.
package log

import (
	"io"
	"log"
)

// Output is where logger will write to. Prefix is what will precede log text.
type Output struct {
	Out    io.Writer
	Prefix string
}

// Config allows for configuring the logger outputs for errors and information
// type log messages.
type Config struct {
	Err Output
	Inf Output
}

// Log is able to write information to a chosen output.
type Log struct {
	err *log.Logger
	inf *log.Logger
}

// New provides a pointer to a new instance of the Log object.
func New(c Config) *Log {
	return &Log{
		inf: log.New(c.Inf.Out, c.Inf.Prefix, 0),
		err: log.New(c.Err.Out, c.Err.Prefix, 0),
	}
}

// Info outputs information from the application.
func (log *Log) Info(format string, as ...interface{}) {
	log.inf.Printf(format+"\n", as...)
}

// Error outputs error information from the application.
func (log *Log) Error(format string, as ...interface{}) {
	log.err.Printf(format+"\n", as...)
}
