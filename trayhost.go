/*
   Trayhost provides boilerplate for placing a Go application
   in the task bar (system tray, notification area, or dock)
   in a consistent manner across multiple platforms. Currently,
   there is built-in support for Windows, Mac OSX, and Linux
   systems that support GTK+ 3 status icons (including
   Gnome 2, KDE 4, Cinnamon, MATE and other desktop
   environments).

   The indended usage is for redistributable web applications
   that require access to the client's system in ways that
   make remotely hosted web apps not possible, but still
   wish to use the web browser as the primary user interface.

   This allows for developing cross-platform applications with
   deep access to the system.

   A tray icon will be placed with menu options to open a
   URL, giving users easy access to the web-based user
   interface. On OSX, the icon will reside in the dock
   per Apple's design guidelines rather than in the
   so-called "Menu Extras" location.

   Further information can be found at the project's
   home at http://github.com/cratonica/trayhost

   Clint Caywood

   http://github.com/cratonica/trayhost
*/
package trayhost

import (
	"reflect"
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

var isExiting bool
var urlPtr unsafe.Pointer

// Run the host system's event loop
func EnterLoop(title string, imageData []byte) {
	defer C.free(urlPtr)

	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	// Copy the image data into unmanaged memory
	cImageData := C.malloc(C.size_t(len(imageData)))
	defer C.free(cImageData)
	var cImageDataSlice []C.uchar
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&cImageDataSlice))
	sliceHeader.Cap = len(imageData)
	sliceHeader.Len = len(imageData)
	sliceHeader.Data = uintptr(cImageData)

	for i, v := range imageData {
		cImageDataSlice[i] = C.uchar(v)
	}

	// Enter the loop
	C.native_loop(cTitle, &cImageDataSlice[0], C.uint(len(imageData)))

	// If reached, user clicked Exit
	isExiting = true
}

// Set the URL that the tray icon will open in a browser
func SetUrl(url string) {
	if isExiting {
		return
	}
	cs := C.CString(url)
	C.free(urlPtr)
	urlPtr = unsafe.Pointer(cs)
	C.set_url(cs)
}
