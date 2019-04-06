NAME := gaproxy
USER := ianfoo
OUT ?= bin/$(NAME)
IMAGE_TAG ?= $(USER)/$(NAME)
HTTP_PORT ?= 5050

docker:
	docker build . -t $(IMAGE_TAG)
	if [[ -n "$(EXTRA_TAG)" ]]; then docker tag $(IMAGE_TAG) $(EXTRA_TAG); fi

docker-run:
	docker run --rm --name $(NAME) -p $(HTTP_PORT):$(HTTP_PORT) $(IMAGE_TAG)

build:
	GO111MODULE=on go build -o $(OUT) ./...

.PHONY: docker docker-run

