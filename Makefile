default: build/gigawallet

.PHONY: clean
clean:
	rm -rf ./build

build/gigawallet: clean
	mkdir -p build/
	go build -o build/gigawallet ./cmd/gigawallet/main.go 
