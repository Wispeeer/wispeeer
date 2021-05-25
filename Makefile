NAME    := wispeeer
SOURCE  := cmd/${NAME}/main.go
BINARY  := bin/${NAME}

# This version-strategy uses git tags to set the version string
VERSION := $(shell git describe --tags --always --dirty="-dev")
UPDATE  := $(shell date +"%Y.%m.%d %X")

#native host
NATIVEOS   := $(shell go env GOOS)
NATIVEARCH := $(shell go env GOARCH)

all: install

define gobuild
    @echo "making ${BINARY}_$(1)_$(2)"
    @mkdir -p bin
    @echo "Wispeeer Version:${VERSION}\nWispeeer Last Update:${UPDATE}"
    @GOOS=$(1) GOARCH=$(2) go build -ldflags "                  \
        -installsuffix 'static'                                 \
        -s -w                                                   \
        -X '$(shell go list -m)/pkg/version.VERSION=${VERSION}' \
        -X '$(shell go list -m)/pkg/version.UPDATE=${UPDATE}'   \
        "                                                       \
        -o ${BINARY}_$(1)_$(2) ${SOURCE}
endef

.PHONY: build
build:          ## build this app.
	$(call gobuild,${NATIVEOS},${NATIVEARCH})

.PHONY: install
install:        ## install this app.
	@echo "${NAME} installing ..."
	@echo "Wispeeer Version:${VERSION}\nWispeeer Last Update:${UPDATE}"
	@go install -ldflags "                                 		\
        -installsuffix 'static'                                 \
        -s -w                                                   \
        -X '$(shell go list -m)/pkg/version.VERSION=${VERSION}' \
        -X '$(shell go list -m)/pkg/version.UPDATE=${UPDATE}'   \
        "                                                       \
        ./...

.PHONY: release
release:        ## build for Multi-platform (go tool dist list).
	$(call gobuild,linux,amd64)
	$(call gobuild,darwin,amd64)
	$(call gobuild,windows,amd64)
	$(call gobuild,linux,arm)
	$(call gobuild,linux,arm64)
	$(call gobuild,darwin,arm64)

.PHONY: help
help:           ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

.PHONY: version
version:        ## Show the app version.
	@echo "Version :$(VERSION)"
	@echo "Last Update :$(UPDATE)"

.PHONY: clean
clean:          ## Clean build cache.
	@rm -rf bin
	@echo "clean [ ok ]"
