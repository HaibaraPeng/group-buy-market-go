# Group Buy Market (Go)

A distributed group buying marketplace built with Go, following Domain-Driven Design (DDD) principles and Clean Architecture.

## Project Structure

This project follows a DDD-based layered architecture with additional design patterns like Strategy Tree and Responsibility Chain Pattern for flexible business logic processing:

```
├── api/                   # API definition files (protobuf)
│   ├── v1/                # API v1 definitions
│   │   ├── activity.proto # Activity service definitions
│   │   ├── dcc.proto      # Dynamic configuration service definitions
│   │   ├── index.proto    # Index service definitions
│   │   ├── tag.proto      # Tag service definitions
│   │   └── trade.proto    # Trade service definitions
├── cmd/                   # Application entry points
│   └── groupbuy/          # Main application
│       ├── main.go        # Entry point
│       ├── wire.go        # Wire dependency injection definitions
│       └── wire_gen.go    # Auto-generated Wire code
├── configs/               # Configuration files
│   └── config.yaml        # Main configuration file
├── internal/              # Private application code
│   ├── common/            # Common utilities and constants
│   │   ├── consts/        # Constants and error codes
│   │   ├── design/        # Design pattern implementations
│   │   │   ├── link/      # Responsibility chain pattern implementations
│   │   │   │   ├── model1/ # First responsibility chain model
│   │   │   │   └── model2/ # Second responsibility chain model
│   │   │   └── tree/      # Strategy tree pattern implementation
│   │   ├── exception/     # Exception handling utilities
│   │   ├── utils/         # Utility functions
│   │   │   ├── date_util.go    # Date utility functions
│   │   │   └── user_util.go    # User utility functions
│   ├── conf/              # Configuration protobuf definitions
│   ├── domain/            # Domain layer (entities, value objects, aggregates, domain services)
│   │   ├── activity/      # Activity domain
│   │   │   ├── biz/       # Business logic implementations
│   │   │   │   ├── discount/  # Discount calculation services
│   │   │   │   └── trial/     # Trial calculation services
│   │   │   │       ├── core/      # Core trial components
│   │   │   │       ├── factory/   # Strategy factory
│   │   │   │       ├── node/      # Strategy tree nodes (RootNode, SwitchNode, TagNode, MarketNode, EndNode, ErrorNode)
│   │   │   │       └── thread/    # Thread tasks
│   │   │   └── model/     # Domain models and value objects
│   │   ├── tag/           # Tag domain
│   │   │   └── model/     # Tag models
│   │   └── trade/         # Trade domain
│   │       ├── biz/       # Trade business logic
│   │       │   ├── lock/         # Trade lock business logic
│   │       │   │   └── filter/   # Trade lock filters (ActivityUsabilityRuleFilter, UserTakeLimitRuleFilter, TradeLockRuleFilterFactory)
│   │       │   └── settlement/   # Trade settlement business logic
│   │       │       └── filter/   # Trade settlement filters (SCRuleFilter, SettableRuleFilter, OutTradeNoRuleFilter, EndRuleFilter, TradeSettlementRuleFilterFactory)
│   │       ├── model/     # Trade models
│   │       └── trade.go   # Trade interface definition
│   ├── infrastructure/    # Infrastructure layer (database, external services)
│   │   ├── adapter/       # Adapters for external services
│   │   │   ├── port/      # Port interfaces
│   │   │   └── repository/    # Repository implementations
│   │   ├── dao/           # Data Access Objects
│   │   ├── data/          # Data layer with database and redis clients
│   │   ├── dcc/           # Dynamic configuration client
│   │   ├── event/         # Event publishing and listening implementations
│   │   │   ├── listener/  # Event listeners
│   │   │   └── publish/   # Event publishers
│   │   ├── gateway/       # Gateway implementations
│   │   ├── job/           # Scheduled job implementations
│   │   ├── po/            # Persistent Objects (data models)
│   │   └── provider.go    # DCC provider implementation
│   ├── server/            # Server initialization
│   └── service/           # Service layer (application services)
│       ├── activity_service.go # Activity application service
│       ├── dcc_service.go      # Dynamic configuration application service
│       ├── index_service.go    # Index application service
│       ├── tag_service.go      # Tag application service
│       └── trade_service.go    # Trade application service
├── third_party/           # Third-party proto files
│   └── google/api/        # Google API annotations and HTTP rules
├── README.md              # This file
└── openapi.yaml           # OpenAPI specification generated from protobuf
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

### 2. Application Layer (`internal/service`)
Contains application-specific business logic:
- Use cases
- Service implementations:
  - ActivityService (marketing trial service)
  - IndexService (group buying market configuration queries)
  - TagService (tag batch job service)
  - DccService (dynamic configuration service)
  - TradeService (trade order service)

### 3. Infrastructure Layer (`internal/infrastructure`)
Contains technology-specific implementations:
- Database implementations (DAOs)
- External service clients
- Repository implementations
- Cache implementations (Redis)
- Data layer with database and Redis clients (Data struct)
- Event publishing and listening implementations
- Job scheduling implementations
- Dynamic Configuration Center (DCC)

### 4. Interfaces Layer (`internal/server`)
Contains adapters that convert external requests into internal formats:
- HTTP server initialization
- Service registration

## Key Components

### Strategy Tree Pattern
The project implements a strategy tree pattern for handling group buying calculations:
- **RootNode**: Entry point of the strategy tree
- **SwitchNode**: Determines if marketing activities are enabled (handles downgrade switches and cut range checks)
- **TagNode**: Checks if user belongs to target audience
- **MarketNode**: Calculates marketing discounts using various discount calculation services
- **EndNode**: Finalizes the calculation and returns results
- **ErrorNode**: Handles error cases

### Discount Calculation Services
Various discount calculation algorithms:
- **ZJCalculateService**: Direct reduction calculation
- **ZKCalculateService**: Discount calculation
- **MJCalculateService**: Full reduction calculation
- **NCalculateService**: N yuan purchase calculation

### Responsibility Chain Pattern Filters
The project implements responsibility chain pattern filters for trade rule validation in the trade settlement domain:
- **SCRuleFilter**: Validates channel blacklist using DCC configuration
- **SettableRuleFilter**: Validates transaction time against group buy valid time range
- **OutTradeNoRuleFilter**: Validates external trade number existence
- **EndRuleFilter**: Finalizes the settlement process and returns results
- **TradeSettlementRuleFilterFactory**: Assembles the filter chain for trade settlement validation

Additional responsibility chain filters for trade lock:
- **ActivityUsabilityRuleFilter**: Validates if the activity is usable
- **UserTakeLimitRuleFilter**: Validates if user has reached participation limit
- **TradeLockRuleFilterFactory**: Assembles the filter chain for trade lock validation

### Dynamic Configuration Center (DCC)
The project includes a Dynamic Configuration Center for runtime configuration management:
- **DCC Service**: Provides configuration retrieval and update capabilities
- **Channel Blacklist Configuration**: Allows dynamic management of blocked source-channel combinations
- **Runtime Configuration**: Enables configuration changes without application restart
- **Downgrade Switch**: Provides capability to disable features during high load
- **User Cut Range**: Supports percentage-based feature rollout

### Event Publishing System
The project includes an event publishing system for message-based communication:
- **RabbitMQ Event Publisher**: Publishes events to RabbitMQ queues
- **Topic-based messaging**: Uses topic exchanges for flexible routing
- **Integration testing**: Includes integration tests for real RabbitMQ connections

### Job Scheduling System
The project includes a scheduled job system for background tasks:
- **Group Buy Notification Job**: Handles group buy completion notifications
- **Cron-based scheduling**: Uses cron expressions for flexible scheduling
- **Context-aware execution**: Maintains context across job executions

### Error Code Management
Comprehensive error code system with business-specific error codes for different domains:
- **System-level errors**: SUCCESS, UN_ERROR, ILLEGAL_PARAMETER, INDEX_EXCEPTION, UPDATE_ZERO
- **Business-level errors**: E0001-E0007 for general business operations
- **Activity-related errors**: E0101-E0106 for activity and settlement operations
- **Consistent error handling**: Standardized error codes across Java and Go implementations

### Date Utilities
The project includes date utility functions for common time operations:
- **DifferenceDateTime2Str**: Calculates the difference between two time values and formats as HH:MM:SS

### Data Persistence
The project uses modern data access patterns with GORM for database operations and Redis for caching:
- **MySQL with GORM**: Provides ORM capabilities with database abstraction
- **Redis for caching**: Implements caching and bitmap operations for user targeting
- **DAO Layer**: Data Access Objects abstract database operations
- **Repository Pattern**: Provides domain-specific data access interfaces
- **Persistent Objects (PO)**: Data models that map to database tables
- **RabbitMQ Integration**: Message queue support for asynchronous processing

## Getting Started

### Prerequisites
- Go 1.16 or higher (recommended: Go 1.23.12)
- Redis (for caching and bitmap operations)
- MySQL (for data persistence)
- RabbitMQ (for event publishing)

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

## Configuration
Configuration files are located in the `configs/` directory. The project uses protobuf for configuration definitions and YAML for runtime configurations. The DCC (Dynamic Configuration Center) allows runtime configuration changes without restarts.

The main configuration file (`configs/config.yaml`) includes:
- Server settings (HTTP and gRPC addresses and timeouts)
- Database configuration (MySQL connection details)
- Redis configuration (address, timeouts)
- RabbitMQ configuration (addresses, credentials, queue settings)

## Dependency Injection
This project uses [Google Wire](https://github.com/google/wire) for dependency injection.

To regenerate Wire injection code:
```bash
# Install wire if not already installed
go install github.com/google/wire/cmd/wire@latest

# Generate dependency injection code
cd cmd/groupbuy
wire
```

## Testing
```bash
go test ./...
```

The project includes both unit tests and integration tests (marked with `//go:build integration` tags).

## API Documentation
API endpoints are defined in the `api/` directory using Protocol Buffers. The service definitions include HTTP bindings for RESTful access to gRPC services. The main services are:
- IndexService: For group buying market configuration queries
- TagService: For user targeting and segmentation
- TradeService: For order processing and management
- DccService: For dynamic configuration management

An OpenAPI specification is automatically generated and available in the `openapi.yaml` file for REST API documentation.

## Architecture Patterns
The application implements several architectural patterns:
- **Domain-Driven Design (DDD)**: Clear separation of domain logic from infrastructure concerns
- **Strategy Pattern**: For handling different discount calculation methods
- **Chain of Responsibility**: For processing trade rules and validations
- **Event-Driven Architecture**: Using RabbitMQ for asynchronous communication
- **Repository Pattern**: Abstracting data access logic
- **Dependency Injection**: Using Google Wire for managing dependencies