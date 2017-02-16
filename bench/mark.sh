#!/bin/sh
gcc ctest.c -o ctest && ./ctest
go build main.go && ./main | grep -v '^[A4]'
