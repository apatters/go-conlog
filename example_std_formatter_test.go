package conlog_test

import (
	"os"

	"github.com/apatters/go-conlog"
)

func ExampleStdFormatter() {
	// Basic logging using the default constructor.
	log := conlog.NewLogger()
	log.SetErrorOutput(os.Stdout) // All output goes to stdout for this example.
	log.Info("This is an info message without a leader.")
	log.Warning("This is a warning message without a leader.")
	log.Error("This is an error message without a leader.")

	// Basic logging showing log level leaders.
	formatter := conlog.NewStdFormatter()
	formatter.Options.LogLevelFmt = conlog.LogLevelFormatLongTitle
	log.SetFormatter(formatter)
	log.Info("This is an info message with a leader.")
	log.Warning("This is a warning message with a leader.")
	log.Error("This is an error message with a leader.")

	// The leader can be colorized if going to a tty. If output is
	// going to a file or pipe, no color is used.
	log.SetLevel(conlog.DebugLevel)
	formatter = conlog.NewStdFormatter()
	formatter.Options.LogLevelFmt = conlog.LogLevelFormatLongTitle
	formatter.Options.ShowLogLevelColors = true
	log.SetFormatter(formatter)
	log.Debug("Debug messages are blue.")
	log.Info("Info messages are green.")
	log.Warn("Warning messages are yellow.")
	log.Error("Error messages are red.")

	// You can show the traditional logrus log levels using the
	// formatter LogLevelFormatShort option:
	formatter = conlog.NewStdFormatter()
	formatter.Options.LogLevelFmt = conlog.LogLevelFormatShort
	log.SetFormatter(formatter)
	log.Debug("Debug message with a short leader.")
	log.Info("Info message with a short leader.")
	log.Warn("Warning message with a short leader.")
	log.Error("Error message with a short leader.")

	// You can show long form log levels in lower-case using the
	// formatter LogLevelFormatLongLower option:
	formatter = conlog.NewStdFormatter()
	formatter.Options.LogLevelFmt = conlog.LogLevelFormatLongLower
	log.SetFormatter(formatter)
	log.Debug("Debug message with a long, lowercase leader.")
	log.Info("Info message with a long, lowercase leader.")
	log.Warn("Warning message with a long, lowercase leader.")
	log.Error("Error message with a long, lowercase leader.")

	/*
		// You can show time stamps in wall clock time with various
		// formats.
		formatter = conlog.NewStdFormatter()
		formatter.Options.TimestampType = conlog.TimestampTypeWall
		formatter.Options.LogLevelFmt = conlog.LogLevelFormatLongTitle
		log.Formatter = formatter
		log.Info("Info message with wall clock time (default RFC3339 format).")
		formatter = conlog.NewStdFormatter()
		formatter.Options.TimestampType = conlog.TimestampTypeWall
		formatter.Options.LogLevelFmt = conlog.LogLevelFormatLongTitle
		formatter.Options.WallclockTimestampFmt = time.ANSIC
		log.SetFormatter(formatter)
		log.Info("Info message with wall clock time (ANSIC format).")
		formatter = conlog.NewStdFormatter()
		formatter.Options.TimestampType = conlog.TimestampTypeWall
		formatter.Options.LogLevelFmt = conlog.LogLevelFormatLongTitle
		formatter.Options.WallclockTimestampFmt = "Jan _2 15:04:05"
		log.SetFormatter(formatter)
		log.Info("Info message with wall clock time (custom format).")

		// You can show time stamps with elapsed time and in various
		// formats.
		formatter = conlog.NewStdFormatter()
		formatter.Options.TimestampType = conlog.TimestampTypeElapsed
		formatter.Options.LogLevelFmt = conlog.LogLevelFormatLongTitle
		log.SetFormatter(formatter)
		log.Info("Info message with elapsed time (start).")
		time.Sleep(time.Second)
		log.Info("Info message with elapsed time (wait one second).")
		formatter = conlog.NewStdFormatter()
		formatter.Options.TimestampType = conlog.TimestampTypeElapsed
		formatter.Options.LogLevelFmt = conlog.LogLevelFormatLongTitle
		formatter.Options.ElapsedTimestampFmt = "%02d"
		log.SetFormatter(formatter)
		log.Info("Info message with elapsed time with custom format.")
	*/

	// Output:
	// This is an info message without a leader.
	// This is a warning message without a leader.
	// This is an error message without a leader.
	// Info This is an info message with a leader.
	// Warning This is a warning message with a leader.
	// Error This is an error message with a leader.
	// Debug Debug messages are blue.
	// Info Info messages are green.
	// Warning Warning messages are yellow.
	// Error Error messages are red.
	// DEBU Debug message with a short leader.
	// INFO Info message with a short leader.
	// WARN Warning message with a short leader.
	// ERRO Error message with a short leader.
	// debug Debug message with a long, lowercase leader.
	// info Info message with a long, lowercase leader.
	// warning Warning message with a long, lowercase leader.
	// error Error message with a long, lowercase leader.

}
