# Variables
BINARY=microservice
PROTO_DIR=proto
PROTO_FILES=$(PROTO_DIR)/*.proto
DOCKER_IMAGE=microservice:latest

# Default target
all: build

# Build Go binary
build:
	go build -o $(BINARY) ./cmd/server

# Run locally (without Docker)
run: build
	./$(BINARY)

# Generate gRPC code from proto
proto:
	protoc --go_out=. --go-grpc_out=. $(PROTO_FILES)

# Run tests (unit)
test:
	go test ./... -v

# Run integration tests (Postgres required)
integration-test:
	go test ./... -v -tags=integration

# Docker build
docker-build:
	docker-compose build

# Docker up
docker-up:
	docker-compose up

# Docker run
docker-run:
	docker-compose build
	docker-compose up

# Docker down
docker-down:
	docker-compose down

# Clean build artifacts
clean:
	rm -f $(BINARY)
