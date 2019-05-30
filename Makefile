PROJECT?=github.com/gpoulter/upgraded-broccoli
VERSION?=0.0.1

COMMIT := git-$(shell git rev-parse --short HEAD)
BUILD_TIME := $(shell date -u '+%Y-%m-%d_%H:%M:%S')

test:
	go test --race ./...

build:
	go build \
		-ldflags "-s -w -X ${PROJECT}/internal/diagnostics.Version=${VERSION} \
		-X ${PROJECT}/internal/diagnostics.Commit=${COMMIT} \
		-X ${PROJECT}/internal/diagnostics.BuildTime=${BUILD_TIME}" \
		-o bin/tenerife ${PROJECT}/cmd/tenerife