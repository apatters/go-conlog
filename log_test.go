package conlog_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/apatters/go-conlog"
	"github.com/stretchr/testify/assert"
)

// Create a logger to string using the specified log level.
func newSimpleLogger(level conlog.Level) (*conlog.Logger, *bytes.Buffer, *bytes.Buffer) {
	var outBuf = make([]byte, 0, 256)
	var errOutBuf = make([]byte, 0, 256)
	var out = bytes.NewBuffer(outBuf)
	var errOut = bytes.NewBuffer(errOutBuf)

	formatter := conlog.NewStdFormatter()
	formatter.Options.LogLevelFmt = conlog.LogLevelFormatShort
	log := conlog.NewLogger()
	log.SetOutput(out)
	log.SetErrorOutput(errOut)
	log.SetFormatter(formatter)
	log.SetLevel(level)

	return log, out, errOut
}

type testPrint struct {
	FnName string
	Fn     func(...interface{})
	Level  conlog.Level
}

type testPrintf struct {
	FnName string
	Fn     func(string, ...interface{})
	Level  conlog.Level
}

func TestLog_PrintStyle(t *testing.T) {
	logger, out, errOut := newSimpleLogger(conlog.DebugLevel)
	var outTests = []testPrint{
		{"Debug", logger.Debug, conlog.DebugLevel},
		{"Debugln", logger.Debugln, conlog.DebugLevel},
		{"Info", logger.Info, conlog.InfoLevel},
		{"Infoln", logger.Infoln, conlog.InfoLevel},
		{"Warn", logger.Warn, conlog.WarnLevel},
		{"Warnln", logger.Warnln, conlog.WarnLevel},
		{"Warning", logger.Warning, conlog.WarnLevel},
		{"Warningln", logger.Warningln, conlog.WarnLevel},
	}

	var errOutTests = []testPrint{
		{"Error", logger.Error, conlog.ErrorLevel},
		{"Errorln", logger.Errorln, conlog.ErrorLevel},
	}

	for _, test := range outTests {
		logger.SetLevel(test.Level)
		testStr := fmt.Sprintf("%s test", strings.Title(test.Level.String()))
		cmpStr := fmt.Sprintf("%s %s\n", strings.ToUpper(test.Level.String())[0:4], testStr)
		test.Fn(testStr)
		t.Logf("func: %s", test.FnName)
		t.Logf("test string = %q", testStr)
		t.Logf("out string =  %q", out.String())
		t.Logf("cmp string =  %q", cmpStr)
		assert.Equal(t, cmpStr, out.String())
		out.Reset()
	}

	for _, test := range errOutTests {
		logger.SetLevel(test.Level)
		testStr := fmt.Sprintf("%s test", strings.Title(test.Level.String()))
		cmpStr := fmt.Sprintf("%s %s\n", strings.ToUpper(test.Level.String())[0:4], testStr)
		test.Fn(testStr)
		t.Logf("func: %s", test.FnName)
		t.Logf("test string = %q", testStr)
		t.Logf("out string =  %q", errOut.String())
		t.Logf("cmp string =  %q", cmpStr)
		assert.Equal(t, cmpStr, errOut.String())
		errOut.Reset()
	}
}

func TestLog_PrintfStyle(t *testing.T) {
	logger, out, errOut := newSimpleLogger(conlog.DebugLevel)
	var outTests = []testPrintf{
		{"Debugf", logger.Debugf, conlog.DebugLevel},
		{"Infof", logger.Infof, conlog.InfoLevel},
		{"Warnf", logger.Warnf, conlog.WarnLevel},
		{"Warningf", logger.Warningf, conlog.WarnLevel},
	}
	var errOutTests = []testPrintf{
		{"Errorf", logger.Errorf, conlog.ErrorLevel},
	}

	for _, test := range outTests {
		logger.SetLevel(test.Level)
		testStr := fmt.Sprintf("%s test", strings.Title(test.Level.String()))
		cmpStr := fmt.Sprintf("%s formatted %s\n", strings.ToUpper(test.Level.String())[0:4], testStr)
		test.Fn("formatted %s", testStr)
		t.Logf("func: %s", test.FnName)
		t.Logf("test string = %q", testStr)
		t.Logf("out string =  %q", out.String())
		t.Logf("cmp string =  %q", cmpStr)
		assert.Equal(t, cmpStr, out.String())
		out.Reset()
	}
	for _, test := range errOutTests {
		logger.SetLevel(test.Level)
		testStr := fmt.Sprintf("%s test", strings.Title(test.Level.String()))
		cmpStr := fmt.Sprintf("%s formatted %s\n", strings.ToUpper(test.Level.String())[0:4], testStr)
		test.Fn("formatted %s", testStr)
		t.Logf("func: %s", test.FnName)
		t.Logf("test string = %q", testStr)
		t.Logf("out string =  %q", errOut.String())
		t.Logf("cmp string =  %q", cmpStr)
		assert.Equal(t, cmpStr, errOut.String())
		out.Reset()
	}
}

func TestLog_DisablePrint(t *testing.T) {
	logger, out, _ := newSimpleLogger(conlog.DebugLevel)

	enabledTestStr := "Print enabled test"
	cmpStr := enabledTestStr
	logger.Print(enabledTestStr)
	t.Logf("enabled string = %q", enabledTestStr)
	t.Logf("out string =     %q", out.String())
	t.Logf("cmp string =     %q", cmpStr)
	assert.Equal(t, cmpStr, out.String())

	logger.SetPrintEnabled(false)
	disabledTestStr := "Print disabled test"
	logger.Print(disabledTestStr)
	// Note output is cumulative, so we check for previous
	// results.
	cmpStr = enabledTestStr
	t.Logf("out string =      %q", out.String())
	t.Logf("cmp string =      %q", cmpStr)
	t.Logf("disabled string = %q", disabledTestStr)
	assert.Equal(t, cmpStr, out.String())

	logger.SetPrintEnabled(true)
	reenabledTestStr := "Print reenabled test"
	logger.Print(reenabledTestStr)
	// Note output is cumulative, so we check for previous
	// results.
	cmpStr = enabledTestStr + reenabledTestStr
	t.Logf("reenabled string = %q", reenabledTestStr)
	t.Logf("out string =       %q", out.String())
	t.Logf("cmp string =       %q", cmpStr)
	assert.Equal(t, cmpStr, out.String())
}

func TestLog_BasicPanic(t *testing.T) {
	logger, _, errOut := newSimpleLogger(conlog.PanicLevel)

	defer func() {
		if r := recover(); r == nil {
			t.Log("Panic recovery failed.")
			t.Fail()
		}
	}()

	testStr := "Panic test"
	cmpStr := fmt.Sprintf("level=panic msg=\"%s\"\n", testStr)
	t.Logf("testStr    = %q", testStr)
	t.Logf("cmpStr     = %q", cmpStr)
	logger.Panic(testStr)
	// We never get here.
	t.Logf("out string = %q", errOut.String())
	assert.Equal(t, cmpStr, errOut.String())
}
