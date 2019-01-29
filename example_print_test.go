// Copyright 2019 Secure64 Software Corporation. All rights reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package conlog_test

import (
	"fmt"

	"github.com/apatters/go-conlog"
)

// There are a set of Print* output methods that behave exactly like
// the corresponding fmt.Print* functions except that output goes to
// the log.Out writer. Output can be suppressed using the
// SetPrintEnabled() method (enabled by default). Note, Print* methods
// are not in any way governed by the log level.
//
// Command-line programs should generally use these Print* methods in
// favor of the corresponding fmt versions if using this logging
// system.
//
// Unlike the log level output methods, the Print and Printf methods
// do not output a trailing newline.
func ExampleLogger_Print() {
	// Initilize basic logging using the default constructor.
	log := conlog.NewLogger()

	log.Print("A simple message with an added trailing newline.", "\n")
	log.Print("Compress some adjacent strings", "1st", "2nd", 3, "4th", "5th", 6, 7, "\n")
	log.Print("Print a number with an added trailing newline: ", 4.0, "\n")
	log.Println("Print a number with an implicit newline:", 4.0)
	log.Printf("Print a formatted number with an added newline: %f\n", 4.0)

	// We can suppress Print* output using the SetPrintEnabled
	// method. This feature can be useful for programs that have a
	// --verbose option. You can disable Print* output by default
	// and then enable it when the verbose flag it set.
	log.SetPrintEnabled(false)
	fmt.Printf("Printing enabled: %t\n", log.GetPrintEnabled())
	log.Println("This print message is suppressed.")
	log.SetPrintEnabled(true)
	fmt.Printf("Printing enabled: %t\n", log.GetPrintEnabled())
	log.Println("Print* messages are no longer suppressed.")

	// Output:
	// A simple message with an added trailing newline.
	// Compress some adjacent strings1st2nd34th5th6 7
	// Print a number with an added trailing newline: 4
	// Print a number with an implicit newline: 4
	// Print a formatted number with an added newline: 4.000000
	// Printing enabled: false
	// Printing enabled: true
	// Print* messages are no longer suppressed.
}
