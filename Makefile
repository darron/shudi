GIT_COMMIT=$(shell git rev-parse HEAD)
SHUDI_VERSION=$(shell ./version)
COMPILE_DATE=$(shell date -u +%Y%m%d.%H%M%S)
BUILD_FLAGS=-X main.CompileDate=$(COMPILE_DATE) -X main.GitCommit=$(GIT_COMMIT) -X main.Version=$(SHUDI_VERSION)

all: build

deps:
	go get -u github.com/spf13/cobra
	go get -u github.com/PagerDuty/godspeed
	go get -u github.com/hashicorp/consul/api
	go get -u github.com/progrium/basht
	go get -u github.com/CiscoCloud/consul-cli

format:
	gofmt -w .

clean:
	rm -f bin/shudi || true

build: clean
	go build -ldflags "$(BUILD_FLAGS)" -o bin/shudi main.go

gziposx:
	gzip bin/shudi
	mv bin/shudi.gz bin/shudi-$(SHUDI_VERSION)-darwin.gz

linux: clean
	GOOS=linux GOARCH=amd64 go build -ldflags "$(BUILD_FLAGS)" -o bin/shudi main.go

gziplinux:
	gzip bin/shudi
	mv bin/shudi.gz bin/shudi-$(SHUDI_VERSION)-linux-amd64.gz

release: clean build gziposx clean linux gziplinux clean

consul:
	consul agent -data-dir `mktemp -d` -bootstrap -server -bind=127.0.0.1 1>/dev/null &

consul_kill:
	ps auxwww | grep "[c]onsul agent.*tmp.*bind.127.*" | cut -d ' ' -f 3 | xargs kill

test: wercker

wercker: consul
	basht test/tests.bash
