# Go parameters
MAIN_PATH=main.go
BINARY_NAME=$(BINARY_PATH)/tiwttor
BINARY_PATH=bin

run:
	go build -o $(BINARY_NAME) -race $(MAIN_PATH)
	./$(BINARY_NAME)

test:
	go test -race -v -timeout=10s ./...

clean:
	go clean $(MAIN_PATH)
	rm -f $(BINARY_PATH)/*
