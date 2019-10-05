.DEFAULT_GOAL := default

IMAGE := averagemarcus/chaos-gaming:latest

.PHONY: format # Format all go code
format:
	@gofmt -s -w .

.PHONY: build # Build the Go binary
build:
	@go build -o chaos-gaming cmd/chaos-gaming/main.go

.PHONY: run # Run the application
run:
	@go run cmd/chaos-gaming/main.go

.PHONY: docker-build # Build the docker image
docker-build:
	@docker build -t $(IMAGE) .

.PHONY: test # Run all tests
test:
	@go vet && go test

.PHONY: publish # Publish the docker image to the Artifactory registry
publish:
	@docker push $(IMAGE)

.PHONY: help # Show this list of commands
help:
	@grep '^.PHONY: .* #' Makefile | sed 's/\.PHONY: \(.*\) # \(.*\)/\1: \2/' | expand -t20

default: format test build

