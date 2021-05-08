export GOARCH=arm
export GOARM=7
export GOOS=linux

.PHONY: no
.PHONY: all
.PHONY: test
.PHONY: clean
no:
	go env -w  GO111MODULE=off
all: no
	go build -v -o xblinkt

test:
	go test -v

clean:
	go clean
