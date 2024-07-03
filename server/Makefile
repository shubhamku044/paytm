GOCMD=go
GOBUILD=$(GOCMD) build
BUILD_DIR=bin
MAIN_FILE=./cmd/paytm/main.go
BINARY_NAME=paytm

build:
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_FILE) -v

run-build:
	$(BUILD_DIR)/$(BINARY_NAME)

dev:
	air

run:
	go run $(MAIN_FILE)