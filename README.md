# StudentApi

A simple REST API for managing student records, built with Go and SQLite.

## Overview

StudentApi is a lightweight RESTful API that allows you to perform CRUD (Create, Read, Update, Delete) operations on student data. It uses SQLite as the database backend and is designed with a clean architecture using Go's standard library and minimal external dependencies.

## Features

- **Create Student**: Add new student records
- **Read Students**: Retrieve individual students by ID or get a list of all students
- **Update Student**: Modify existing student information
- **Delete Student**: Remove student records
- **Input Validation**: Validates required fields (name, email, age)
- **Graceful Shutdown**: Handles server shutdown gracefully
- **Structured Logging**: Uses slog for logging

## Project Structure

```
.
├── cmd/
│   └── studentApi/
│       └── main.go              # Application entry point
├── config/
│   └── local.yaml               # Configuration file
├── internal/
│   ├── config/
│   │   └── config.go            # Configuration loading
│   ├── http/
│   │   └── handlers/
│   │       └── student/
│   │           └── student.go   # HTTP handlers for student operations
│   ├── storage/
│   │   ├── storage.go           # Storage interface
│   │   └── sqlite/
│   │       └── sqlite.go        # SQLite implementation
│   ├── types/
│   │   └── types.go             # Data types
│   └── utils/
│       └── response/
│           └── response.go      # Response utilities
├── storage/
│   └── storage.db               # SQLite database file (auto-created)
├── go.mod                       # Go module file
├── go.sum                       # Go dependencies
└── README.md                    # This file
```

## Prerequisites

- Go 1.25.6 or later
- SQLite (automatically handled by the application)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/iShinzoo/studentApi.git
   cd studentApi
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. (Optional) Build the application:
   ```bash
   go build -o bin/studentApi cmd/studentApi/main.go
   ```

## Configuration

The application uses a YAML configuration file. A sample configuration is provided in `config/local.yaml`:

```yaml
env: "dev"
storage_path: "storage/storage.db"
http_server:
  address: "localhost:8082"
```

You can create different configuration files for different environments (e.g., `config/prod.yaml`).

## Usage

Run the application with the configuration file:

```bash
go run cmd/studentApi/main.go -config config/local.yaml
```

Or if you built the binary:

```bash
./bin/studentApi -config config/local.yaml
```

The server will start on the address specified in the configuration (default: `localhost:8082`).

## API Endpoints

### Create a Student
- **POST** `/api/students`
- **Body**: JSON
  ```json
  {
    "name": "John Doe",
    "email": "john.doe@example.com",
    "age": 20
  }
  ```
- **Response**: `{"id": 1}`

### Get a Student by ID
- **GET** `/api/students/{id}`
- **Response**: Student object or error

### Get All Students
- **GET** `/api/students`
- **Response**: Array of student objects

### Update a Student
- **PUT** `/api/students/{id}`
- **Body**: JSON (same as create)
- **Response**: `{"message": "student updated successfully"}`

### Delete a Student
- **DELETE** `/api/students/{id}`
- **Response**: `{"message": "student deleted successfully"}`

## Data Model

```go
type Student struct {
    Id    int64  `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
    Age   int    `json:"age"`
}
```

All fields except `id` are required for create/update operations.

## Dependencies

- `github.com/go-playground/validator/v10`: Input validation
- `github.com/ilyakaznacheev/cleanenv`: Configuration loading
- `modernc.org/sqlite`: SQLite driver
- Standard library packages for HTTP, JSON, logging, etc.

## Development

The project follows Go best practices with:
- Clean architecture separation
- Interface-based design for storage layer
- Structured error handling
- Input validation
- Graceful server shutdown

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

