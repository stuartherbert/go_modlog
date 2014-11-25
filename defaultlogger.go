// Copyright (c) 2014-present Stuart Herbert
// Released under the 3-clause BSD license
package modlog

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/stuartherbert/go_extras/extrafmt"
)

// if the user just wants to use us as a drop-in replacement for the stdlib's
// log package, everything will go through our default logger
var defaultLogger *Logger

func init() {
	// we have to do this to guarantee order
	initOptionsWhitelist()
	defaultLogger = NewLogger(
		SetStdlibFlags(log.LstdFlags),
	)
	defaultLogger.SetOutput(os.Stdout)
}

func DefaultLogger() *Logger {
	return defaultLogger
}

func SetOptions(logOptions ...LogOption) {
	defaultLogger.SetOptions(logOptions...)
}

func Flags() int {
	return defaultLogger.Flags()
}

func Output(calldepth int, s string) error {
	return defaultLogger.Output(calldepth, s)
}

func Prefix() string {
	return defaultLogger.Prefix()
}

func SetFlags(flags int) {
	defaultLogger.SetFlags(flags)
}

func SetPrefix(prefix string) {
	defaultLogger.SetPrefix(prefix)
}

func SetOutput(out io.Writer) {
	defaultLogger.SetOutput(out)
}

func Tracef(format string, args ...interface{}) {
	defaultLogger.AddLogEntry(TraceLevel, "", fmt.Sprintf(format, args...))
}

func Trace(args ...interface{}) {
	defaultLogger.AddLogEntry(TraceLevel, "", fmt.Sprint(args...))
}

func Traceln(args ...interface{}) {
	defaultLogger.AddLogEntry(TraceLevel, "", extrafmt.Sprintnln(args...))
}

func Debugf(format string, args ...interface{}) {
	defaultLogger.AddLogEntry(DebugLevel, "", fmt.Sprintf(format, args...))
}

func Debug(args ...interface{}) {
	defaultLogger.AddLogEntry(DebugLevel, "", fmt.Sprint(args...))
}

func Debugln(args ...interface{}) {
	defaultLogger.AddLogEntry(DebugLevel, "", extrafmt.Sprintnln(args...))
}

func Infof(format string, args ...interface{}) {
	defaultLogger.AddLogEntry(InfoLevel, "", fmt.Sprintf(format, args...))
}

func Info(args ...interface{}) {
	defaultLogger.AddLogEntry(InfoLevel, "", fmt.Sprint(args...))
}

func Infoln(args ...interface{}) {
	defaultLogger.AddLogEntry(InfoLevel, "", extrafmt.Sprintnln(args...))
}

func Noticef(format string, args ...interface{}) {
	defaultLogger.AddLogEntry(NoticeLevel, "", fmt.Sprintf(format, args...))
}

func Notice(args ...interface{}) {
	defaultLogger.AddLogEntry(NoticeLevel, "", fmt.Sprint(args...))
}

func Noticeln(args ...interface{}) {
	defaultLogger.AddLogEntry(NoticeLevel, "", extrafmt.Sprintnln(args...))
}

func Warnf(format string, args ...interface{}) {
	defaultLogger.AddLogEntry(WarnLevel, "", fmt.Sprintf(format, args...))
}

func Warn(args ...interface{}) {
	defaultLogger.AddLogEntry(WarnLevel, "", fmt.Sprint(args...))
}

func Warnln(args ...interface{}) {
	defaultLogger.AddLogEntry(WarnLevel, "", extrafmt.Sprintnln(args...))
}

func Errorf(format string, args ...interface{}) {
	defaultLogger.AddLogEntry(ErrorLevel, "", fmt.Sprintf(format, args...))
}

func Error(args ...interface{}) {
	defaultLogger.AddLogEntry(ErrorLevel, "", fmt.Sprint(args...))
}

func Errorln(args ...interface{}) {
	defaultLogger.AddLogEntry(ErrorLevel, "", extrafmt.Sprintnln(args...))
}

func Criticalf(format string, args ...interface{}) {
	defaultLogger.AddLogEntry(CriticalLevel, "", fmt.Sprintf(format, args...))
}

func Critical(args ...interface{}) {
	defaultLogger.AddLogEntry(CriticalLevel, "", fmt.Sprint(args...))
}

func Criticalln(args ...interface{}) {
	defaultLogger.AddLogEntry(CriticalLevel, "", extrafmt.Sprintnln(args...))
}

func Alertf(format string, args ...interface{}) {
	defaultLogger.AddLogEntry(AlertLevel, "", fmt.Sprintf(format, args...))
}

func Alert(args ...interface{}) {
	defaultLogger.AddLogEntry(AlertLevel, "", fmt.Sprint(args...))
}

func Alertln(args ...interface{}) {
	defaultLogger.AddLogEntry(AlertLevel, "", extrafmt.Sprintnln(args...))
}

func Emergencyf(format string, args ...interface{}) {
	defaultLogger.AddLogEntry(EmergencyLevel, "", fmt.Sprintf(format, args...))
}

func Emergency(args ...interface{}) {
	defaultLogger.AddLogEntry(EmergencyLevel, "", fmt.Sprint(args...))
}

func Emergencyln(args ...interface{}) {
	defaultLogger.AddLogEntry(EmergencyLevel, "", extrafmt.Sprintnln(args...))
}

func Fatalf(format string, args ...interface{}) {
	defaultLogger.AddLogEntry(FatalLevel, "", fmt.Sprintf(format, args...))
}

func Fatal(args ...interface{}) {
	defaultLogger.AddLogEntry(FatalLevel, "", fmt.Sprint(args...))
}

func Fatalln(args ...interface{}) {
	defaultLogger.AddLogEntry(FatalLevel, "", extrafmt.Sprintnln(args...))
}

func Panicf(format string, args ...interface{}) {
	defaultLogger.AddLogEntry(PanicLevel, "", fmt.Sprintf(format, args...))
}

func Panic(args ...interface{}) {
	defaultLogger.AddLogEntry(PanicLevel, "", fmt.Sprint(args...))
}

func Panicln(args ...interface{}) {
	defaultLogger.AddLogEntry(PanicLevel, "", extrafmt.Sprintnln(args...))
}

func Printf(format string, args ...interface{}) {
	defaultLogger.AddLogEntry(InfoLevel, "", fmt.Sprintf(format, args...))
}

func Print(args ...interface{}) {
	defaultLogger.AddLogEntry(InfoLevel, "", fmt.Sprint(args...))
}

func Println(args ...interface{}) {
	defaultLogger.AddLogEntry(InfoLevel, "", extrafmt.Sprintnln(args...))
}
