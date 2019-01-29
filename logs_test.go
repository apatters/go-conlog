// Copyright 2019 Secure64 Software Corporation. All rights reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package conlog_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/apatters/go-conlog"
	"github.com/stretchr/testify/assert"
)

const (
	numLoggers = 2
)

func newSimpleLoggers(level conlog.Level) ([]conlog.ConLogger, []*bytes.Buffer) {
	formatter := conlog.NewStdFormatter()
	formatter.Options.LogLevelFmt = conlog.LogLevelFormatShort

	loggers := make([]conlog.ConLogger, numLoggers)
	outs := make([]*bytes.Buffer, numLoggers)
	for i := 0; i < numLoggers; i++ {
		buf := make([]byte, 0, 256)
		out := bytes.NewBuffer(buf)
		outs[i] = out

		logger := conlog.NewLogger()
		logger.SetOutput(out)
		logger.SetErrorOutput(out)
		logger.SetFormatter(formatter)
		logger.SetLevel(level)
		loggers[i] = conlog.ConLogger(logger)
	}

	return loggers, outs
}

func TestLoggersPrintStyle(t *testing.T) {
	loggersList, outs := newSimpleLoggers(conlog.DebugLevel)

	loggers := conlog.NewLoggers(loggersList...)
	var tests = []testPrint{
		{"Debug", loggers.Debug, conlog.DebugLevel},
		{"Debugln", loggers.Debugln, conlog.DebugLevel},
		{"Info", loggers.Info, conlog.InfoLevel},
		{"Infoln", loggers.Infoln, conlog.InfoLevel},
		{"Warn", loggers.Warn, conlog.WarnLevel},
		{"Warnln", loggers.Warnln, conlog.WarnLevel},
		{"Warning", loggers.Warning, conlog.WarnLevel},
		{"Warningln", loggers.Warningln, conlog.WarnLevel},
		{"Error", loggers.Error, conlog.ErrorLevel},
		{"Errorln", loggers.Errorln, conlog.ErrorLevel},
	}

	for _, test := range tests {
		loggers.SetLevel(test.Level)
		testStr := fmt.Sprintf("%s test", strings.Title(test.Level.String()))
		cmpStr := fmt.Sprintf("%s %s\n", strings.ToUpper(test.Level.String())[0:4], testStr)
		test.Fn(testStr)
		for _, out := range outs {
			t.Logf("func: %s", test.FnName)
			t.Logf("test string = %q", testStr)
			t.Logf("out string  = %q", out.String())
			t.Logf("cmp string  = %q", cmpStr)
			assert.Equal(t, cmpStr, out.String())
			out.Reset()
		}
	}
}

func TestLoggers_Print(t *testing.T) {
	const level = conlog.PanicLevel
	loggerList, outs := newSimpleLoggers(level)

	loggers := conlog.NewLoggers(loggerList...)

	testStr := "Print test"
	cmpStr := testStr
	loggers.Print(testStr)
	for i := 0; i < numLoggers; i++ {
		t.Logf("testStr  = %q", testStr)
		t.Logf("outs[%d] = %q", i, outs[i].String())
		t.Logf("cmpStr   = %q", cmpStr)
		assert.Equal(t, cmpStr, outs[i].String())
	}
}

func TestLoggers_SetPrintEnabled(t *testing.T) {
	const level = conlog.PanicLevel
	loggerList, outs := newSimpleLoggers(level)

	loggers := conlog.NewLoggers(loggerList...)

	testStrEnabled := "Print test enabled"
	cmpStr := testStrEnabled
	loggers.Print(testStrEnabled)
	for i := 0; i < numLoggers; i++ {
		t.Logf("testStrEnabled  = %q", testStrEnabled)
		t.Logf("outs[%d]        = %q", i, outs[i].String())
		t.Logf("cmpStr          = %q", cmpStr)
		assert.Equal(t, cmpStr, outs[i].String())
	}

	loggers.SetPrintEnabled(false)
	testStrDisabled := "Print test disabled"
	cmpStr = testStrEnabled
	loggers.Print(testStrEnabled)
	for i := 0; i < numLoggers; i++ {
		if outs[i].String() != cmpStr {
			t.Logf("testStrDisnabled = %q", testStrDisabled)
			t.Logf("outs[%d]         = %q", i, outs[i].String())
			t.Logf("cmpStr           = %q", cmpStr)
			assert.Equal(t, cmpStr, outs[i].String())
		}
	}

	loggers.SetPrintEnabled(true)
	testStrReenabled := "Print test reenabled"
	cmpStr = testStrEnabled + testStrReenabled
	loggers.Print(testStrReenabled)
	for i := 0; i < numLoggers; i++ {
		t.Logf("testStrReenabled = %q", testStrReenabled)
		t.Logf("outs[%d]         = %q", i, outs[i].String())
		t.Logf("cmpStr           = %q", cmpStr)
		assert.Equal(t, cmpStr, outs[i].String())
	}
}

func TestLoggers_Panic(t *testing.T) {
	const level conlog.Level = conlog.PanicLevel
	loggerList, outs := newSimpleLoggers(level)
	loggers := conlog.NewLoggers(loggerList...)

	defer func() {
		if r := recover(); r == nil {
			t.Log("Panic recovery failed.")
			t.Fail()
		}
	}()

	testStr := fmt.Sprintf("%s test", level)
	cmpStr := fmt.Sprintf("level=%s msg=\"%s\"\n", level.String(), testStr)
	loggers.Panic(testStr)
	// Never get here.
	for i := 0; i < numLoggers; i++ {
		t.Logf("testStr  = %q", testStr)
		t.Logf("outs[%d] = %q", i, outs[i].String())
		t.Logf("cmpStr   = %q", cmpStr)
		assert.Equal(t, cmpStr, outs[i].String())
	}
}
