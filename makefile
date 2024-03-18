APP_NAME=deck-api
BUILD_DIR=build
SRC_DIR=.

.PHONY: clean build run
build:	clean
	go build -o $(BUILD_DIR)/$(APP_NAME) $(SRC_DIR)/main.go
serve: build
	./$(BUILD_DIR)/$(APP_NAME) serve
test: 
	go test ./...
coverage: 
	go test -cover -coverprofile=c.out ./... && go tool cover -html="c.out"