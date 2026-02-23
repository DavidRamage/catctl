BINARY_NAME=catctl
build:
	go build -o $(BINARY_NAME) main.go

configfiles:
	mkdir -p ~/.catctl
	cp configs/commands.yaml ~/.catctl/commands.yaml
	cp configs/serial.yaml ~/.catctl/serial.yaml
clean:
	go clean	
	rm -f $(BINARY_NAME)
