package main

import "C"
import (
	"unsafe"
)

/**
 * Export must be in a file separate from CGO code to avoid duplicate symbols
 */

//export onMessageGo
func onMessageGo(msg unsafe.Pointer) {
	recvMessage(msg)
}
