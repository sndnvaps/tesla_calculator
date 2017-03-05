#!/bin/sh
goqt_rcc -go main -o tc_qrc.go tc.qrc
go build -ldflags "-r ." -o Tesla_calculator
