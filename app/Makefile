PROJECT ?= github.com/davyj0nes/k8s-example-app/app
USERNAME ?= davyj0nes
APP = k8s-example-app
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
	$(call blue, "# Running clean...")
	@rm -f ${APP}
	@rm -f tmp.yaml

build: clean
	$(call blue, "# Running build...")
	@dep ensure
	@CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} $(GO-BUILD) -o ${APP}

container: build
	$(call blue, "# Running container...")
	@docker build -t ${USERNAME}/${APP}:${RELEASE} .
	@docker tag ${USERNAME}/${APP}:${RELEASE} ${USERNAME}/${APP}:latest

push: container
	$(call blue, "# Running push...")
	@docker push docker.io/${USERNAME}/${APP}:${RELEASE}

run: container
	$(call yellow, "# Running run...")
	@docker stop ${USERNAME}/${APP}:${RELEASE} || true && docker rm ${USERNAME}/${APP}:${RELEASE} || true
	@docker run --name ${APP} -p ${PORT}:8080 --rm ${USERNAME}/${APP}:${RELEASE}
		
deploy: push
	$(call yellow, "# Running deploy...")
	@for t in $(shell find ./kubernetes -type f -name "*.yaml"); do \
        echo ---;  cat $$t | \
        	sed -E "s/.Release/$(RELEASE)/g"; \
    done > tmp.yaml
	@kubectl apply -f tmp.yaml
	@rm -f tmp.yaml

test:
	$(call yellow, "# Running test...")
	@go test -v -race $(GO-FILES)

define blue
	@tput setaf 4
	@echo $1
	@tput sgr0
endef

define yellow
	@tput setaf 3 
	@echo $1
	@tput sgr0
endef