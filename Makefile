default: build
BUILD_DATE=$(shell date +'%y.%m.%d %H:%M:%S')

clean:
	rm -f ./basstest

build:
	CGO_ENABLED=1 \
	go build -o basstest *.go

build-linux:
	CGO_ENABLED=1 \
	GOARCH=amd64 \
	GOOS=linux \
	go build -ldflags "-X 'main.Version=${DRONE_TAG} SHA:${DRONE_COMMIT_SHA}' -X 'main.BuildDate=${BUILD_DATE}'" -o basstest *.go

deps:
	go install github.com/xlab/c-for-go@latest

generate:
	c-for-go ./gobass.yaml

test:
	go test ./
