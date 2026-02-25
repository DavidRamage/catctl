BINARY_NAME=catctl
build:
	go build -o $(BINARY_NAME) main.go

configfiles:
	mkdir -p ~/.catctl
	cp configs/config.yaml ~/.catctl/config.yaml
	cp configs/radios.yaml ~/.catctl/radios.yaml
clean:
	go clean	
	rm -f $(BINARY_NAME)
