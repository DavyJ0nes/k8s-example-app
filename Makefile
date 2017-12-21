PROJECT ?= github.com/davyj0nes/k8s-example-app
USERNAME ?= davyj0nes
APP ?= k8s-example-app
PORT ?= 8080

GOOS ?= linux
GOARCH ?= amd64

RELEASE ?= 0.0.3
COMMIT ?= $(shell git rev-parse --short HEAD)
BUILD_TIME ?= $(shell date -u '+%Y-%m-%d_%H:%M:%S')

clean:
	rm -f ${APP}

build: clean
	dep ensure
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build \
		-ldflags "-s -w \
		-X ${PROJECT}/version.Release=${RELEASE} \
		-X ${PROJECT}/version.Commit=${COMMIT} \
		-X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
		-o ${APP}

run: build
	./${APP} -port ${PORT}

container: build
	docker build -t ${USERNAME}/${APP}:${RELEASE} .
	docker tag ${USERNAME}/${APP}:${RELEASE} ${USERNAME}/${APP}:latest

push: container
	docker push docker.io/${USERNAME}/${APP}:${RELEASE}

run: container
	docker stop ${USERNAME}/${APP}:${RELEASE} || true && docker rm ${USERNAME}/${APP}:${RELEASE} || true
	docker run --name ${APP} -p ${PORT}:${PORT} --rm ${USERNAME}/${APP}:${RELEASE}
		

test:
	go test -v -race ./...
