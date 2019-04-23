#!/bin/bash

# export exec name
Exec=tesla_calculator
# Setup Release folder
ReleaseFolder=tesla_calculator.AppDir
des=$PWD/$ReleaseFolder
VERSION=$(git rev-parse --short HEAD)
# Start build process
function Build {
go build -ldflags "-r ." -o $des/usr/bin/$Exec
}


# Clean all the Release Folder && Files
function Clean {
if [ -f $Exec ]
then
 rm -rf $Exec
fi

if [ -f $Exec-$(uname -s)-$(uname -m).tar.gz ]
then
 rm -rf $Exec-$(uname -s)-$(uname -m).tar.gz
fi
}


function Pack {
cp /usr/bin/desktop-file-validate $des/usr/bin/
linuxdeployqt $des/usr/share/applications/$Exec.desktop -appimage \
  -executable=$des/usr/bin/desktop-file-validate
#appimagetool  $ReleaseFolder
}

Clean
Build
Pack




