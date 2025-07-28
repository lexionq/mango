BINARY_NAME=mango
BUILD_DIR=build
INSTALL_DIR=/usr/bin

.PHONY: all build run clean fmt install

all: build

build:
	go build -o $(BUILD_DIR)/$(BINARY_NAME) main.go

run:
	go run main.go

clean:
	rm -rf $(BUILD_DIR)

fmt:
	go fmt ./...

install: build
	chmod +x $(BUILD_DIR)/$(BINARY_NAME)
	sudo cp $(BUILD_DIR)/$(BINARY_NAME) $(INSTALL_DIR)/$(BINARY_NAME)

