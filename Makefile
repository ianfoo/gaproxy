NAME := gaproxy
USER := ianfoo

BIN_DIR ?= bin
BIN ?= $(BIN_DIR)/$(NAME)
IMAGE_TAG ?= $(USER)/$(NAME)
HTTP_PORT ?= 5050

GO_FILES := $(shell find . -name '*.go')

docker:
	docker build . -t $(IMAGE_TAG)
	if [[ -n "$(EXTRA_TAG)" ]]; then docker tag $(IMAGE_TAG) $(EXTRA_TAG); fi

docker-run: docker
	docker run --rm --name $(NAME) -p $(HTTP_PORT):$(HTTP_PORT) $(IMAGE_TAG)

build: $(GO_FILES)
	GO111MODULE=on go build -o $(BIN) cmd/gaproxy-server/main.go

run: build
	$(BIN)

clean:
	rm -rf $(BIN_DIR)

.PHONY: clean docker docker-run

