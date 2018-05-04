#!/bin/bash
# alias gopath='export GOPATH=`pwd`'
# go       build & clear
#

export GOPATH=`pwd`
export GO15VENDOREXPERIMENT=1

# Source function library.
if [[ $EUID -ne 0 ]]; then
	echo Only root can access run
	exit 1
fi
rc=0

SRC_PATH=wishlily.coding

case "$1" in
	restore)
		cd src/$SRC_PATH/
		godep restore
		cd ../..
	;;
	build)
		go vet $SRC_PATH/...
		#go build -o xxx $SRC_PATH/xxx
	;;
	test)
		gotestcover -v -coverprofile=cover.out $SRC_PATH/...
	;;
	cover)
		go tool cover -html=cover.out -o coverage.html
	;;
	clean)
		rm -f *.log *.out *.html *.gz *.err bin/*
	;;
	*)
	echo $"Usage: $0 {build|test|cover|clean}"
	exit 2
esac

exit $rc