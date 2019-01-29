// Copyright 2019 Secure64 Software Corporation. All rights reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package conlog_test

import (
	"os"

	log "github.com/apatters/go-conlog"
)

func ExampleStdLogger() {
	log.Print("Use Print() to print a message.\n")
	log.Println("Use Println to output a message.")
	log.Printf("Use Printf() to output a %s message.\n", "formatted")

	log.SetErrorOutput(os.Stdout) // go test only looks at stdout.
	log.Error("Use Error() to print an error message.")

	// Output:
	// Use Print() to print a message.
	// Use Println to output a message.
	// Use Printf() to output a formatted message.
	// Use Error() to print an error message.
}
