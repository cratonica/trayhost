TrayHost
========

   __TrayHost__ provides boilerplate for placing a Go application
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

The Interesting Part
----------------------

Because TrayHost is boilerplate (instead of being a library you link in), you will need to fork or copy the code base for use in your project. This primarily due to the way that the code links against the native libraries, and the fact that the platform-specific code is compiled with generated code. I agree with you that this is unfortunate.

Once you've forked the code base, you will find what you need to get started in the __example.go__ file:

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

OSX will throw a runtime error if __TrayLoop__ is called on a child thread, so the first thing you must do is lock the OS thread. Your application code will need to run on a child goroutine. __SetTrayUrl__ can be called lazily if you need to take some time to determine what port you are running on. 

Build Environment
--------------------------
Before continuing, make sure that your GOPATH environment variable is set, and that you have Git and Mercurial installed and that __go__, __git__, and __hg__ are in your PATH.

Cross-compilation is not currently supported, so you will need access to a machine running the platform that you wish to target. 

Generally speaking, make sure that your system is capable of doing [cgo](http://golang.org/doc/articles/c_go_cgo.html) builds.

#### Linux
In addition to the essential GNU build tools, you will need to have the GTK+ 3.0 development headers installed.

#### Windows
To do cgo builds, you will need to install [MinGW](http://www.mingw.org/). In order to prevent the terminal window from appearing when your application runs, you'll need access to a copy of [editbin.exe](http://msdn.microsoft.com/en-us/library/xd3shwhf.aspx) which comes packaged with Microsoft's C/C++ build tools.

#### Mac OSX
__Note__: TrayHost requires __Go 1.1__ when targetting Mac OSX, or linking will fail due to issues with previous versions of Go and Mach-O binaries.

You'll need the "Command Line Tools for Xcode", which can be installed using Xcode. You should be able to run the __cc__ command from a terminal window.

Building
-----------
Once your build environment is configured, try to build the example app by running

    go install

from the base of your project directory. If all goes well, you should be able to run

    $GOPATH/trayhost

and see a new icon appear in the system tray (or dock on OSX). 

#### Disabling the Command Prompt Window on Windows
The [editbin](http://msdn.microsoft.com/en-us/library/xd3shwhf.aspx) tool will allow you to change the subsystem of the output executable so that users won't see the command window while your application is running. The easiest way to do this is to open the Visual Studio Command Prompt from the start menu (or, alternatively, find __vcvarsall.bat__ in your Visual Studio installation directory and CALL it passing the __x86__ argument). Once you are in this environment, issue the command:

    editbin.exe /SUBSYSTEM:WINDOWS path\to\program.exe

Now when you run the program, you won't see a terminal window.

Generating the Tray Icon
------------------------------------
Included in the project is a tool for generating the icon that gets displayed in the system tray. An icon sized 64x64 pixels should suffice, but there aren't any restrictions here as the system will take care of fitting it (just don't get carried away). 

Icons are embedded into the application by generating a C array containing the byte data using the [2carray](http://github.com/cratonica/2carray) tool, which will automatically be installed if it is missing. The generated .h file will be compiled into the output program.

#### Linux/OSX
From the project root, run __make_icon.sh__, followed by the path to a PNG file to use. For example:

    ./make_icon.sh ~/Documents/MyIcon.png

#### Windows
From the project root, run __make_icon.bat__, followed by the path to a Windows ICO file to use. If you need to create an ICO file, the online tool [ConvertICO](http://convertico.com/) can do this painlessly. 

Example:

    ./make_icon.sh ~/Documents/MyIcon.ico
