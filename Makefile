BIN_DIR := bin
PROJECT_NAME := book-store-api-server
TAG ?= latest

all: build

$(BIN_DIR):
	mkdir -p $(BIN_DIR)

build: $(BIN_DIR)
	go build -o $(BIN_DIR)/$(PROJECT_NAME)

clean:
	rm -rf $(BIN_DIR)

docker-build: build
	docker build -t $(PROJECT_NAME):$(TAG) .


.PHONY: all build clean docker-build
