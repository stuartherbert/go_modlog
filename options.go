// Copyright (c) 2014-present Stuart Herbert
// Released under the 3-clause BSD license
package modlog

import (
	"io"
)

// LogOption is the signature that all logging option functions must match
//
// See http://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
// for the thinking behind functional options
type LogOption func(*Logger) error

// SetDefaultOutput() tells the logger to write to the given output
func SetDefaultOutput(out io.Writer) LogOption {
	return func(self *Logger) error {
		self.AddOutput("default", out)
		return nil
	}
}

// SetStdlibFlags() tells the logger to honour flags supported by stdlib's
// log package
func SetStdlibFlags(flags int) LogOption {
	return func(self *Logger) error {
		self.SetFlags(flags)
		return nil
	}
}

// SetStdlibPrefix() tells the logger to use the given prefix at the front
// of log messages
func SetStdlibPrefix(prefix string) LogOption {
	return func(self *Logger) error {
		self.SetPrefix(prefix)
		return nil
	}
}
