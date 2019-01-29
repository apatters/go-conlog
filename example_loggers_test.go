// Copyright 2019 Secure64 Software Corporation. All rights reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package conlog_test

// This example illustrates how to use the Loggers object to send
// output to multiple logs. The Loggers object is typically used in a
// command-line program to send output to both the TTY/console and to
// a log file with one call.

import (
	_ "fmt"
	"io/ioutil"
	"os"

	"github.com/apatters/go-conlog"
)

func ExampleLoggers() {
	// Initialize the log going to the TTY/console.
	ttyLog := conlog.NewLogger()
	ttyLog.SetLevel(conlog.InfoLevel)
	formatter := conlog.NewStdFormatter()
	formatter.Options.LogLevelFmt = conlog.LogLevelFormatLongLower
	formatter.Options.ShowLogLevelColors = true
	ttyLog.SetFormatter(formatter)

	// Initialize the log going to a file.
	logFile, _ := ioutil.TempFile("", "mylogfile-")
	defer os.Remove(logFile.Name()) // You normally wouldn't delete the file.
	fileLog := conlog.NewLogger()
	fileLog.SetLevel(conlog.DebugLevel)
	fileLog.SetOutput(logFile)
	fileLog.SetErrorOutput(logFile)
	formatter = conlog.NewStdFormatter()
	formatter.Options.LogLevelFmt = conlog.LogLevelFormatShort
	formatter.Options.TimestampType = conlog.TimestampTypeWall
	formatter.Options.WallclockTimestampFmt = "15:04:05"
	fileLog.SetFormatter(formatter)

	// Initialize the multi-logger. We will use this one when we
	// want output to go to both logs.
	bothLogs := conlog.NewLoggers(ttyLog, fileLog)

	// We can send messages to individual logs or to both logs.
	ttyLog.Info("This info message only goes to the TTY.")
	fileLog.Info("This info message only goes to the log file.")
	bothLogs.Info("This info message goes to both the TTY and the log file.")

	// Individual log levels are honored.
	bothLogs.Info("This message goes to both logs because they both have log levels >= InfoLevel.")
	bothLogs.Debug("This message only goes the log file because its log level is DebugLevel.")

	// We can use the logs.SetLevel method to set the log level for all logs.
	bothLogs.SetLevel(conlog.DebugLevel)
	bothLogs.Debug("This debug message goes to both the TTY and the log file now that they are DebugLevel.")

	// We can enable/disable Print* methods for all logs using SetPrintEnabled.
	bothLogs.SetPrintEnabled(false)
	bothLogs.Print("This message is suppressed on both the TTY and the log file.")
	bothLogs.SetPrintEnabled(true)
	bothLogs.Print("This message goes to both TTY and the log file now that Print is re-enabled.")

	// We can use Fatal* and Panic* methods, but the messages will
	// only go the first log as the program will terminate before
	// getting to subsequent logs.
	var err error
	if err != nil {
		bothLogs.Fatal("This fatal message only goes to the TTY.")
	}
	var impossibleCond bool
	if impossibleCond {
		bothLogs.Panic("This panic message only goes to the TTY. The panic() output always goes to stderr.")
	}

	// Output:
	// info This info message only goes to the TTY.
	// info This info message goes to both the TTY and the log file.
	// info This message goes to both logs because they both have log levels >= InfoLevel.
	// debug This debug message goes to both the TTY and the log file now that they are DebugLevel.
	// This message goes to both TTY and the log file now that Print is re-enabled.

}
