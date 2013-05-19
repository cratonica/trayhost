package main

import (
	"unsafe"
)

/*
#cgo linux pkg-config: gtk+-3.0
#cgo linux CFLAGS: -DLINUX
#cgo windows CFLAGS: -DWIN32
#cgo darwin CFLAGS: -DDARWIN -x objective-c
#cgo darwin LDFLAGS: -framework Cocoa
#include <stdlib.h>
#include "platform/platform.h"
*/
import "C"

// Run the host system's event loop
func TrayLoop(title string) {
	cs := C.CString(title)
	C.native_loop(cs)

	// If reached, user clicked Exit
	C.free(unsafe.Pointer(cs))
}

var urlPtr unsafe.Pointer

// Set the URL that the tray icon will open in a browser
func SetTrayUrl(url string) {
	cs := C.CString(url)
	if urlPtr != unsafe.Pointer(nil) {
		C.free(urlPtr)
	}
	urlPtr = unsafe.Pointer(cs)
	C.set_url(cs)
}
