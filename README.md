# Golang DDD Demo

This repository is a demonstration of Domain-Driven Design (DDD) principles in a Golang application.

## DDD Project Structure

- **cmd/**: Application entry points (main packages)
- **internal/**: Private application and domain code
  - **domain/**: Domain layer (entities, aggregates, value objects, repositories, domain services)
  - **application/**: Application services, DTOs, use cases
  - **infrastructure/**: Implementation of repository interfaces, external services, database, etc.
  - **interfaces/**: HTTP handlers, gRPC, CLI, etc.
- **pkg/**: Public libraries for use by other projects

## Getting Started

1. Clone the repo
2. Run `go mod init <module-name>`
3. Start building your DDD-based Go application

---
Feel free to modify the structure as your project evolves.
