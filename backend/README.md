# Bookmarks Backend

A Go backend service for managing website bookmarks with automatic metadata scraping.

## Features

- RESTful API for bookmark management
- Automatic metadata extraction (title, description, favicon)
- PostgreSQL database storage
- CORS support for frontend integration
- Graceful shutdown handling

## Prerequisites

- Go 1.21 or higher
- PostgreSQL 12 or higher
- Docker and Docker Compose (for local development)

## Setup

1. Start the PostgreSQL database:
```bash
docker-compose up -d postgres
```

2. Run the database migrations:
```bash
docker exec -i bookmarks_db psql -U postgres -d bookmarks_db < migrations/001_create_bookmarks_table.sql
```

3. Configure environment variables (optional):
```bash
export PORT=8081  # Default: 8081
export DATABASE_URL="postgres://postgres:postgres@localhost:5432/bookmarks_db?sslmode=disable"
```

## Development

1. Run the server:
```bash
go run cmd/server/main.go
```

The server will start on port 8081 (or the configured PORT).

## Project Structure

- `cmd/server`: Main application entry point
- `internal/api`: HTTP handlers and routing
- `internal/models`: Data models
- `internal/scraper`: Webpage metadata scraping
- `internal/storage`: Database operations
- `migrations`: SQL migration files

## API Documentation

The API is documented using OpenAPI 3.0. See `openapi.yaml` for the full specification.

### Endpoints

#### Create Bookmark
```http
POST /api/bookmarks
Content-Type: application/json

{
    "url": "https://example.com"
}
```

#### List Bookmarks
```http
GET /api/bookmarks
```

#### Get Bookmark
```http
GET /api/bookmarks/{id}
```

#### Delete Bookmark
```http
DELETE /api/bookmarks/{id}
```

## Error Handling

The API returns appropriate HTTP status codes:

- 200: Success
- 400: Bad Request (invalid input)
- 404: Not Found
- 500: Internal Server Error

## Security

- Input validation for URLs
- Prepared statements for database queries
- CORS headers for frontend integration
- Request timeouts
- Connection pooling

## Testing

Run the tests:
```bash
go test ./...
```

## Performance

- Connection pooling for database
- Proper indexing on database tables
- Configurable timeouts for HTTP operations
- Graceful shutdown handling