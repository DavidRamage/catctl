BINARY_NAME=catctl
build:
	go build -o $(BINARY_NAME) main.go
configfiles:
	mkdir -p ~/.catctl
	sudo cp configs/commands.yaml ~./catctl/commands.yaml
clean:
	go clean	
	rm -f $(BINARY_NAME)
