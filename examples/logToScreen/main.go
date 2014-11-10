// Copyright (c) 2014-present Stuart Herbert
// Released under a 3-clause BSD license
//
// This simple example shows you what the default output looks like
package main

import (
	log "github.com/stuartherbert/go_modlog"
)

func main() {
	log.Traceln("entering")
	log.Debugln("this is a debug message")
	log.Infoln("this is an info message")
	log.Noticeln("this is a notice")
	log.Warnln("this is a warning")
	log.Errorln("this is an error")
	log.Alertln("this is an alert")
	log.Criticalln("things are going critical now")
	log.Emergencyln("this is a state of emergency")

	// now we hide some of the messages
	log.SetOptions(log.SetMinLogLevel(log.InfoLevel))
	log.Traceln("you shouldn't be able to see this")
	log.Debugln("you shouldn't be able to see this")
	log.Infoln("you should be able to see this")
}
