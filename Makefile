PROJECT_NAME=thdl

default: build

help:
	@echo "Build targets:"
	@echo "  all      Run fmt vet build."
	@echo "  build    Build binary."
	@echo "  default  Run build."
	@echo "Quality targets:"
	@echo "  fmt       Format files with go fmt."
	@echo "  vet       Examine go sources with go vet."
	@echo "  errcheck  Examine go sources with errcheck."
	@echo "Test targets:"
	@echo "  test        Run go test."
	@echo "  test-check  Run check command tests."
	@echo "  test-all    Run all tests."
	@echo "Other targets:"
	@echo "  help  Print help message."


# Build targets
all: fmt vet build

build:
	go build -v -o $(PROJECT_NAME) ./cmd/thdl


# Quality targets
fmt:
	go fmt ./...

vet:
	go vet ./...

errcheck:
	errcheck -verbose ./...


# Test targets
test:
	go test ./...

test-check:
	@./scripts/test-check.sh

test-all: test test-check


# Installation targets
install:
	cp $(PROJECT_NAME) /usr/bin

uninstall:
	rm /usr/bin/$(PROJECT_NAME)
