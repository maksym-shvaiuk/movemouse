.PHONY: build-linux build-windows build-macos-intel build-macos-silicon build-all clean

BINARY_NAME := movemouse
BIN_DIR := bin

build:
	go build -o $(BIN_DIR)/$(BINARY_NAME) .

# Clean build artifacts
clean:
	rm -rf $(BIN_DIR)/*
