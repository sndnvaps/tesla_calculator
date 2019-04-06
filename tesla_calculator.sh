#!/bin/bash
appname=tesla_calculator
dirname=/usr/local/teslacoil_cal
LD_LIBRARY_PATH=$dirname/lib:$LD_LIBRARY_PATH
export LD_LIBRARY_PATH
$dirname/bin/$appname "$@"
