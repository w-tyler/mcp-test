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

---

# Golang DDD 示例

本仓库用于演示如何在 Golang 项目中实践领域驱动设计（DDD）架构。

## DDD 项目结构

- **cmd/**：应用程序入口（main 包）
- **internal/**：私有应用和领域代码
  - **domain/**：领域层（实体、聚合、值对象、仓储接口、领域服务）
  - **application/**：应用服务、DTO、用例
  - **infrastructure/**：仓储接口实现、外部服务、数据库等
  - **interfaces/**：接口层，如 HTTP 处理器、gRPC、CLI 等
- **pkg/**：可供其他项目使用的公共库

## 快速开始

1. 克隆本仓库
2. 运行 `go mod init <module-name>`
3. 按照 DDD 架构开始构建你的 Go 应用

---
如有需要，可根据项目实际情况调整结构。
