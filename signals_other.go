//go:build windows || plan9

package main

import (
	"os"
)

var exitSignals = []os.Signal{os.Interrupt}
