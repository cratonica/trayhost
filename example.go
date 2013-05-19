/*
   Example program for using trayhost
*/
package main

import (
	"fmt"
	"runtime"
)

func main() {
	// TrayLoop must be called on the OS's main thread
	runtime.LockOSThread()

	go func() {
		// Run your application/server code in here. Most likely you will
		// want to start an HTTP server that the user can hit with a browser
		// by clicking the tray icon.

		// Be sure to call this to link the tray icon to the target url
		SetTrayUrl("http://github.com/cratonica/trayhost")
	}()

	// Enter the host system's event loop
	TrayLoop("My Go App")

	// This is only reached once the user chooses the Exit menu item
	fmt.Println("Exiting")
}
