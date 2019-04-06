#!/bin/bash
Exec=tesla_calculator

function Install {
if [ -f $Exec-$(uname -s)-$(uname -m).tar.gz ]
then
  tar -xzf $Exec-$(uname -s)-$(uname -m).tar.gz -C /usr/local
else
  echo "You need to get the latest Release from https://github.com/sndnvaps/tesla_calculator/releases"
fi
}

Install
   
