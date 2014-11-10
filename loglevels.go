// Copyright (c) 2014-present Stuart Herbert
// Released under the 3-clause BSD license
package modlog

// LogLevel is how we represent log levels internally
type LogLevel uint8

// a list of the log levels that you can set
const (
	EmergencyLevel = LogLevel(0)
	AlertLevel     = LogLevel(1)
	CriticalLevel  = LogLevel(2)
	ErrorLevel     = LogLevel(3)
	WarnLevel      = LogLevel(4)
	NoticeLevel    = LogLevel(5)
	InfoLevel      = LogLevel(6)
	DebugLevel     = LogLevel(7)
	TraceLevel     = LogLevel(8)
	PanicLevel     = LogLevel(0)
	FatalLevel     = LogLevel(0)
)

// a helper for translating strings to log levels
var LogLevels = map[string]LogLevel{
	"emergency": 0,
	"emerg":     0,
	"alert":     1,
	"critical":  2,
	"crit":      2,
	"error":     3,
	"err":       3,
	"warn":      4,
	"notice":    5,
	"not":       5,
	"info":      6,
	"debug":     7,
	"trace":     8,
	"panic":     0,
	"fatal":     0,
}

var LogLevelNames = map[LogLevel]string{
	EmergencyLevel: "EMERGENCY",
	AlertLevel:     "ALERT",
	CriticalLevel:  "CRITICAL",
	ErrorLevel:     "ERROR",
	WarnLevel:      "WARNING",
	NoticeLevel:    "NOTICE",
	InfoLevel:      "INFO",
	DebugLevel:     "DEBUG",
	TraceLevel:     "TRACE",
}

var LogLevelShortNames = map[LogLevel]string{
	EmergencyLevel: "EMERG ",
	AlertLevel:     "ALERT ",
	CriticalLevel:  "CRIT  ",
	ErrorLevel:     "ERROR ",
	WarnLevel:      "WARN  ",
	NoticeLevel:    "NOTICE",
	InfoLevel:      "INFO  ",
	DebugLevel:     "DEBUG ",
	TraceLevel:     "TRACE ",
}

func (self *LogLevel) String() string {
	return LogLevelNames[*self]
}

func (self *LogLevel) ShortString() string {
	return LogLevelShortNames[*self]
}
