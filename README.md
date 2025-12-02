# Group Buy Market (Go)

A distributed group buying marketplace built with Go, following Domain-Driven Design (DDD) principles and Clean Architecture.

## Project Structure

This project follows a DDD-based layered architecture:

```
├── api/                   # API definition files (protobuf, swagger, etc.)
├── cmd/                   # Application entry points
│   └── groupbuy/          # Main application
│       └── main.go        # Entry point
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
- Entities: [Product](file:///C:/Users/Roc/Documents/Code/Me/group-buy-market-go/internal/domain/model.go#L4-L8), [User](file:///C:/Users/Roc/Documents/Code/Me/group-buy-market-go/internal/domain/model.go#L11-L15), [Order](file:///C:/Users/Roc/Documents/Code/Me/group-buy-market-go/internal/domain/model.go#L18-L23)
- Value Objects
- Aggregates
- Domain Services
- Repository Interfaces

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
go run cmd/groupbuy/main.go
```

### Building the Application
```bash
go build -o bin/groupbuy cmd/groupbuy/main.go
```

## Dependency Injection
This project uses [Google Wire](https://github.com/google/wire) for dependency injection.

To regenerate Wire injection code:
```bash
wire cmd/groupbuy/internal/wire.go
```

## Configuration
Configuration files are located in the `configs/` directory.

## Testing
```bash
go test ./...
```