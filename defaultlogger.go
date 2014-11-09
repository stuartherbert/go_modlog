// Copyright (c) 2014-present Stuart Herbert
// Released under the 3-clause BSD license
package modlog

// if the user just wants to use us as a drop-in replacement for the stdlib's
// log package, everything will go through our default logger
var defaultLogger *Logger

func init() {
	defaultLogger = NewLogger()
}

func Tracef(format string, args ...interface{}) {
	defaultLogger.Tracef(format, args...)
}

func Trace(args ...interface{}) {
	defaultLogger.Trace(args...)
}

func Traceln(args ...interface{}) {
	defaultLogger.Traceln(args...)
}

func Debugf(format string, args ...interface{}) {
	defaultLogger.Debugf(format, args...)
}
func Debug(args ...interface{}) {
	defaultLogger.Debug(args...)
}
func Debugln(args ...interface{}) {
	defaultLogger.Debugln(args...)
}

func Infof(format string, args ...interface{}) {
	defaultLogger.Infof(format, args...)
}
func Info(args ...interface{}) {
	defaultLogger.Info(args...)
}
func Infoln(args ...interface{}) {
	defaultLogger.Infoln(args...)
}

func Noticef(format string, args ...interface{}) {
	defaultLogger.Noticef(format, args...)
}
func Notice(args ...interface{}) {
	defaultLogger.Notice(args...)
}
func Noticeln(args ...interface{}) {
	defaultLogger.Noticeln(args...)
}

func Warnf(format string, args ...interface{}) {
	defaultLogger.Warnf(format, args...)
}

func Warn(args ...interface{}) {
	defaultLogger.Warn(args...)
}

func Warnln(args ...interface{}) {
	defaultLogger.Warnln(args...)
}

func Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args...)
}

func Error(args ...interface{}) {
	defaultLogger.Error(args...)
}

func Errorln(args ...interface{}) {
	defaultLogger.Errorln(args...)
}

func Criticalf(format string, args ...interface{}) {
	defaultLogger.Criticalf(format, args...)
}

func Critical(args ...interface{}) {
	defaultLogger.Critical(args...)
}

func Criticalln(args ...interface{}) {
	defaultLogger.Criticalln(args...)
}

func Alertf(format string, args ...interface{}) {
	defaultLogger.Alertf(format, args...)
}

func Alert(args ...interface{}) {
	defaultLogger.Alert(args...)
}

func Alertln(args ...interface{}) {
	defaultLogger.Alertln(args...)
}

func Emergencyf(format string, args ...interface{}) {
	defaultLogger.Emergencyf(format, args...)
}

func Emergency(args ...interface{}) {
	defaultLogger.Emergency(args...)
}

func Emergencyln(args ...interface{}) {
	defaultLogger.Emergencyln(args...)
}

func Fatal(args ...interface{}) {
	defaultLogger.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	defaultLogger.Fatalf(format, args...)
}

func Fatalln(args ...interface{}) {
	defaultLogger.Fatalln(args...)
}

func Flags() int {
	return defaultLogger.Flags()
}

func Output(calldepth int, s string) error {
	return defaultLogger.Output(calldepth, s)
}

func Panic(args ...interface{}) {
	defaultLogger.Panic(args...)
}

func Panicf(format string, args ...interface{}) {
	defaultLogger.Panicf(format, args...)
}

func Panicln(args ...interface{}) {
	defaultLogger.Panicln(args...)
}

func Prefix() string {
	return defaultLogger.Prefix()
}

func Print(args ...interface{}) {
	defaultLogger.Print(args...)
}

func Printf(format string, args ...interface{}) {
	defaultLogger.Printf(format, args...)
}

func Println(args ...interface{}) {
	defaultLogger.Println(args...)
}

func SetFlags(flags int) {
	defaultLogger.SetFlags(flags)
}

func SetPrefix(prefix string) {
	defaultLogger.SetPrefix(prefix)
}
