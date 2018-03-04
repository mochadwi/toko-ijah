#!/bin/sh

#run http server using gin http hot reload
dep ensure
gin -a 7575 run main.go