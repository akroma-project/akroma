.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/akroma-project/akroma/cmd/swarm
	go build -o bin/geth     github.com/akroma-project/akroma/cmd/geth
	go build -o bin/bootnode github.com/akroma-project/akroma/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/akroma-project/akroma/swarm/...
