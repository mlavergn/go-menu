###############################################
#
# Makefile
#
###############################################

.DEFAULT_GOAL := build

VERSION := 1.0.0

#
# Build settings
#
GOOS = darwin
GOARCH = amd64

build:
	GOOS=${GOOS} GOARCH=${GOARCH} CC=clang go build -o menu main

clean:
	-rm -f menu

run:
	./menu

st:
	open -a SourceTree .
