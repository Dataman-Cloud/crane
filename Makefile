PACKAGES = $(shell go list ./src/...)

.PHONY: build doc fmt lint run test vet test-cover-html test-cover-func collect-cover-data

## OS checking
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
#GOPATH := ${PWD}/vendor:${GOPATH}
# export GOPATH

# Used to populate version variable in main package.
VERSION=$(shell git describe --always --tags)
BUILD_TIME=$(shell date -u +%Y-%m-%d:%H-%M-%S)
GO_LDFLAGS=-ldflags "-X `go list ./src/version`.Version=$(VERSION) -X `go list ./src/version`.BuildTime=$(BUILD_TIME)"

default: build

build: fmt
	@echo "🐳 $@"
	 ${BUILD_OPTS} go build ${GO_LDFLAGS} -v -o ./bin/crane ./src/

rel: fmt
	@echo "🐳 $@"
	${BUILD_OPTS} go build -v -o ../rel/crane ./src/

doc:
	@echo "🐳 $@"
	godoc -http=:6060 -index

# http://golang.org/cmd/go/#hdr-Run_gofmt_on_package_sources
fmt:
	@echo "🐳 $@"
	go fmt ./src/...

# https://github.com/golang/lint
# go get github.com/golang/lint/golint
lint:
	@echo "🐳  $@"
	@test -z "$$(golint ./src/... | tee /dev/stderr)"

run: build
	@echo "🐳 $@"
	./bin/crane

test:
	@echo "🐳 $@"
	go test -cover=true ./src/...

# http://godoc.org/code.google.com/p/go.tools/cmd/vet
# go get code.google.com/p/go.tools/cmd/vet
vet:
#	go vet ./src/...
#

clean:
	@echo "🐳 $@"
	rm -rf ../bin/* ../rel/*

collect-cover-data:
	@echo "🐳 $@"
	@echo "mode: count" > coverage-all.out
	$(foreach pkg,$(PACKAGES),\
		go test -v -coverprofile=coverage.out -covermode=count $(pkg) || exit $$?;\
		if [ -f coverage.out ]; then\
		    tail -n +2 coverage.out >> coverage-all.out;\
                fi\
		;)

test-cover-html:
	@echo "🐳 $@"
	go tool cover -html=coverage-all.out -o coverage.html

test-cover-func:
	@echo "🐳 $@"
	go tool cover -func=coverage-all.out
