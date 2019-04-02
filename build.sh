#!/bin/sh
go build -ldflags "-r ." -o tesla_calculator main.go mainform.go CoefficientOfCoupling.go PrimaryCoilForm.go SecCoilInfo.go SparkLength.go sphere.go images_bindata.go
