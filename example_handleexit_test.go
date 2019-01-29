// Copyright 2019 Secure64 Software Corporation. All rights reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package conlog_test

import (
	"github.com/apatters/go-conlog"
)

// The *Fatal*WithExitCode routines use a panic/recover mechanism to
// exit with a specific exit code. This mechanism requires that we
// call the HandleExit function as the last function in main() if any
// Fatal* method is used. This is usually best done by creating a
// "deferred" call at the beginning of main() before any other
// deferred calls.

// We define an exit "wrapper" to be called in main if we want to
// explicitly exit. We do not have to call this sort of function to
// fall-through exit (with exit code 0) out of main(). You should not
// call this sort of function outside the main package.
func exit(code int) {
	panic(conlog.Exit{code})
}

func ExampleHandleExit() {
	// The exit routines use a panic/recover mechanism to exit
	// with a specific exit code. We need to call the recovery
	// routine in a defer as the first defer in main() so that it
	// gets called last.
	defer conlog.HandleExit()

	// We can exit explicitly with a specified exit code using the
	// exit() function we defined earlier. We don't need to use
	// this call to exit with exit code 0 when falling-through the
	// end of main.
	exit(5)
}
