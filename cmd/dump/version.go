// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"fmt"
	"runtime"
)

// Application name and version constants.
const (
	AppName         = "dump"
	AppVersionMajor = 0
	AppVersionMinor = 2
)

// Version returns the application version as a string.
func Version() string {
	return fmt.Sprintf("%s %d.%d (Go runtime %s).\nCopyright (c) 2010-2015, Jim Teeuwen\nCopyright (c) 2016, Andreas T Jonsson.",
		AppName, AppVersionMajor, AppVersionMinor, runtime.Version())
}
