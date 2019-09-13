export GO111MODULE=on

all: imageflux-cli

imageflux-cli: cmd/imageflux-cli/*.go imageflux/*.go
	go build cmd/imageflux-cli/imageflux-cli.go

bundle:
	dep ensure

check:
	go test -test.v ./imageflux

fmt:
	go fmt ./...

imports:
	goimports -w .

clean:
	rm -rf imageflux-cli
