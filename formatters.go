// Copyright (c) 2014-present Stuart Herbert
// All rights reserved
package modlog

import (
	"fmt"
	"log"
)

// LogFormatter is the signature that all output formatters much satisfy
type LogFormatter func(*Logger, *LogEntry) string

// FormattableLog is the interface for anything that supports formatting the
// metadata for a log message
type FormattableLog interface {
	AddFormatter(string, LogFormatter)
	RemoveFormatter(string)
}

// a list of the standard slots to insert formatters into
const (
	FormatTimestamp = "timestamp"
	FormatLogLevel  = "loglevel"
	FormatModule    = "module"
)

// DateTimeFormatter converts the 'When' field to the date/time format
// specified by the SetFlags() call
func StdlibDateTimeFormatter(logger *Logger, entry *LogEntry) string {
	output := ""

	// do we need to work out the date?
	if logger.StdlibFlags&log.Ldate != 0 {
		year, month, day := entry.When.Date()
		output = fmt.Sprintf("%04d/%02d/%02d", year, month, day)
	}

	// what about the time?
	if logger.StdlibFlags&log.Ltime != 0 {
		if len(output) > 0 {
			output = output + " "
		}

		hour, mins, secs := entry.When.Clock()
		output = output + fmt.Sprintf("%02d:%02d:%02d", hour, mins, secs)
	}

	// nanoseconds, anyone?
	if logger.StdlibFlags&log.Lmicroseconds != 0 {
		if len(output) > 0 {
			output = output + " "
		}

		ms := entry.When.Nanosecond() / 1e3
		output = output + fmt.Sprintf(".%06d", ms)
	}

	// all done
	return output
}

func StandardLogLevelFormatter(logger *Logger, entry *LogEntry) string {
	return fmt.Sprintf("%-9s", entry.LogLevel.String())
}

func ShortLogLevelFormatter(logger *Logger, entry *LogEntry) string {
	return entry.LogLevel.ShortString()
}

func StdlibPrefixFormatter(logger *Logger, entry *LogEntry) string {
	return logger.StdlibPrefix
}
