GOACH=arm64
GOARM=7
GOOS=linux
all:
	go build -v

test:
	go test -v

clean:
	go clean
