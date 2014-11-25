// Copyright (c) 2014-present Stuart Herbert
// Released under a 3-clause BSD license
package modlog

import (
	"fmt"
	"github.com/stuartherbert/go_extras/extrafmt"
)

type ModPrint func(string, ...interface{})
type ModPrintf func(string, string, ...interface{})
type ModPrintln func(string, ...interface{})

func ModTracef(modName string, format string, args ...interface{}) {
	defaultLogger.AddLogEntry(TraceLevel, modName, fmt.Sprintf(format, args...))
}

func ModTrace(modName string, args ...interface{}) {
	defaultLogger.AddLogEntry(TraceLevel, modName, fmt.Sprint(args...))
}

func ModTraceln(modName string, args ...interface{}) {
	defaultLogger.AddLogEntry(TraceLevel, modName, extrafmt.Sprintnln(args...))
}

func ModDebugf(modName string, format string, args ...interface{}) {
	defaultLogger.AddLogEntry(DebugLevel, modName, fmt.Sprintf(format, args...))
}

func ModDebug(modName string, args ...interface{}) {
	defaultLogger.AddLogEntry(DebugLevel, modName, fmt.Sprint(args...))
}

func ModDebugln(modName string, args ...interface{}) {
	defaultLogger.AddLogEntry(DebugLevel, modName, extrafmt.Sprintnln(args...))
}

func ModInfof(modName string, format string, args ...interface{}) {
	defaultLogger.AddLogEntry(InfoLevel, modName, fmt.Sprintf(format, args...))
}

func ModInfo(modName string, args ...interface{}) {
	defaultLogger.AddLogEntry(InfoLevel, modName, fmt.Sprint(args...))
}

func ModInfoln(modName string, args ...interface{}) {
	defaultLogger.AddLogEntry(InfoLevel, modName, extrafmt.Sprintnln(args...))
}

func ModNoticef(modName string, format string, args ...interface{}) {
	defaultLogger.AddLogEntry(NoticeLevel, modName, fmt.Sprintf(format, args...))
}

func ModNotice(modName string, args ...interface{}) {
	defaultLogger.AddLogEntry(NoticeLevel, modName, fmt.Sprint(args...))
}

func ModNoticeln(modName string, args ...interface{}) {
	defaultLogger.AddLogEntry(NoticeLevel, modName, extrafmt.Sprintnln(args...))
}

func ModWarnf(modName string, format string, args ...interface{}) {
	defaultLogger.AddLogEntry(WarnLevel, modName, fmt.Sprintf(format, args...))
}

func ModWarn(modName string, args ...interface{}) {
	defaultLogger.AddLogEntry(WarnLevel, modName, fmt.Sprint(args...))
}

func ModWarnln(modName string, args ...interface{}) {
	defaultLogger.AddLogEntry(WarnLevel, modName, extrafmt.Sprintnln(args...))
}

func ModErrorf(modName string, format string, args ...interface{}) {
	defaultLogger.AddLogEntry(ErrorLevel, modName, fmt.Sprintf(format, args...))
}

func ModError(modName string, args ...interface{}) {
	defaultLogger.AddLogEntry(ErrorLevel, modName, fmt.Sprint(args...))
}

func ModErrorln(modName string, args ...interface{}) {
	defaultLogger.AddLogEntry(ErrorLevel, modName, extrafmt.Sprintnln(args...))
}

func ModCriticalf(modName string, format string, args ...interface{}) {
	defaultLogger.AddLogEntry(CriticalLevel, modName, fmt.Sprintf(format, args...))
}

func ModCritical(modName string, args ...interface{}) {
	defaultLogger.AddLogEntry(CriticalLevel, modName, fmt.Sprint(args...))
}

func ModCriticalln(modName string, args ...interface{}) {
	defaultLogger.AddLogEntry(CriticalLevel, modName, extrafmt.Sprintnln(args...))
}

func ModAlertf(modName string, format string, args ...interface{}) {
	defaultLogger.AddLogEntry(AlertLevel, modName, fmt.Sprintf(format, args...))
}

func ModAlert(modName string, args ...interface{}) {
	defaultLogger.AddLogEntry(AlertLevel, modName, fmt.Sprint(args...))
}

func ModAlertln(modName string, args ...interface{}) {
	defaultLogger.AddLogEntry(AlertLevel, modName, extrafmt.Sprintnln(args...))
}

func ModEmergencyf(modName string, format string, args ...interface{}) {
	defaultLogger.AddLogEntry(EmergencyLevel, modName, fmt.Sprintf(format, args...))
}

func ModEmergency(modName string, args ...interface{}) {
	defaultLogger.AddLogEntry(EmergencyLevel, modName, fmt.Sprint(args...))
}

func ModEmergencyln(modName string, args ...interface{}) {
	defaultLogger.AddLogEntry(EmergencyLevel, modName, extrafmt.Sprintnln(args...))
}

func ModFatalf(modName string, format string, args ...interface{}) {
	defaultLogger.AddLogEntry(FatalLevel, modName, fmt.Sprintf(format, args...))
}

func ModFatal(modName string, args ...interface{}) {
	defaultLogger.AddLogEntry(FatalLevel, modName, fmt.Sprint(args...))
}

func ModFatalln(modName string, args ...interface{}) {
	defaultLogger.AddLogEntry(FatalLevel, modName, extrafmt.Sprintnln(args...))
}

func ModPanicf(modName string, format string, args ...interface{}) {
	defaultLogger.AddLogEntry(PanicLevel, modName, fmt.Sprintf(format, args...))
}

func ModPanic(modName string, args ...interface{}) {
	defaultLogger.AddLogEntry(PanicLevel, modName, fmt.Sprint(args...))
}

func ModPanicln(modName string, args ...interface{}) {
	defaultLogger.AddLogEntry(PanicLevel, modName, extrafmt.Sprintnln(args...))
}
