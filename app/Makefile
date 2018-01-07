PROJECT ?= github.com/davyj0nes/k8s-example-app/app
USERNAME ?= davyj0nes
APP ?= k8s-example-app
PORT ?= 8080

GOOS ?= linux
GOARCH ?= amd64

RELEASE ?= 0.0.3
COMMIT ?= $(shell git rev-parse --short HEAD)
BUILD_TIME ?= $(shell date -u '+%Y-%m-%d_%H:%M:%S')

LDFLAGS = -ldflags "-s -w -X ${PROJECT}/version.Release=${RELEASE} -X ${PROJECT}/version.Commit=${COMMIT} -X ${PROJECT}/version.BuildTime=${BUILD_TIME}" 

GO-BUILD = go build $(LDFLAGS)

GO-FILES = $(shell go list ./... | grep -v /vendor/)

clean:
	@rm -f ${APP}

build: clean
	@dep ensure
	@CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} $(GO-BUILD) -o ${APP}

container: build
	@docker build -t ${USERNAME}/${APP}:${RELEASE} .
	@docker tag ${USERNAME}/${APP}:${RELEASE} ${USERNAME}/${APP}:latest

push: container
	@docker push docker.io/${USERNAME}/${APP}:${RELEASE}

run: container
	@docker stop ${USERNAME}/${APP}:${RELEASE} || true && docker rm ${USERNAME}/${APP}:${RELEASE} || true
	@docker run --name ${APP} -p ${PORT}:8080 --rm ${USERNAME}/${APP}:${RELEASE}
		

test:
	@go test -v -race $(GO-FILES)
