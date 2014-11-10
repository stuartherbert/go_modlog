// Copyright (c) 2014-present Stuart Herbert
// Released under a 3-clause BSD license
package modlog

import (
	"github.com/stuartherbert/go_options"
)

// a Filter is a function that takes a log entry, and decides whether or not
// the entry can be logged
type LogFilter func(*options.OptionsStore, *LogEntry) bool

type Filterable interface {
	AddFilter(string, LogFilter)
	RemoveFilter(string)
}

// a list of the standard slots to insert filters into
const (
	LogLevelFilter = "loglevel"
)

func FilterLogToMinLevel(os *options.OptionsStore, entry *LogEntry) bool {
	// do we have a minimum level to look out for?
	option, ok := os.Option("minLogLevel")
	if !ok {
		// no we do not
		return true
	}

	minLogLevel := option.(LogLevel)
	// is the entry at a log level we are interested in?
	if entry.LogLevel <= minLogLevel {
		return true
	}

	return false
}
