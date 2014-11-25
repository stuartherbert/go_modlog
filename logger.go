// Copyright (c) 2014-present Stuart Herbert
// Released under the 3-clause BSD license
package modlog

import (
	"fmt"
	"io"
	_ "log"
	"os"
	"sync"

	"github.com/stuartherbert/go_extras/extrafmt"
	"github.com/stuartherbert/go_options"
)

// See https://tools.ietf.org/html/rfc5424 for a list of the official
// log levels
type Logger struct {
	// a list of filters to apply before we send the message out to our
	// list of outputs
	Filters map[string]LogFilter

	// the different outputs to write to
	Outputs map[string]*LogOutput

	// StdlibFlags are the flags also supported by the stdlib's log package
	StdlibFlags int

	// StdlibPrefix is the log message prefix supported by the stdlib's
	// log package
	StdlibPrefix string

	// Settings is a generic databag
	Options *options.OptionsStore

	// avoids race conditions
	mu sync.RWMutex
}

// New() returns a new Logger for you to embed and/or use
//
// it is compatible with the stdlib's log.New() function
func New(out io.Writer, prefix string, flag int) *Logger {
	return NewLogger(
		SetDefaultOutput(out),
		SetStdlibPrefix(prefix),
		SetStdlibFlags(flag),
	)
}

// NewLogger() is a flexible alternative to New()
func NewLogger(logOptions ...LogOption) *Logger {
	// create a new logger
	retval := &Logger{
		Outputs: make(map[string]*LogOutput),
		Filters: make(map[string]LogFilter),
		Options: options.NewOptionsStore(optionsWhitelist),
	}

	retval.SetOptions(logOptions...)

	// make sure that we have at least one output
	if len(retval.Outputs) == 0 {
		retval.createDefaultOutput(os.Stderr)
	}

	// all done
	return retval
}

// createDefaultOutput() creates an output called 'default'
//
// This is called from our New() and NewLogger() functions to ensure that
// there is at least one output to write to
func (self *Logger) createDefaultOutput(out io.Writer) {
	output := self.AddOutput("default", out)
	output.AddFormatter(FormatTimestamp, StdlibDateTimeFormatter).
		AddFormatter(FormatLogLevel, ShortLogLevelFormatter)
}

// SetOptions applies the list of options to the current logger
func (self *Logger) SetOptions(logOptions ...LogOption) {
	// apply any user-provided options
	for _, option := range logOptions {
		err := option(self)
		if err != nil {
			// well that's no good :(
			//
			// as logging is a fundamental component, treat errors here
			// as fatal
			panic(fmt.Sprintf("Unable to set log option; error is: %s\n", err.Error()))
		}
	}
}

func (self *Logger) AddOutput(name string, out io.Writer) *LogOutput {
	self.mu.Lock()
	defer self.mu.Unlock()

	output := NewLogOutput(out, DefaultOutputWriter)

	self.Outputs[name] = output
	return output
}

func (self *Logger) RemoveOutput(name string) {
	self.mu.Lock()
	defer self.mu.Unlock()

	delete(self.Outputs, name)
}

func (self *Logger) GetOutput(name string) *LogOutput {
	self.mu.Lock()
	defer self.mu.Unlock()

	output, _ := self.Outputs[name]
	return output
}

func (self *Logger) AddFilter(name string, filter LogFilter) {
	self.mu.Lock()
	defer self.mu.Unlock()

	self.Filters[name] = filter
}

func (self *Logger) RemoveFilter(name string) {
	self.mu.Lock()
	defer self.mu.Unlock()

	delete(self.Filters, name)
}

func (self *Logger) AddLogEntry(level LogLevel, module string, message string) {
	entry := NewLogEntry(level, module, message)
	self.processEntry(entry)
}

func (self *Logger) processEntry(entry *LogEntry) {
	self.mu.Lock()
	defer self.mu.Unlock()

	// does this entry pass the filters?
	for _, filter := range self.Filters {
		ok := filter(self.Options, entry)
		if !ok {
			// we're done
			return
		}
	}

	// send this out to all of our outputs
	for _, output := range self.Outputs {
		output.ProcessEntry(self, entry)
	}
}

func (self *Logger) Tracef(format string, args ...interface{}) {
	self.AddLogEntry(TraceLevel, "", fmt.Sprintf(format, args...))
}

func (self *Logger) Trace(args ...interface{}) {
	self.AddLogEntry(TraceLevel, "", fmt.Sprint(args...))
}

func (self *Logger) Traceln(args ...interface{}) {
	self.AddLogEntry(TraceLevel, "", extrafmt.Sprintnln(args...))
}

func (self *Logger) Debugf(format string, args ...interface{}) {
	self.AddLogEntry(DebugLevel, "", fmt.Sprintf(format, args...))
}

func (self *Logger) Debug(args ...interface{}) {
	self.AddLogEntry(DebugLevel, "", fmt.Sprint(args...))
}

func (self *Logger) Debugln(args ...interface{}) {
	self.AddLogEntry(DebugLevel, "", extrafmt.Sprintnln(args...))
}

func (self *Logger) Infof(format string, args ...interface{}) {
	self.AddLogEntry(InfoLevel, "", fmt.Sprintf(format, args...))
}

func (self *Logger) Info(args ...interface{}) {
	self.AddLogEntry(InfoLevel, "", fmt.Sprint(args...))
}

