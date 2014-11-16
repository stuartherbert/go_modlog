// Copyright (c) 2014-present Stuart Herbert
// Released under the 3-clause BSD license
package modlog

import (
	"fmt"
	"io"
	"sync"

	"github.com/stuartherbert/go_options"
)

// OutputWriter is the function that does the final writing to the output
type OutputWriter func(io.Writer, *LogEntry, map[string]string)

func DefaultOutputWriter(out io.Writer, entry *LogEntry, data map[string]string) {
	if len(entry.Module) > 0 {
		fmt.Fprintf(out, "%s|%s|%s: %s", data[FormatTimestamp], data[FormatLogLevel], entry.Module, entry.Message)
	} else {
		fmt.Fprintf(out, "%s|%s|%s", data[FormatTimestamp], data[FormatLogLevel], entry.Message)
	}
}

func StdlibOutputWriter(out io.Writer, entry *LogEntry, data map[string]string) {
	if len(entry.Module) > 0 {
		fmt.Fprintf(out, "%s %s: %s", data[FormatTimestamp], entry.Module, entry.Message)
	} else if len(data[FormatTimestamp]) > 0 {
		fmt.Fprintf(out, "%s %s", data[FormatTimestamp], entry.Message)
	} else {
		fmt.Fprint(out, entry.Message)
	}
}

// LogOutput represents a single log destination
type LogOutput struct {
	Out        io.Writer
	Filters    map[string]LogFilter
	Formatters map[string]LogFormatter
	mu         sync.Mutex
	Writer     OutputWriter
	Options    *options.OptionsStore
}

// NewLogOutput() creates a new LogOutput
func NewLogOutput(out io.Writer, writer OutputWriter) *LogOutput {
	retval := &LogOutput{
		Out:        out,
		Filters:    make(map[string]LogFilter),
		Formatters: make(map[string]LogFormatter),
		Writer:     writer,
		Options:    options.NewOptionsStore(optionsWhitelist),
	}

	return retval
}

func (self *LogOutput) AddFilter(name string, filter LogFilter) *LogOutput {
	self.mu.Lock()
	defer self.mu.Unlock()

	self.Filters[name] = filter

	return self
}

func (self *LogOutput) RemoveFilter(name string) *LogOutput {
	self.mu.Lock()
	defer self.mu.Unlock()

	delete(self.Filters, name)

	return self
}

func (self *LogOutput) AddFormatter(name string, formatter LogFormatter) *LogOutput {
	self.mu.Lock()
	defer self.mu.Unlock()

	self.Formatters[name] = formatter

	return self
}

func (self *LogOutput) RemoveFormatter(name string) *LogOutput {
	self.mu.Lock()
	defer self.mu.Unlock()

	delete(self.Formatters, name)

	return self
}

func (self *LogOutput) SetWriter(writer OutputWriter) *LogOutput {
	self.Writer = writer
	return self
}

func (self *LogOutput) ProcessEntry(logger *Logger, entry *LogEntry) {
	self.mu.Lock()
	defer self.mu.Unlock()

	// does the log entry pass our filters?
	for _, filter := range self.Filters {
		ok := filter(self.Options, entry)
		if !ok {
			// we're done here
			return
		}
	}

	// run things through our formatters to create the extra fields that
	// are wanted
	data := make(map[string]string)
	for name, formatter := range self.Formatters {
		data[name] = formatter(logger, entry)
	}

	// now we need to write the output
	self.Writer(self.Out, entry, data)
}
