# Group Buy Market (Go)

A distributed group buying marketplace built with Go, following Domain-Driven Design (DDD) principles and Clean Architecture.

## Project Structure

This project follows a DDD-based layered architecture:

```
├── api/                   # API definition files (protobuf, swagger, etc.)
├── cmd/                   # Application entry points
│   └── groupbuy/          # Main application
│       ├── main.go        # Entry point
│       ├── wire.go        # Wire dependency injection definitions
│       └── wire_gen.go    # Auto-generated Wire code
├── common/                # Common utilities and design patterns
│   └── design/tree/       # Strategy tree pattern implementation
├── configs/               # Configuration files
├── docs/                  # Documentation
├── internal/              # Private application code
│   ├── conf/              # Configuration protobuf definitions
│   ├── domain/            # Domain layer (entities, value objects, aggregates, domain services)
│   │   ├── activity/      # Activity domain
│   │   │   ├── model/     # Domain models and value objects
│   │   │   └── service/   # Domain services
│   │   │       ├── discount/  # Discount calculation services
│   │   │       └── trial/     # Trial calculation services
│   │   │           ├── core/      # Core trial components
│   │   │           ├── factory/   # Strategy factory
│   │   │           ├── node/      # Strategy tree nodes
│   │   │           └── thread/    # Thread tasks
│   │   └── tag/           # Tag domain
│   │       ├── model/     # Tag models
│   │       └── service/   # Tag services
│   ├── infrastructure/    # Infrastructure layer (database, external services)
│   │   ├── adapter/       # Adapters for external services
│   │   │   └── repository/    # Repository implementations
│   │   ├── cache/         # Cache implementations
│   │   ├── dao/           # Data Access Objects
│   │   └── po/            # Persistent Objects (data models)
│   └── server/            # Server initialization
└── third_party/           # Third-party proto files
```

## Layers Overview

### 1. Domain Layer (`internal/domain`)
Contains enterprise business logic and domain entities:
- Entities: MarketProductEntity, TrialBalanceEntity
- Value Objects: GroupBuyActivityDiscountVO, SkuVO, SCSkuActivityVO
- Aggregates
- Domain Services: Discount calculation services, Trial calculation services
- Repository Interfaces
- Strategy Pattern Nodes (RootNode, SwitchNode, TagNode, MarketNode, EndNode, ErrorNode)

### 2. Application Layer (`internal/application`)
Contains application-specific business logic:
- Use cases
- DTOs (Data Transfer Objects)
- Orchestration of domain objects

### 3. Infrastructure Layer (`internal/infrastructure`)
Contains technology-specific implementations:
- Database implementations (DAOs)
- External service clients
- Repository implementations
- Cache implementations (Redis)

### 4. Interfaces Layer (`internal/interfaces`)
Contains adapters that convert external requests into internal formats:
- HTTP handlers
- gRPC services

## Key Components

### Strategy Tree Pattern
The project implements a strategy tree pattern for handling group buying calculations:
- **RootNode**: Entry point of the strategy tree
- **SwitchNode**: Determines if marketing activities are enabled
- **TagNode**: Checks if user belongs to target audience
- **MarketNode**: Calculates marketing discounts
- **EndNode**: Finalizes the calculation and returns results
- **ErrorNode**: Handles error cases

### Discount Calculation Services
Various discount calculation algorithms:
- **ZJCalculateService**: Direct reduction calculation
- **ZKCalculateService**: Discount calculation
- **MJCalculateService**: Full reduction calculation
- **NCalculateService**: N yuan purchase calculation

## Getting Started

### Prerequisites
- Go 1.16 or higher
- Redis (for caching and bitmap operations)
- MySQL (for data persistence)

### Installation
```bash
go mod tidy
```

### Running the Application
```bash
# Run with standard HTTP server
go run cmd/groupbuy/main.go
```

### Building the Application
```bash
# Standard build
go build -o bin/groupbuy cmd/groupbuy/main.go
```

## Dependency Injection
This project uses [Google Wire](https://github.com/google/wire) for dependency injection.

To regenerate Wire injection code:
```bash
# Install wire if not already installed
go install github.com/google/wire/cmd/wire@latest

# Generate dependency injection code
cd cmd/groupbuy
go generate
```

## Configuration
Configuration files are located in the `configs/` directory. The project uses protobuf for configuration definitions.

## Testing
```bash
go test ./...
```