func (self *Logger) Infoln(args ...interface{}) {
	self.AddLogEntry(InfoLevel, "", extrafmt.Sprintnln(args...))
}

func (self *Logger) Noticef(format string, args ...interface{}) {
	self.AddLogEntry(NoticeLevel, "", fmt.Sprintf(format, args...))
}

func (self *Logger) Notice(args ...interface{}) {
	self.AddLogEntry(NoticeLevel, "", fmt.Sprint(args...))
}

func (self *Logger) Noticeln(args ...interface{}) {
	self.AddLogEntry(NoticeLevel, "", extrafmt.Sprintnln(args...))
}

func (self *Logger) Warnf(format string, args ...interface{}) {
	self.AddLogEntry(WarnLevel, "", fmt.Sprintf(format, args...))
}

func (self *Logger) Warn(args ...interface{}) {
	self.AddLogEntry(WarnLevel, "", fmt.Sprint(args...))
}

func (self *Logger) Warnln(args ...interface{}) {
	self.AddLogEntry(WarnLevel, "", extrafmt.Sprintnln(args...))
}

func (self *Logger) Errorf(format string, args ...interface{}) {
	self.AddLogEntry(ErrorLevel, "", fmt.Sprintf(format, args...))
}

func (self *Logger) Error(args ...interface{}) {
	self.AddLogEntry(ErrorLevel, "", fmt.Sprint(args...))
}

func (self *Logger) Errorln(args ...interface{}) {
	self.AddLogEntry(ErrorLevel, "", extrafmt.Sprintnln(args...))
}

func (self *Logger) Criticalf(format string, args ...interface{}) {
	self.AddLogEntry(CriticalLevel, "", fmt.Sprintf(format, args...))
}

func (self *Logger) Critical(args ...interface{}) {
	self.AddLogEntry(CriticalLevel, "", fmt.Sprint(args...))
}

func (self *Logger) Criticalln(args ...interface{}) {
	self.AddLogEntry(CriticalLevel, "", extrafmt.Sprintnln(args...))
}

func (self *Logger) Alertf(format string, args ...interface{}) {
	self.AddLogEntry(AlertLevel, "", fmt.Sprintf(format, args...))
}

func (self *Logger) Alert(args ...interface{}) {
	self.AddLogEntry(AlertLevel, "", fmt.Sprint(args...))
}

func (self *Logger) Alertln(args ...interface{}) {
	self.AddLogEntry(AlertLevel, "", extrafmt.Sprintnln(args...))
}

func (self *Logger) Emergencyf(format string, args ...interface{}) {
	self.AddLogEntry(EmergencyLevel, "", fmt.Sprintf(format, args...))
}

func (self *Logger) Emergency(args ...interface{}) {
	self.AddLogEntry(EmergencyLevel, "", fmt.Sprint(args...))
}

func (self *Logger) Emergencyln(args ...interface{}) {
	self.AddLogEntry(EmergencyLevel, "", extrafmt.Sprintnln(args...))
}

func (self *Logger) Fatal(args ...interface{}) {
	self.AddLogEntry(FatalLevel, "", fmt.Sprint(args...))
}

func (self *Logger) Fatalf(format string, args ...interface{}) {
	self.AddLogEntry(FatalLevel, "", fmt.Sprintf(format, args...))
}

func (self *Logger) Fatalln(args ...interface{}) {
	self.AddLogEntry(FatalLevel, "", extrafmt.Sprintnln(args...))
}

func (self *Logger) Panic(args ...interface{}) {
	self.AddLogEntry(PanicLevel, "", fmt.Sprint(args...))
}

func (self *Logger) Panicf(format string, args ...interface{}) {
	self.AddLogEntry(PanicLevel, "", fmt.Sprintf(format, args...))
}

func (self *Logger) Panicln(args ...interface{}) {
	self.AddLogEntry(PanicLevel, "", extrafmt.Sprintnln(args...))
}

func (self *Logger) Print(args ...interface{}) {
	self.AddLogEntry(InfoLevel, "", fmt.Sprint(args...))
}

func (self *Logger) Printf(format string, args ...interface{}) {
	self.AddLogEntry(InfoLevel, "", fmt.Sprintf(format, args...))
}

func (self *Logger) Println(args ...interface{}) {
	self.AddLogEntry(InfoLevel, "", extrafmt.Sprintnln(args...))
}

func (self *Logger) Flags() int {
	self.mu.Lock()
	defer self.mu.Unlock()

	return self.StdlibFlags
}

// SetFlags() allows you to set the flags that are also supported by the
// stdlib's log package
func (self *Logger) SetFlags(flag int) {
	self.mu.Lock()
	defer self.mu.Unlock()

	self.StdlibFlags = flag
}

func (self *Logger) Prefix() string {
	self.mu.RLock()
	defer self.mu.RUnlock()

	return self.StdlibPrefix
}

func (self *Logger) SetPrefix(prefix string) {
	self.mu.Lock()
	defer self.mu.Unlock()

	self.StdlibPrefix = prefix
}

func (self *Logger) Output(calldepth int, s string) error {
	return nil
}

// SetOutput() allows you to set the output to write log messages to
func (self *Logger) SetOutput(out io.Writer) {
	self.AddOutput("default", out).
		AddFormatter(FormatTimestamp, StdlibDateTimeFormatter).
		AddFormatter(FormatModule, StdlibPrefixFormatter).
		AddFormatter(FormatFilename, StdlibFileFormatter).
		SetWriter(StdlibOutputWriter)
}
