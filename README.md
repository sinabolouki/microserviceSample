# 🧪 Microservice Architecture with GraphQL Gateway (Go + gRPC + PostgreSQL)

## 🎯 Goal

This project demonstrates a microservices architecture using **Go**, **gRPC**, **PostgreSQL**, and a unified **GraphQL Gateway**. It allows:

- **User Management**: create and retrieve users, fetch their orders
- **Catalogue Management**: create and retrieve catalogue items, fetch orders containing an item
- **Order Management**: create orders with positions referencing catalogue items

---

## 🧱 Services Overview

| Service           | Port  | Description                                 |
| ----------------- | ----- | ------------------------------------------- |
| User Service      | 50051 | Manages user data                           |
| Catalogue Service | 50052 | Manages catalogue item data                 |
| Order Service     | 50053 | Manages orders and validates user/items     |
| GraphQL Gateway   | 8080  | Exposes unified GraphQL API to the frontend |

---

## ⚙️ Setup Instructions

### 1. ✅ Prerequisites

Ensure you have the following installed:

- Go 1.22+
- PostgreSQL
- `protoc` (Protocol Buffers compiler)
- `gqlgen` (`go install github.com/99designs/gqlgen@latest`)

---

### 2. 🛠 Database Setup

Create the necessary PostgreSQL databases:

```bash
createdb microservice_example
```

Run schema files from each service:

```bash
psql microservice_example < user-service/schema/init.sql
psql microservice_example < catalogue-service/schema/init.sql
psql microservice_example < order-service/schema/init.sql
```

---

### 3. 🚀 Running the Services

Run each service in its own terminal:

#### User Service

```bash
cd user-service
go run main.go service.go
```

#### Catalogue Service

```bash
cd catalogue-service
go run main.go service.go
```

#### Order Service

```bash
cd order-service
go run main.go service.go
```

#### GraphQL Gateway

```bash
cd gateway
go run server.go
```

---

### 🔐 Environment Variables (Optional)

Each service supports an environment variable `POSTGRES_CONN` for DB connection override:

```bash
export POSTGRES_CONN=postgres://postgres:password@localhost:5432/userdb?sslmode=disable
```

Defaults are used if not provided.

---

## 🔍 Testing

Visit: [http://localhost:8080/](http://localhost:8080/) to open GraphQL Playground.

You can:

- Create users and catalogue items
- Create orders referencing users and items
- Fetch users with nested order history
- Fetch items with related orders

See the provided `graphql_queries.graphql` file for full examples.

---

## 🧠 Technologies Used

- Go 1.22
- gRPC
- PostgreSQL
- gqlgen (GraphQL)
- Protobuf
- UUIDs for IDs

---

## ✅ Features

- Modular microservices with separate databases
- Unified API via GraphQL
- gRPC communication with service validation
- PostgreSQL persistence per service

---

## 🧪 Testing Workflow

1. Create a user
2. Create multiple catalogue items
3. Create an order referencing user and item IDs
4. Query user and item data with nested orders and positions