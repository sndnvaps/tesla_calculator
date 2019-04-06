#!/bin/bash

InstallDir=/usr/local/teslacoil_cal
# Remove the Tesla Coil Calculator from the system
function UnInstall {
if [ -d $InstallDir ]
then
  echo "start UnInsatll Process"
  rm -rf $InstallDir
  echo "Done!"
fi
}

UnInstall

