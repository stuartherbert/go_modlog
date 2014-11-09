// Copyright (c) 2014-present Stuart Herbert
// Released under a 3-clause BSD license
package modlog

// a Filter is a function that takes a log entry, and decides whether or not
// the entry can be logged
type LogFilter func(*Logger, *LogEntry) bool

type FilterableLog interface {
	AddFilter(string, LogFilter)
	RemoveFilter(string)
}

// a list of the standard slots to insert filters into
const (
	LogLevelFilter = "loglevel"
)

func RestrictToLogLevel(logger *Logger, entry *LogEntry) bool {
	// is the entry at a log level we are interested in?
	if entry.LogLevel <= logger.MinLogLevel {
		return true
	}

	return false
}
