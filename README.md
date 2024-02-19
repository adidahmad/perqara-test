[Work In Progress]
# CRUD API

## Table of Contents
- [Table of Contents](#table-of-contents)
- [Introduction](#introduction)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
  - [Installation](#installation)
  - [Configuration](#configuration)
- [Project Structure](#project-structure)
- [Project Command](#project-command)
- [Package Used](#package-used)
- [Postman Collection](#postman-collection)

## Introduction

This Project Implement CRUD Users API

Implement hexagonal architecture (port & adapter). This architectural paradigm allows for the establishment of clear boundaries and a solid separation of concerns within our application. Ensure separation between the core business logic and the supporting peripheral systems. This results in clean boundaries and ensures that changes in peripheral systems do not impact the core business logic. This separation is crucial as it enables easy management, maintenance and testing of the application. The isolated business logic simplifies the testing process as it is not directly tied to external systems. here is the short explanation of key structures:

- **domain**:
  The domain directory holds the core business logic related to notes. Here, you'd find the models, entities and the rules specific to handling notes. This is the heart of the application, where the fundamental operations and validations regarding notes are defined.
- **port**:
  The port directory within the business/notes section includes interfaces or abstract definitions that define how the core business logic interacts with the external world. It encapsulates the contracts or interfaces used by the adapters and defines what operations the adapters should implement to communicate with the core.
- **service**:
  This folder contains additional business logic or services that support the main note-related functionalities. These services might include auxiliary operations or complex algorithms related to notes that aren't part of the core domain logic but are necessary for the application.
- **repository**:
  The repository folder typically holds the concrete implementations of interfaces defined in the port folder. It manages the data persistence and retrieval, interacting with the database or any external data storage.

## Prerequisites

- golang version 1.20+
- mysql database

## Getting Started
### Installation

- install golang
- install mysql
- clone this repo
  ```bash
  git clone https://github.com/adidahmad/perqara-test.git
  ```

### Configuration
create configuration file by copying from example. then fill it. 

```bash
cp env/config_example.json env/config.json
```

## Project Structure
```bash
├── README.md
├── app
│   ├── controllers # Contains controllers for various components along with request handling logic.
│   │   └── users
│   │       ├── controller.go
│   │       └── request.go
│   └── routes
│       └── routes.go
├── cmd
│   ├── migration # Contains SQL files for each database migration script to create tables.
│   │   ├── files
│   │   │   ├── 20240129110109_create_users_table.sql
│   │   └── main.go
│   └── server # Holds the main application server configuration.
│       └── main.go
├── config # Manages how to parse configuration files
│   ├── config.go
│   ├── database.go
├── core # The core logic of the application, divided into separate modules.
│   └── users
│       ├── domain # contains domain models/entities representing core business concepts
│       │   └── domain.go
│       ├── entity # contains models/entities representing database communcation concepts
│       │   └── entity.go
│       ├── port # It contains the interfaces definition used to communicate between layers
│       │   ├── mocks # 
│       │   │   └── port.go
│       │   └── port.go
│       ├── repository
│       │   └── repository.go
│       └── service # entry points to the core and each one of them implements the corresponding port 
│           ├── service.go
│           └── service_test.go # unit test for user service
├── env #  Contains environment-related configurations in JSON format.
│   └── config_example.json
├── go.mod
├── go.sum
├── main.go
├── modules # just like repostitory modules 
│   ├── databases
│   │   ├── gorm
│   │   │   └── gorm.go
│   │   └── mysql
│   │       └── mysql.go
└── utils
    └── validator
        └── validator.go
```

## Project Command
- run project
  ```bash
  go run main.go
  ```
- migrate up
  ```bash
  go run cmd/migration/main.go up
  ```

## Package Used
- **Echo Framework:**
   - Echo is a fast and minimalist web framework for Go. It is used to build web applications and RESTful APIs.
   - Echo provides features like routing, middleware support, and context management, making it easy to develop robust and scalable web applications.
- **Viper:**
   - Viper is a configuration management library for Go. It is used to read and parse configuration files, environment variables, and other configuration sources.
   - Viper simplifies the process of managing application configurations, supporting multiple configuration formats and providing a convenient way to access configuration values.
- **Goose:**
   - Goose is a database migration tool for Go applications. It helps manage database schema changes over time by providing a way to version control and apply migrations.
   - With Goose, developers can maintain the database schema in a structured and versioned manner, making it easy to update databases across different environments.
- **GORM:**
   - GORM is an Object Relational Mapping (ORM) library for Go. It simplifies database interactions by providing a high-level, expressive API for CRUD operations and queries.
   - GORM supports multiple database backends, offers features like associations and hooks, and is commonly used to interact with relational databases in a Go application.
- **Gomock:**
   - Gomock is a mocking framework for Go. It is used in unit testing to create mock implementations of interfaces or functions, allowing developers to isolate and test specific components of their code.
   - Gomock helps in writing effective unit tests by replacing external dependencies with controlled mock objects, ensuring tests focus on the functionality being tested.

## Postman Collection
```bash
- 
```

