fmt:
	go fmt ./...

build:
	go build

build_linux:
	GOOS=linux GOARCH=amd64 go build

build_darwin:
	GOOS=darwin go build

build_windows:
	GOOS=windows go build
	