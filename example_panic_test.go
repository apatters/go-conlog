// Copyright 2019 Secure64 Software Corporation. All rights reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package conlog_test

import (
	"github.com/apatters/go-conlog"
)

// The Panic* methods are used to make logging the output of failed
// assertions consistent with the rest of the conlog logging system.
func ExampleLogger_Panic() {
	// The Panicln, Panic, and Panicf methods output a panic level
	// message and then call the standard builtin panic() function
	// with the message. The panic message goes to the the ErrOut
	// Writer, but the underlying panic output will always go to
	// stderr.

	// Initilize basic logging using the default constructor.
	log := conlog.NewLogger()

	var impossibleCond bool
	if impossibleCond {
		log.Panicln("Panic message to log with panic output to stderr: ", impossibleCond)
	}
	if impossibleCond {
		log.Panic("Panic message to log with panic output to stderr:", impossibleCond)
	}
	if impossibleCond {
		log.Panicf("Panic message to log with panic output to stderr: %t", impossibleCond)
	}
}
