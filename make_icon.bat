@ECHO OFF
REM My DOS skills are horrid. Here comes the spaghetti...

IF "%GOPATH%"=="" GOTO NOGO
IF NOT EXIST %GOPATH%\bin\2carray.exe GOTO INSTALL
:POSTINSTALL
IF "%1"=="" GOTO NOICO
IF NOT EXIST %1 GOTO BADFILE
ECHO Creating platform\windows\icon.h
TYPE %1 | %GOPATH%\bin\2carray ICON_ICO > platform\windows\icon.h
GOTO DONE

:CREATEFAIL
ECHO Unable to create output file
GOTO DONE

:INSTALL
ECHO Installing 2carray...
go get github.com/cratonica/2carray
IF ERRORLEVEL 1 GOTO GETFAIL
GOTO POSTINSTALL

:GETFAIL
ECHO Failure running go get github.com/cratonica/2carray.  Ensure that go and git are in PATH
GOTO DONE

:NOGO
ECHO GOPATH environment variable not set
GOTO DONE

:NOICO
ECHO Please specify a .ico file
GOTO DONE

:BADFILE
ECHO %1 is not a valid file
GOTO DONE

:DONE

