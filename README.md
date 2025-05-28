# go-monolith
A sample Go monolith showcasing clean, modular architecture and best practices for scalable backend systems.  
This project demonstrates how to structure a production-ready Go codebase for extensibility, maintainability, and clarity.

## Overview
This repository features two services, **accounts** and **payments**, organized as independent modules within a single Go application.  
The structure and layering are designed to make it easy to extract services into standalone microservices as the codebase evolves.

**Key features:**
- Modular, domain-driven folder structure
- Example SQL migrations for database setup
- Simple Makefile-based workflow
- Ready for containerized/local development

## Project Structure
```
go-monolith/
├── api/
│   └── graphql/            # GraphQL schema definitions
├── cmd/
│   ├── server/             # Entrypoint: dependency injection, router, env setup
│   └── main.go
├── config/                 # App configuration files and helpers
├── db/                     # DB config, connections, migrations, shared access logic
│   └── migrations/         # SQL migrations for accounts and payments
├── pkg/                    # Shared utilities (logger, error handling, etc.)
├── service/
│   ├── accounts/           # Account service module: functions, converters, validators, etc.
│   │   └── data/           # ORM models and account repository (data layer)
│   ├── payments/           # Payments service module: functions, converters, validators, etc.
│   │   └── data/           # ORM models and payment repository (data layer)
├── Makefile
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── go.sum
```

## Getting Started

### Prerequisites
- [Go](https://golang.org/doc/install) 1.23+
- [Docker](https://docs.docker.com/get-docker/) (for local database)
- [golang-migrate](https://github.com/golang-migrate/migrate) (optional, for DB migrations)

### Common Commands
```sh
make run         # Build and run the binary and database
make log         # Watch logs from Docker container
make down        # Tear down the binary and database
make generate    # Generate GraphQL types
make test        # Run unit tests
```

### Extending the Project
This project is structured so that each service (e.g., `accounts`, `payments`) could easily be extracted to its own microservice.

To add new modules, simply mirror the pattern in `service/` with the `data` layer.