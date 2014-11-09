// Copyright (c) 2014-present Stuart Herbert
// All rights reserved
package modlog

import (
	"time"
)

// LogFields is arbitrary data attached to a LogEntry
type LogFields map[string]interface{}

// LogEntry is a single log message that the caller wants to output somewhere
type LogEntry struct {
	// what level is this entry for?
	LogLevel LogLevel

	// who/what is writing to the log?
	Module string

	// what is the message?
	Message string

	// Any additional information?
	Data LogFields

	// when was the message generated?
	When time.Time
}

// NewLogEntry() creates a new log entry
//
// This will then get passed through the filter list and eventually will be
// send to all of the different log writers
func NewLogEntry(level LogLevel, module string, message string) *LogEntry {
	// create a new LogEntry
	retval := &LogEntry{
		LogLevel: level,
		Module:   module,
		Message:  message,
		Data:     make(LogFields, 5),
		When:     time.Now(),
	}

	// all done
	return retval
}
