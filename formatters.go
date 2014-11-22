// Copyright (c) 2014-present Stuart Herbert
// All rights reserved
package modlog

import (
	"fmt"
	"log"
	"path"
	"runtime"
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
	FormatFilename  = "filename"
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
	//
	// note: Ltime is optional if Lmicroseconds is set
	//
	// Implicit behaviour :(
	if logger.StdlibFlags&log.Ltime != 0 || logger.StdlibFlags&log.Lmicroseconds != 0 {
		if len(output) > 0 {
			output = output + " "
		}

		hour, mins, secs := entry.When.Clock()
		output = output + fmt.Sprintf("%02d:%02d:%02d", hour, mins, secs)
	}

	// nanoseconds, anyone?
	if logger.StdlibFlags&log.Lmicroseconds != 0 {
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

func StdlibFileFormatter(logger *Logger, entry *LogEntry) string {
	if logger.StdlibFlags&(Lshortfile|Llongfile) == 0 {
		return ""
	}

	// get the caller
	_, file, line, ok := runtime.Caller(5)
	if !ok {
		return "unknown:00"
	}

	// Lshortfile overrides Llongfile in the upstream_tests
	if logger.StdlibFlags&Lshortfile != 0 {
		return fmt.Sprintf("%s:%d", path.Base(file), line)
	} else {
		return fmt.Sprintf("%s:%d", file, line)
	}
}
