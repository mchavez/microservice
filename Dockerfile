# Build stage
FROM golang:1.24 AS builder

# Install protoc and necessary development tools
RUN apt-get update && apt-get install -y \
    protobuf-compiler \
    build-essential \
    git \
    && rm -rf /var/lib/apt/lists/*

# Install Go Protobuf plugins
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Set GOPATH and add it to PATH
ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH

# Sets the working directory inside the container
WORKDIR /app

# Copies the Go module files.
COPY go.mod go.sum ./

# Downloads the application's dependencies.
RUN go mod download

# Copy all files
COPY . .

# Generate proto code
RUN protoc --go_out=. --go-grpc_out=. proto/user.proto

# Build binary
# CGO_ENABLED=0 disables C bindings, and GOOS=linux ensures it's built for Linux.
RUN CGO_ENABLED=0 GOOS=linux go build -o microservice ./cmd/server

# RUN ls -l microservice
# RUN test -f microservice
# Copy all files
# COPY ./microservice /usr/local/bin/

# Stage 2: Create a minimal runtime image
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/microservice .

# Expose ports for REST + gRPC
EXPOSE 8080 50051

CMD ["./server"]

