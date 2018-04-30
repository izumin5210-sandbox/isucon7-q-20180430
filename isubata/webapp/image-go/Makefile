GOPATH := ${PWD}
export GOPATH

build:
	go build -v image-server
	go build -v image-extractor

vet:
	go vet ./src/image-server/...
