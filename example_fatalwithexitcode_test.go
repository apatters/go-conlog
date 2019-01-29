// Copyright 2019 Secure64 Software Corporation. All rights reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package conlog_test

import (
	"github.com/apatters/go-conlog"
)

// The Fatal*WithExitCode methods are used to make logging the output
// of fatal conditions consistent with the rest of the conlog logging
// system. These methods should be only be used in a "main" package
// and never be buried in other packages (you should be returning
// errors there instead).
//
// The *Fatal*WithExitCode routines use a panic/recover mechanism to
// exit with a specific exit code. This mechanism requires that we
// call the HandleExit function as the last function in main() if any
// Fatal* method is used. This is usually best done by creating a
// "deferred" call at the beginning of main() before any other
// deferred calls.

func ExampleLogger_FatalWithExitCode() {
	// The exit routines use a panic/recover mechanism to exit
	// with a specific exit code. We need to call the recovery
	// routine in a defer as the first defer in main() so that it
	// gets called last.
	defer conlog.HandleExit()

	// Initilize basic logging using the default constructor.
	log := conlog.NewLogger()

	var err error

	// The Fatal*WithExitCode methods work like Fatalln, Fatal,
	// and Fatalf except that you can specify the exit code to
	// use.
	if err != nil {
		log.FatallnWithExitCode(2, "Fatal message exiting with exit code (note we have added a space): ", 2)
	}
	if err != nil {
		log.FatalWithExitCode(2, "Fatal message exiting with exit code:", 2)
	}
	if err != nil {
		log.FatalfWithExitCode(2, "Fatal message exiting with exit code %d", 2)
	}

	// The Fatal*IfError methods output a fatal message and exit
	// with a specified exit code if the the err parameter is not
	// nil. They are used as a short-cut for the common:
	//
	//  if err != nil {
	//      log.FatalWithExitCode(...)
	//  }
	//
	// golang programming construct.
	log.FatallnIfError(err, 3, "Fatal message if err != nil exiting with exit code: ", 3)
	log.FatalIfError(err, 3, "Fatal message if err != nil exiting with exit code:", 3)
	log.FatalfIfError(err, 3, "Fatal message if err != nil exiting with exit code %d, err = %s", 3, err)
}
