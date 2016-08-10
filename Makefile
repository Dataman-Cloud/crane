.PHONY: build doc fmt lint run test vet test-cover-html test-cover-func collect-cover-data

OS := $(shell uname)
ifeq ($(OS),Darwin)
	BUILD_OPTS=
else
	BUILD_OPTS=CGO_ENABLED=0 GOOS=linux GOARCH=amd64
endif

# Prepend our vendor directory to the system GOPATH
# so that import path resolution will prioritize
# our third party snapshots.
export GO15VENDOREXPERIMENT=1
# GOPATH := ${PWD}/vendor:${GOPATH}
# export GOPATH

default: build

build: fmt
	 ${BUILD_OPTS} go build -ldflags "-X version.BuildTime `date -u +.%Y%m%d.%H%M%S` -X version.Version 0.1-`git rev-parse --short HEAD`" -v -o ./bin/rolex ./src/

rel: fmt
	${BUILD_OPTS} go build -v -o ../rel/rolex ./src/

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
	./bin/rolex

test:
	go test ./src/...

# http://godoc.org/code.google.com/p/go.tools/cmd/vet
# go get code.google.com/p/go.tools/cmd/vet
vet:
#	go vet ./src/...
#

clean:
	rm -rf ../bin/* ../rel/*

PACKAGES = $(shell find ./src/ -type d -not -path '*/\.*')
collect-cover-data:
	echo "mode: count" > coverage-all.out
	@$(foreach pkg,$(PACKAGES),\
		if grep -Fxq "$(pkg)" dirs_to_ignore_test.txt; then\
			echo "Skipping parent dir $(pkg)";\
                else\
		       go test -v -coverprofile=coverage.out -covermode=count $(pkg) || exit $$?;\
		       if [ -f coverage.out ]; then\
		           tail -n +2 coverage.out >> coverage-all.out;\
                       fi\
		fi;)

test-cover-html:
	go tool cover -html=coverage-all.out -o coverage.html

test-cover-func:
	go tool cover -func=coverage-all.out
