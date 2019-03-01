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

func Test_FatalWithExitCode(t *testing.T) {
	logger, out, errOut := newSimpleLogger(conlog.FatalLevel)

	testVals := []interface{}{"This is a fatal", 1, 2, "abc", "def"}
	cmpStr := "FATA This is a fatal1 2abcdef\n"
	logger.FatalWithExitCode(-1, testVals...)
	t.Logf("test vals = %v", testVals)
	t.Logf("out string =  %q", out.String())
	t.Logf("cmp string =  %q", cmpStr)
	assert.Empty(t, out.String())
	assert.Equal(t, cmpStr, errOut.String())
	out.Reset()
	errOut.Reset()
}

func Test_FatallnWithExitCode(t *testing.T) {
	logger, out, errOut := newSimpleLogger(conlog.FatalLevel)

	testVals := []interface{}{"This is a fatal", 1, 2, "abc", "def"}
	cmpStr := "FATA This is a fatal 1 2 abc def\n"
	logger.FatallnWithExitCode(-1, testVals...)
	t.Logf("test vals = %v", testVals)
	t.Logf("out string =  %q", out.String())
	t.Logf("errOut string =  %q", out.String())
	t.Logf("cmp string =  %q", cmpStr)
	assert.Empty(t, out.String())
	assert.Equal(t, cmpStr, errOut.String())
	out.Reset()
	errOut.Reset()
}

func Test_FatalfWithExitCode(t *testing.T) {
	logger, out, errOut := newSimpleLogger(conlog.FatalLevel)

	testFmt := "%s %d-%d"
	testArgs := []interface{}{
		"formatted fatal error:",
		1,
		2,
	}
	cmpStr := "FATA formatted fatal error: 1-2\n"
	logger.FatalfWithExitCode(-1, testFmt, testArgs...)
	t.Logf("test fmt = %q", testFmt)
	t.Logf("test args = %v", testArgs)
	t.Logf("out string = %q", out.String())
	t.Logf("errOut string =  %q", errOut.String())
	t.Logf("cmp string =  %q", cmpStr)
	assert.Empty(t, out.String())
	assert.Equal(t, cmpStr, errOut.String())
	out.Reset()
	errOut.Reset()
}

func Test_FatalIfError(t *testing.T) {
	logger, out, errOut := newSimpleLogger(conlog.FatalLevel)

	t.Logf("With error:")
	testVals := []interface{}{"This is a fatal", 1, 2, "abc", "def"}
	cmpStr := "FATA This is a fatal1 2abcdef\n"
	logger.FatalIfError(fmt.Errorf("an error"), -1, testVals...)
	t.Logf("test vals = %v", testVals)
	t.Logf("out string =  %q", out.String())
	t.Logf("errOut string =  %q", errOut.String())
	t.Logf("cmp string =  %q", cmpStr)
	assert.Empty(t, out.String())
	assert.Equal(t, cmpStr, errOut.String())
	out.Reset()
	errOut.Reset()

	t.Logf("Without error:")
	logger.FatalIfError(nil, -1, testVals...)
	t.Logf("test vals = %v", testVals)
	t.Logf("out string =  %q", out.String())
	t.Logf("errOut string =  %q", errOut.String())
	assert.Empty(t, out.String())
	assert.Empty(t, errOut.String())
	out.Reset()
	errOut.Reset()
}

func Test_FatallnIfError(t *testing.T) {
	logger, out, errOut := newSimpleLogger(conlog.FatalLevel)

	t.Logf("With error:")
	testVals := []interface{}{"This is a fatal", 1, 2, "abc", "def"}
	cmpStr := "FATA This is a fatal 1 2 abc def\n"
	logger.FatallnIfError(fmt.Errorf("an error"), -1, testVals...)
	t.Logf("test vals = %v", testVals)
	t.Logf("out string =  %q", out.String())
	t.Logf("errOut string =  %q", errOut.String())
	t.Logf("cmp string =  %q", cmpStr)
	assert.Empty(t, out.String())
	assert.Equal(t, cmpStr, errOut.String())
	out.Reset()
	errOut.Reset()

	t.Logf("Without error:")
	logger.FatalIfError(nil, -1, testVals...)
	t.Logf("test vals = %v", testVals)
	t.Logf("out string =  %q", out.String())
	t.Logf("errOut string =  %q", errOut.String())
	assert.Empty(t, out.String())
	assert.Empty(t, errOut.String())
	out.Reset()
	errOut.Reset()
}

func Test_FatalfIfError(t *testing.T) {
	logger, out, errOut := newSimpleLogger(conlog.FatalLevel)

	testFmt := "%s %d-%d"
	testArgs := []interface{}{
		"formatted fatal error:",
		1,
		2,
	}
	cmpStr := "FATA formatted fatal error: 1-2\n"

	t.Logf("With error:")
	logger.FatalfIfError(fmt.Errorf("an error"), -1, testFmt, testArgs...)
	t.Logf("test fmt = %q", testFmt)
	t.Logf("test args = %v", testArgs)
	t.Logf("out string = %q", out.String())
	t.Logf("errOut string =  %q", errOut.String())
	t.Logf("cmp string =  %q", cmpStr)
	assert.Empty(t, out.String())
	assert.Equal(t, cmpStr, errOut.String())
	out.Reset()
	errOut.Reset()

	t.Logf("Without error:")
	t.Logf("test fmt = %q", testFmt)
	t.Logf("test args = %v", testArgs)
	t.Logf("out string = %q", out.String())
	t.Logf("errOut string =  %q", errOut.String())
	assert.Empty(t, out.String())
	assert.Empty(t, errOut.String())
	out.Reset()
	errOut.Reset()
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

func TestLog_PrintExceptions(t *testing.T) {
	logger, out, _ := newSimpleLogger(conlog.DebugLevel)

	// Try a string containing formatting characters.
	testStr := "%s abc %%n %d %t %v"
	cmpStr := testStr
	logger.Print(testStr)
	t.Logf("test string = %q", testStr)
	t.Logf("out string =     %q", out.String())
	t.Logf("cmp string =     %q", cmpStr)
	assert.Equal(t, cmpStr, out.String())
}
