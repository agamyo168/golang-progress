.PHONY: build run dev clean test lint

BINARY=arr-dashboard
MAIN=cmd/server/main.go

build:
	go build -o $(BINARY) $(MAIN)

run: build
	./$(BINARY)

dev:
	air

clean:
	rm -f $(BINARY)

test:
	go test ./...

lint:
	go vet ./...
