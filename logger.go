// Copyright (c) 2014-present Stuart Herbert
// Released under the 3-clause BSD license
package modlog

import (
	"fmt"
	"io"
	_ "log"
	"os"
	"sync"

	"github.com/stuartherbert/go_options"
)

// See https://tools.ietf.org/html/rfc5424 for a list of the official
// log levels
type Logger struct {
	// a list of filters to apply before we send the message out to our
	// list of outputs
	Filters map[string]LogFilter

	// the different outputs to write to
	Outputs map[string]LogOutput

	// StdlibFlags are the flags also supported by the stdlib's log package
	StdlibFlags int

	// StdlibPrefix is the log message prefix supported by the stdlib's
	// log package
	StdlibPrefix string

	// Settings is a generic databag
	Options *options.OptionsStore

	// avoids race conditions
	mu sync.Mutex
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
		Outputs: make(map[string]LogOutput),
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
	self.AddOutput("default", out)
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

func (self *Logger) AddOutput(name string, out io.Writer) {
	self.mu.Lock()
	defer self.mu.Unlock()

	output := NewLogOutput(out, DefaultOutputWriter)
	output.AddFormatter(TimestampFormatter, StdlibDateTimeFormatter)

	self.Outputs[name] = *output
}

func (self *Logger) RemoveOutput(name string) {
	self.mu.Lock()
	defer self.mu.Unlock()

	delete(self.Outputs, name)
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
	entry := NewLogEntry(TraceLevel, "", fmt.Sprintf(format, args...))
	self.processEntry(entry)
}

func (self *Logger) Trace(args ...interface{}) {
	entry := NewLogEntry(TraceLevel, "", fmt.Sprint(args...))
	self.processEntry(entry)
}

func (self *Logger) Traceln(args ...interface{}) {
	entry := NewLogEntry(TraceLevel, "", fmt.Sprintln(args...))
	self.processEntry(entry)
}

func (self *Logger) Debugf(format string, args ...interface{}) {
	entry := NewLogEntry(DebugLevel, "", fmt.Sprintf(format, args...))
	self.processEntry(entry)
}

func (self *Logger) Debug(args ...interface{}) {
	entry := NewLogEntry(DebugLevel, "", fmt.Sprint(args...))
	self.processEntry(entry)
}

func (self *Logger) Debugln(args ...interface{}) {
	entry := NewLogEntry(DebugLevel, "", fmt.Sprintln(args...))
	self.processEntry(entry)
}

func (self *Logger) Infof(format string, args ...interface{}) {
	entry := NewLogEntry(InfoLevel, "", fmt.Sprintf(format, args...))
	self.processEntry(entry)
}

func (self *Logger) Info(args ...interface{}) {
	entry := NewLogEntry(InfoLevel, "", fmt.Sprint(args...))
	self.processEntry(entry)
}

func (self *Logger) Infoln(args ...interface{}) {
	entry := NewLogEntry(InfoLevel, "", fmt.Sprintln(args...))
	self.processEntry(entry)
}

func (self *Logger) Noticef(format string, args ...interface{}) {
	entry := NewLogEntry(NoticeLevel, "", fmt.Sprintf(format, args...))
	self.processEntry(entry)
}

func (self *Logger) Notice(args ...interface{}) {
	entry := NewLogEntry(NoticeLevel, "", fmt.Sprint(args...))
	self.processEntry(entry)
}

func (self *Logger) Noticeln(args ...interface{}) {
	entry := NewLogEntry(NoticeLevel, "", fmt.Sprintln(args...))
	self.processEntry(entry)
}

func (self *Logger) Warnf(format string, args ...interface{}) {
	entry := NewLogEntry(WarnLevel, "", fmt.Sprintf(format, args...))
	self.processEntry(entry)
}

func (self *Logger) Warn(args ...interface{}) {
	entry := NewLogEntry(WarnLevel, "", fmt.Sprint(args...))
	self.processEntry(entry)
}

func (self *Logger) Warnln(args ...interface{}) {
	entry := NewLogEntry(WarnLevel, "", fmt.Sprintln(args...))
	self.processEntry(entry)
}

func (self *Logger) Errorf(format string, args ...interface{}) {
	entry := NewLogEntry(ErrorLevel, "", fmt.Sprintf(format, args...))
	self.processEntry(entry)
}

func (self *Logger) Error(args ...interface{}) {
	entry := NewLogEntry(ErrorLevel, "", fmt.Sprint(args...))
	self.processEntry(entry)
}

func (self *Logger) Errorln(args ...interface{}) {
	entry := NewLogEntry(ErrorLevel, "", fmt.Sprintln(args...))
	self.processEntry(entry)
}

func (self *Logger) Criticalf(format string, args ...interface{}) {
	entry := NewLogEntry(CriticalLevel, "", fmt.Sprintf(format, args...))
	self.processEntry(entry)
}

func (self *Logger) Critical(args ...interface{}) {
	entry := NewLogEntry(CriticalLevel, "", fmt.Sprint(args...))
	self.processEntry(entry)
}

func (self *Logger) Criticalln(args ...interface{}) {
	entry := NewLogEntry(CriticalLevel, "", fmt.Sprintln(args...))
	self.processEntry(entry)
}

func (self *Logger) Alertf(format string, args ...interface{}) {
	entry := NewLogEntry(AlertLevel, "", fmt.Sprintf(format, args...))
	self.processEntry(entry)
}

func (self *Logger) Alert(args ...interface{}) {
	entry := NewLogEntry(AlertLevel, "", fmt.Sprint(args...))
	self.processEntry(entry)

}

func (self *Logger) Alertln(args ...interface{}) {
	entry := NewLogEntry(AlertLevel, "", fmt.Sprintln(args...))
	self.processEntry(entry)
}

func (self *Logger) Emergencyf(format string, args ...interface{}) {
	entry := NewLogEntry(EmergencyLevel, "", fmt.Sprintf(format, args...))
	self.processEntry(entry)
}

func (self *Logger) Emergency(args ...interface{}) {
	entry := NewLogEntry(EmergencyLevel, "", fmt.Sprint(args...))
	self.processEntry(entry)

}
func (self *Logger) Emergencyln(args ...interface{}) {
	entry := NewLogEntry(EmergencyLevel, "", fmt.Sprintln(args...))
	self.processEntry(entry)
}

func (self *Logger) Fatal(args ...interface{}) {
	entry := NewLogEntry(FatalLevel, "", fmt.Sprint(args...))
	self.processEntry(entry)
}

func (self *Logger) Fatalf(format string, args ...interface{}) {
	entry := NewLogEntry(FatalLevel, "", fmt.Sprintf(format, args...))
	self.processEntry(entry)

}

func (self *Logger) Fatalln(args ...interface{}) {
	entry := NewLogEntry(FatalLevel, "", fmt.Sprintln(args...))
	self.processEntry(entry)

}

func (self *Logger) Flags() int {
	self.mu.Lock()
	defer self.mu.Unlock()

	return self.StdlibFlags
}

func (self *Logger) Output(calldepth int, s string) error {
	return nil
}
func (self *Logger) Panic(args ...interface{}) {
	entry := NewLogEntry(PanicLevel, "", fmt.Sprint(args...))
	self.processEntry(entry)
}

func (self *Logger) Panicf(format string, args ...interface{}) {
	entry := NewLogEntry(PanicLevel, "", fmt.Sprintf(format, args...))
	self.processEntry(entry)
}

func (self *Logger) Panicln(args ...interface{}) {
	entry := NewLogEntry(PanicLevel, "", fmt.Sprintln(args...))
	self.processEntry(entry)
}

func (self *Logger) Prefix() string {
	self.mu.Lock()
	defer self.mu.Unlock()

	return self.StdlibPrefix
}

func (self *Logger) Print(args ...interface{}) {
}

func (self *Logger) Printf(format string, args ...interface{}) {
	entry := NewLogEntry(InfoLevel, "", fmt.Sprintf(format, args...))
	self.processEntry(entry)
}

func (self *Logger) Println(v ...interface{}) {

}

func (self *Logger) SetFlags(flag int) {
	self.mu.Lock()
	defer self.mu.Unlock()

	self.StdlibFlags = flag
}

func (self *Logger) SetPrefix(prefix string) {
	self.mu.Lock()
	defer self.mu.Unlock()

	self.StdlibPrefix = prefix
}
