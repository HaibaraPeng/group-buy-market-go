# Group Buy Market (Go)

A distributed group buying marketplace built with Go, following Domain-Driven Design (DDD) principles and Clean Architecture.

## Project Structure

This project follows a DDD-based layered architecture:

```
├── api/                   # API definition files (protobuf, swagger, etc.)
├── cmd/                   # Application entry points
│   └── groupbuy/          # Main application
│       └── main.go        # Entry point
│       └── main_kratos.go # Kratos framework entry point
├── configs/               # Configuration files
├── docs/                  # Documentation
├── internal/              # Private application code
│   ├── application/       # Application services
│   ├── domain/            # Domain layer (entities, value objects, aggregates, domain services)
│   ├── infrastructure/    # Infrastructure layer (database, external services)
│   └── interfaces/        # Interfaces layer (HTTP, gRPC handlers)
│       ├── http/          # HTTP handlers and server
│       └── grpc/          # gRPC handlers and server
└── pkg/                   # Shared/cross-project code (if any)
```

## Layers Overview

### 1. Domain Layer (`internal/domain`)
Contains enterprise business logic and domain entities:
- Entities: Product, User, Order
- Value Objects
- Aggregates
- Domain Services
- Repository Interfaces
- Strategy Pattern Nodes (RootNode, SwitchNode, MarketNode, EndNode, ErrorNode)

### 2. Application Layer (`internal/application`)
Contains application-specific business logic:
- Use cases
- DTOs (Data Transfer Objects)
- Orchestration of domain objects

### 3. Infrastructure Layer (`internal/infrastructure`)
Contains technology-specific implementations:
- Database implementations
- External service clients
- Message queues
- File systems

### 4. Interfaces Layer (`internal/interfaces`)
Contains adapters that convert external requests into internal formats:
- HTTP handlers
- gRPC services
- CLI commands

## Getting Started

### Prerequisites
- Go 1.16 or higher

### Installation
```bash
go mod tidy
```

### Running the Application
```bash
# Run with standard HTTP server
go run cmd/groupbuy/main.go

# Run with Kratos framework
go run cmd/groupbuy/main_kratos.go
```

### Building the Application
```bash
# Standard build
go build -o bin/groupbuy cmd/groupbuy/main.go

# Kratos build
go build -o bin/groupbuy_kratos cmd/groupbuy/main_kratos.go
```

## Dependency Injection
This project uses [Google Wire](https://github.com/google/wire) for dependency injection.

To regenerate Wire injection code:
```bash
wire cmd/groupbuy/wire.go
```

## Configuration
Configuration files are located in the `configs/` directory.

## Testing
```bash
go test ./...
```