# Pustaka API

REST API untuk manajemen buku menggunakan Go, Gin Framework, dan PostgreSQL.

## Struktur Project

```
pustaka-api/
├── api/                      # Handler untuk REST API endpoints
│   └── book_handler.go       # Book CRUD handlers
├── config/                   # Konfigurasi aplikasi
│   └── database.go           # Setup database connection
├── exception/                # Custom error handling
│   └── error_handler.go      # Error handler middleware
├── models/                   # Data models/entities
│   └── book.go               # Book entity
├── repository/               # Data access layer
│   └── book_repository.go    # Book repository interface & implementation
├── service/                  # Business logic layer
│   └── book_service.go       # Book service interface & implementation
├── utils/                    # Utility functions dan shared types
│   └── response.go           # Response structures
├── validator/                # Request validation structures
│   └── book_validator.go     # Book request validators
├── main.go                   # Entry point aplikasi
└── go.mod                    # Go module dependencies
```

## Penjelasan Struktur Folder

### 1. **api/**
Berisi HTTP handlers yang menangani request dan response. Layer ini bertanggung jawab untuk:
- Parsing request dari client
- Validasi input
- Memanggil service layer
- Mengembalikan response ke client

### 2. **config/**
Berisi konfigurasi aplikasi seperti:
- Database connection setup
- Environment variables
- Application settings

### 3. **exception/**
Berisi custom error handling:
- Custom error types
- Error handler middleware
- Predefined errors

### 4. **models/**
Berisi entity/model yang merepresentasikan struktur data di database.

### 5. **repository/**
Data access layer yang berinteraksi langsung dengan database:
- Interface untuk abstraksi
- Implementation untuk CRUD operations
- Query ke database menggunakan GORM

### 6. **service/**
Business logic layer yang berisi logika bisnis aplikasi:
- Memanggil repository untuk data access
- Memproses data sesuai business rules
- Mengembalikan result ke handler

### 7. **utils/**
Berisi utility functions dan shared types:
- Response structures
- Helper functions
- Common types

### 8. **validator/**
Berisi request validation structures dengan tag binding untuk Gin.

## API Endpoints

### Books
- `POST /v1/book` - Create a new book
- `GET /v1/book` - Get all books
- `GET /v1/book/:bookId` - Get book by ID
- `PUT /v1/book/:bookId` - Update book by ID
- `DELETE /v1/book/:bookId` - Delete book by ID

## Technology Stack

- **Framework**: Gin Web Framework
- **ORM**: GORM
- **Database**: PostgreSQL
- **Validation**: go-playground/validator

## Running the Application

```bash
go run main.go
```

Server akan berjalan di `http://localhost:3000`

## Database Configuration

Update connection string di `config/database.go`:

```go
dsn := "host=localhost user=youruser password=yourpassword dbname=pustaka_api port=5432 sslmode=disable TimeZone=Asia/Jakarta"
```

## Architecture Pattern

Project ini menggunakan **Clean Architecture** dengan separation of concerns:

1. **Presentation Layer** (api/) - Handles HTTP requests
2. **Business Logic Layer** (service/) - Contains business rules
3. **Data Access Layer** (repository/) - Database operations
4. **Domain Layer** (models/) - Core entities

Keuntungan pattern ini:
- **Testability**: Setiap layer dapat di-test secara independen
- **Maintainability**: Perubahan di satu layer tidak affect layer lain
- **Scalability**: Mudah untuk menambah fitur baru
- **Flexibility**: Mudah untuk mengganti implementasi (misal ganti database)
