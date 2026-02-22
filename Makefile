BINARY_NAME=catctl
build:
	go build -o $(BINARY_NAME) main.go
clean:
	go clean	
	rm -f $(BINARY_NAME)
