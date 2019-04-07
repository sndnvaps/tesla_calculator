#!/bin/bash

# export exec name
Exec=tesla_calculator
# Setup Release folder
ReleaseFolder=teslacoil_cal
des=$PWD/$ReleaseFolder
# Start build process
function Build {
go build -ldflags "-r ." -o $Exec
}


# Clean all the Release Folder && Files
function Clean {
if [ -f $Exec ]
then
 rm -rf $Exec
fi
if [ -d $des ]
then
 rm -rf $des
fi

if [ -f $Exec-$(uname -s)-$(uname -m).tar.gz ]
then
 rm -rf $Exec-$(uname -s)-$(uname -m).tar.gz
fi
}

# After build process, Let's copy all the 
function CopyDependFiles {
if [ ! -d $des ]; then
mkdir -p $des/{bin,icon,lib,plugins}
mkdir -p $des/plugins/platforms
fi
deplist=$(ldd $Exec | awk '{if (match($3,"/")){ printf("%s "),$3 } }')
cp $deplist $des/lib/
cp -rp Depends/qt5.11.1/$(uname -s)_$(uname -m)/*.so $des/plugins/platforms/

}

function CopyExec {
if [ -f $Exec ]
then
  cp $Exec $des/bin/
fi
if [ -f $PWD/images/tesla.ico ]; then
cp $PWD/images/tesla.ico $des/icon/
fi
if [ -f $Exec.desktop ]; then
cp $Exec.desktop $des/bin/
fi
if [ -f $Exec.sh ]; then
cp $Exec.sh $des/bin/
fi
}

function Pack {
tar -czvf $Exec-$(uname -s)-$(uname -m).tar.gz $ReleaseFolder
}

Clean
Build
CopyDependFiles
CopyExec
Pack




