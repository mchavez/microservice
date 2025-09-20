# User Microservice (Go + Gin + gRPC + PostgreSQL)
This project is a clean-architecture microservice built with Golang featuring:

- Gin for REST API
- gRPC for RPC communication
- PostgreSQL as database (via Docker)
- Clean Architecture (Entities → UseCases → Repositories → Delivery)
- Swagger/OpenAPI documentation
- Docker & Docker Compose for containerization
- Unit & Integration Tests
- Makefile for automation

📌 Tech Stack
```bash
Go
Gin
gRPC
PostgreSQL
Swagger
Docker
```

---
## 📂 Project Structure
```bash
├── cmd/
│ └── server/ # Application entry point
├── internal/
│ └── user/
│ ├── delivery/ # REST (Gin) + gRPC delivery
│ │ └── http/
│ │ └── handler.go
│ ├── entity/ # Core business models
│ ├── repository/ # Repository interfaces + implementations
│ └── usecase/ # Business logic
├── proto/ # gRPC proto files
├── docs/ # Swagger docs (auto-generated)
├── Dockerfile
├── docker-compose.yml
├── Makefile
└── README.md
```

---

## Getting Started

### 1. Clone Repository
```bash
git clone https://github.com/mchavez/microservice.git
cd microservice
```

### 2. Build image
```bash
make docker-build
```

### 3. Run containers
```bash
make docker-up
```

### 4. Stop containers
```bash
make docker-down
```

### 5. Run Locally
```bash
make run
```

### 6. Run with Postgres via docker-compose:
This will run Postgres, migrations, then your service. REST at http://localhost:8080.
```bash
docker-compose up --build
```

REST API will be available at:
http://localhost:8080

Swagger UI at:
http://localhost:8080/swagger/index.html

📡 REST API (Gin)
Create User
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Miguel"}'
```

Get Users
```bash
curl http://localhost:8080/users
```

Get User by id
```bash
curl http://localhost:8080/users/1
```

Get User by name
```bash
curl http://localhost:8080/users/search/Miguel
```

gRPC API proto/user.proto
```bash
service UserService {
  rpc GetUsers (ListUsersRequest) returns (ListUsersResponse);
  rpc AddUser (User) returns (User);
  rpc GetUserByID (GetUserByIDRequest) returns (GetUserByIDResponse); // NEW
  rpc GetUsersByName (GetUsersByNameRequest) returns (GetUsersByNameResponse); // NEW
}
```

Use any gRPC client (e.g., evans, grpcurl) to test:
```bash
grpcurl -plaintext localhost:50051 user.UserService/GetUsers

grpcurl -plaintext -d '{"name":"Miguel"}' localhost:50051 user.UserService/AddUser

grpcurl -plaintext -d '{"id":1}' localhost:50051 user.UserService/GetUserByID

grpcurl -plaintext -d '{"name":"Miguel"}' localhost:50051 user.UserService/GetUsersByName
```
    
Installing protoc, protoc-gen-go
```bash
    brew install protobuf
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    export PATH="$PATH:$(go env GOPATH)/bin"
    protoc --go_out=. --go-grpc_out=. proto/user.proto
```

Verifyimg protoc-gen-go
```bash
    source ~/.zshrc
    which protoc-gen-go
    protoc-gen-go --version
```

Installing swagger-go
```bash
go install github.com/swaggo/swag/cmd/swag@latest
go get github.com/swaggo/gin-swagger
go get github.com/swaggo/files
-- fix issue -- https://github.com/swaggo/swag/issues/1568 --
go get -u github.com/swaggo/swag
```

Testing
Run all tests:
```bash
make test
```

Run integration tests (requires DB):
```bash
make integration-test
```

⚙️ Makefile Commands
| Command                 | Description                       |
| ----------------------- | --------------------------------- |
| `make build`            | Build Go binary                   |
| `make run`              | Run app locally                   |
| `make proto`            | Generate gRPC code                |
| `make test`             | Run unit tests                    |
| `make integration-test` | Run integration tests (DB needed) |
| `make docker-build`     | Build Docker image                |
| `make docker-up`        | Start containers                  |
| `make docker-run`       | Build & Start containers          |
| `make docker-down`      | Stop containers                   |
| `make clean`            | Remove binary                     |
