export GOARCH=arm
export GOARM=7
export GOOS=linux
all:
	go build -v

test:
	go test -v

clean:
	go clean
