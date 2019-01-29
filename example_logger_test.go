// Copyright 2019 Secure64 Software Corporation. All rights reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package conlog_test

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/apatters/go-conlog"
)

func ExampleLogger() {
	// Initilize basic logging using the default constructor.
	log := conlog.NewLogger()

	// Messages are output/not output depending on log level. The
	// log level (in ascending order are:
	//
	// PanicLevel
	// FatalLevel
	// ErrorLevel
	// WarnLevel
	// InfoLevel
	// DebugLevel
	log.SetLevel(conlog.DebugLevel)
	log.SetErrorOutput(os.Stdout) // All output goes to stdout.
	formatter := conlog.NewStdFormatter()
	formatter.Options.LogLevelFmt = conlog.LogLevelFormatLongLower
	log.SetFormatter(formatter)
	// log.Panic("This is a panic message.")
	// log.Fatal("This is a fatal message.")
	log.Error("This is an error message.")
	log.Warn("This is a warning message.")
	log.Warning("This is also a warning message.")
	log.Info("This is an info message.")
	log.Debug("This is a debug message.")

	// The default log level is Level.Info. You can set a
	// different log level using SetLevel.
	log.SetLevel(conlog.WarnLevel)
	formatter = conlog.NewStdFormatter()
	formatter.Options.LogLevelFmt = conlog.LogLevelFormatLongLower
	log.SetFormatter(formatter)
	log.Info("This message is above the log level so is not output.")
	log.Warning("This message is at the log level so it is output.")
	log.Error("This message is below the log level so it too is output.")

	// There are three forms of output functions for each log
	// level corresponding to the fmt.Print* functions. The
	// log.Error* output functions have a log.Error(),
	// Log.Errorln, and log.Errorf variations corresponding to
	// fmt.Print(), fmt.Println, and fmt.Printf(). They process
	// parameters in the same way as their fmt counterparts,
	// except that a newline is always output.
	log.SetLevel(conlog.DebugLevel)
	log.Infoln("Print a number with a newline:", 4)
	log.Info("Print a number also with a newline (note we have to add a space): ", 4)
	log.Infof("Print a formatted number with a newline: %d", 4)

	// Output is sent to stderr for log levels of PanicLevel,
	// FatalLevel, and ErrorLevel. Output is sent to stdout for
	// log levels above ErrorLevel. You can change this behavior
	// by setting the Writers in log.out and log.errOut. For
	// example, if we want all output to go to stdout, use:
	log.SetErrorOutput(os.Stdout)
	log.Info("This message is going to stdout.")
	log.Error("This message is now also going to stdout.")

	// We can send output to any Writer. For example, to send
	// output for all levels to a file, we can use:
	logFile, _ := ioutil.TempFile("", "mylogfile-")
	defer os.Remove(logFile.Name()) // You normally wouldn't delete the file.
	log.SetOutput(logFile)
	log.SetErrorOutput(logFile)
	log.Info("This message is going to the logfile.")
	log.Error("This message is also going to the logfile.")
	_ = logFile.Close()
	// Dump to stdout so we can see the results.
	contents, _ := ioutil.ReadFile(logFile.Name())
	fmt.Printf("%s", contents)

	// There are a set of Print* output methods that behave
	// exactly like the corresponding fmt.Print* functions except
	// that output goes to the log.out writer. Output can be
	// suppressed using the SetPrintEnabled() method (enabled by
	// default). Note, Print* methods are not in any way governed
	// by the log level.
	//
	// Command-line programs should generally use these Print*
	// functions in favor of the corresponding fmt versions if
	// using this logging system.
	//
	// Unlike the log level output methods, the Print and Printf
	// methods do not output a trailing newline.
	log.SetOutput(os.Stdout)
	log.SetErrorOutput(os.Stdout)
	log.Println("Print a number with an implicit newline:", 4.0)
	log.Print("Print a number with an added newline (note we have to add a space): ", 4.0, "\n")
	log.Printf("Print a formatted number with an added newline: %f\n", 4.0)

	// We can suppress Print* output using the SetPrintEnabled
	// method. This feature can be useful for programs that have a
	// --verbose option. You can disable Print* output by default
	// and then enable it when the verbose flag it set.
	log.SetPrintEnabled(false)
	fmt.Printf("Printing enabled: %t\n", log.GetPrintEnabled())
	log.Printf("This print message is suppressed.\n")
	log.SetPrintEnabled(true)
	fmt.Printf("Printing enabled: %t\n", log.GetPrintEnabled())
	log.Printf("Print message are no longer suppressed.\n")

	// Output:
	// error This is an error message.
	// warning This is a warning message.
	// warning This is also a warning message.
	// info This is an info message.
	// debug This is a debug message.
	// warning This message is at the log level so it is output.
	// error This message is below the log level so it too is output.
	// info Print a number with a newline: 4
	// info Print a number also with a newline (note we have to add a space): 4
	// info Print a formatted number with a newline: 4
	// info This message is going to stdout.
	// error This message is now also going to stdout.
	// info This message is going to the logfile.
	// error This message is also going to the logfile.
	// Print a number with an implicit newline: 4
	// Print a number with an added newline (note we have to add a space): 4
	// Print a formatted number with an added newline: 4.000000
	// Printing enabled: false
	// Printing enabled: true
	// Print message are no longer suppressed.
}
