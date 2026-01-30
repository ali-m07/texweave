.PHONY: build install test tidy clean

BINARY := texweave
MAIN   := ./cmd/texweave

build:
	go build -o $(BINARY) $(MAIN)

install:
	go install $(MAIN)

test:
	go test ./...

tidy:
	go mod tidy

clean:
	rm -f $(BINARY)
