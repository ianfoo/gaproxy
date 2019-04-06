NAME := gaproxy
USER := ianfoo
PKG ?= github.com/$(USER)/$(NAME)

BIN_DIR ?= bin
BIN ?= $(BIN_DIR)/$(NAME)
IMAGE_TAG ?= $(USER)/$(NAME)

HTTP_PORT ?= 5050
GRPC_PORT ?= 5040

GO_FILES := $(shell find . -name '*.go')

# Set IMAGE_TAG to whatever you like to override default tag.
# e.g.,
# $ IMAGE_TAG=my_docker_user/image_tag:v0.1.1 docker-build
docker-build:
	docker build . -t $(IMAGE_TAG)

docker-run: docker-build
	docker run --rm --name $(NAME) -p $(HTTP_PORT):$(HTTP_PORT) -p $(GRPC_PORT):$(GRPC_PORT) \
		$(IMAGE_TAG) -http.addr :$(HTTP_PORT) -grpc.addr :$(GRPC_PORT)

# Service is generated using Truss.
# See https://github.com/metaverse/truss
truss: service.proto
	truss service.proto --svcout $(PKG)

build: $(GO_FILES)
	GO111MODULE=on go mod tidy
	GO111MODULE=on go build -o $(BIN) cmd/gaproxy-server/main.go

run: build
	$(BIN) -http.addr :$(HTTP_PORT) -grpc.addr :$(GRPC_PORT)

truss-install:
	$(MAKE) -f Makefile.truss

clean:
	rm -rf $(BIN_DIR)

.PHONY: clean docker docker-run run

