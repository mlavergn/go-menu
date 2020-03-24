package main

/*
#import <Foundation/Foundation.h>
#import "cmessage.h"
*/
import "C"

import (
	"unsafe"
)

// COperation provides a Go companion to the C Operation enum
type COperation uint32

const (
	none  COperation = 0
	reset            = 10
	demo             = 20
	quit             = 99
)

// CMessage provides a Go companion to the C Message struct
type CMessage struct {
	op         COperation
	text       string
	err        string
	dataString string
	dataInt    int
	dataBool   bool
	cMessage   *C.struct_CMessage
}

// NewCMessage export
func NewCMessage() *CMessage {
	return &CMessage{
		op:         0,
		text:       "",
		dataString: "",
		dataInt:    0,
		dataBool:   false,
	}
}

// Free releases malloc memory used by C
func (id *CMessage) Free() {
	C.free(unsafe.Pointer(id.cMessage.text))
	C.free(unsafe.Pointer(id.cMessage.dataString))
	id.cMessage = nil
}

// Write a Go message struct to a C pointer
func (id *CMessage) Write() *C.struct_CMessage {
	if id.cMessage != nil {
		id.Free()
	}

	id.cMessage = &C.struct_CMessage{
		op:         C.COperation(id.op),
		text:       C.CString(id.text),
		dataString: C.CString(id.dataString),
		dataInt:    C.int(id.dataInt),
		dataBool:   C.bool(id.dataBool),
	}

	return id.cMessage
}

// Read a C message pointer into Go
func (id *CMessage) Read(ptr unsafe.Pointer) *CMessage {
	cmsg := (*C.struct_CMessage)(ptr)

	id.op = COperation(cmsg.op)
	id.text = C.GoString(cmsg.text)
	id.dataString = C.GoString(cmsg.dataString)
	id.dataInt = int(cmsg.dataInt)
	id.dataBool = bool(cmsg.dataBool)

	return id
}

// unused
func (id *CMessage) pointer(size int) unsafe.Pointer {
	return C.malloc(C.sizeof_char * C.ulong(size))
}

// unused
func (id *CMessage) freePointer(ptr unsafe.Pointer) {
	defer C.free(unsafe.Pointer(ptr))
}
