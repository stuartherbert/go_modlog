// Copyright (c) 2014-present Stuart Herbert
// Released under the 3-clause BSD license
package modlog

import (
	"io"

	"github.com/stuartherbert/go_options"
)

// the default list of options that we support
//
// this is used both in the main logger, and in each output
var optionsWhitelist options.ValidOptions

func init() {
	// setup the list of options that are supported
	optionsWhitelist = make(options.ValidOptions)
	optionsWhitelist["minLogLevel"] = "modlog.LogLevel"
}

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

// SetMinLogLevel() tells the logger to filter out some types of log messages
func SetMinLogLevel(level LogLevel) LogOption {
	return func(self *Logger) error {
		// record the log level
		self.Options.SetOption("minLogLevel", level)

		// add the required filter if needed
		if level < TraceLevel {
			self.AddFilter(LogLevelFilter, FilterLogToMinLevel)
		} else {
			self.RemoveFilter(LogLevelFilter)
		}
		return nil
	}
}
