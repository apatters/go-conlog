// Copyright 2019 Secure64 Software Corporation. All rights reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package conlog_test

import (
	"github.com/apatters/go-conlog"
)

// The Fatal* methods are used to make logging the output of fatal
// conditions consistent with the rest of the conlog logging
// system. These methods should be only be used in a "main" package
// and never be buried in other packages (you should be returning
// errors there instead).
//
// The *Fatal* routines use a panic/recover mechanism to exit with a
// specific exit code. This mechanism requires that we call the
// HandleExit function as the last function in main() if any Fatal*
// method is used. This is usually best done by creating a "deferred"
// call at the beginning of main() before any other deferred calls.

func ExampleLogger_Fatal() {
	// The exit routines use a panic/recover mechanism to exit
	// with a specific exit code. We need to call the recovery
	// routine in a defer as the first defer in main() so that it
	// gets called last.
	defer conlog.HandleExit()

	// Initilize basic logging using the default constructor.
	log := conlog.NewLogger()

	var err error

	// The Fatalln, Fatal, and Fatalf methods output a message and
	// exit with exit code 1. These calls append a newline to the
	// message.
	if err != nil {
		log.Fatalln("Fatal message exiting with exit code (note we have added a space): ", 1)
	}
	if err != nil {
		log.Fatal("Fatal message exiting with exit code:", 1)
	}
	if err != nil {
		log.Fatalf("Fatal message exiting with exit code %d", 1)
	}
}
