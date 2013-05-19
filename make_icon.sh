#/bin/sh

if [ -z "$GOPATH" ]; then
    echo GOPATH environment variable not set
    exit
fi

if [ ! -e "$GOPATH/bin/2carray" ]; then
    echo "Installing 2carray..."
    go get github.com/cratonica/2carray
    if [ $? -ne 0 ]; then
        echo Failure executing go get github.com/cratonica/2carray
        exit
    fi
fi

if [ -z "$1" ]; then
    echo Please specify a PNG file
    exit
fi

if [ ! -f "$1" ]; then
    echo $1 is not a valid file
    exit
fi    

OUTPUT=platform/shared/icon.h
echo Generating $OUTPUT
cat $1 | 2carray ICON_PNG > $OUTPUT
if [ $? -ne 0 ]; then
    echo Failure generating $OUTPUT
    exit
fi
echo Finished
