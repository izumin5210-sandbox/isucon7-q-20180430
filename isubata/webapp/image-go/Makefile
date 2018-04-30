GOPATH := ${PWD}
export GOPATH

build:
	go build -v isubata-image
	go build -v image-extractor

vet:
	go vet ./src/isubata-image/...
