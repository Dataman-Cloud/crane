.PHONY: build doc fmt lint run test vet

# Prepend our vendor directory to the system GOPATH
# so that import path resolution will prioritize
# our third party snapshots.
export GO15VENDOREXPERIMENT=1
# GOPATH := ${PWD}/vendor:${GOPATH}
# export GOPATH

default: build

build: fmt
	go build -ldflags "-X version.BuildTime `date -u +.%Y%m%d.%H%M%S` -X version.Version 0.1-`git rev-parse --short HEAD`" -v -o ../bin/rolex ./src/

rel: fmt
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -o ../rel/rolex ./src/

doc:
	godoc -http=:6060 -index

# http://golang.org/cmd/go/#hdr-Run_gofmt_on_package_sources
fmt:
	go fmt ./src/...

# https://github.com/golang/lint
# go get github.com/golang/lint/golint
lint:
	golint ./src/

run: build
	../bin/rolex

test:
	go test ./src/...

# http://godoc.org/code.google.com/p/go.tools/cmd/vet
# go get code.google.com/p/go.tools/cmd/vet
vet:
#	go vet ./src/...
#

clean:
	rm -rf ../bin/* ../rel/*

