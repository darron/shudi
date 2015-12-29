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
	go get -u github.com/coreos/etcd/Godeps/_workspace/src/golang.org/x/net/context
	go get -u github.com/coreos/etcd/client

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

etcd:
	etcd --data-dir `mktemp -d` --force-new-cluster 1>/dev/null &

etcd_kill:
	ps auxwww | grep "[e]tcd.*tmp.*" | cut -d ' ' -f 3 | xargs kill

consul:
	consul agent -data-dir `mktemp -d` -bootstrap -server -bind=127.0.0.1 1>/dev/null &

consul_kill:
	ps auxwww | grep "[c]onsul agent.*tmp.*bind.127.*" | cut -d ' ' -f 3 | xargs kill

test: wercker

wercker: consul etcd
	basht test/tests.bash
