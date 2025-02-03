# Bookmarks Go Backend

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
- Make (optional, for using Makefile commands)

## Setup

1. Create a PostgreSQL database:
```sql
CREATE DATABASE bookmarks;
```

2. Run the database migrations:
```bash
psql -U postgres -d bookmarks -f migrations/001_create_bookmarks_table.sql
```

3. Configure environment variables (optional):
```bash
export PORT=8080  # Default: 8080
export DATABASE_URL="postgres://postgres:postgres@localhost:5432/bookmarks?sslmode=disable"
```

## Building and Running

1. Build the server:
```bash
go build -o bookmarks-server ./cmd/server
```

2. Run the server:
```bash
./bookmarks-server
```

The server will start on port 8080 (or the configured PORT).

## API Endpoints

### Create Bookmark
```http
POST /api/bookmarks
Content-Type: application/json

{
    "url": "https://example.com"
}
```

### List Bookmarks
```http
GET /api/bookmarks
```

### Get Bookmark
```http
GET /api/bookmarks/{id}
```

### Delete Bookmark
```http
DELETE /api/bookmarks/{id}
```

## Testing with curl

Here are some example curl commands to test the API endpoints:

### Create a new bookmark
```bash
curl -i -X POST \
  -H "Content-Type: application/json" \
  -d '{"url":"https://example.com"}' \
  http://localhost:8081/api/bookmarks
```

### List all bookmarks
```bash
curl -i -X GET http://localhost:8081/api/bookmarks
```

### Get a specific bookmark (replace {id} with actual bookmark ID)
```bash
curl -i -X GET http://localhost:8081/api/bookmarks/1
```

### Delete a bookmark (replace {id} with actual bookmark ID)
```bash
curl -i -X DELETE http://localhost:8081/api/bookmarks/1
```

Note: The `-i` flag is included to show the response headers along with the body.

## Development

The project follows a standard Go project layout:

- `cmd/server`: Main application entry point
- `internal/api`: HTTP handlers and routing
- `internal/models`: Data models
- `internal/scraper`: Webpage metadata scraping
- `internal/storage`: Database operations
- `migrations`: SQL migration files

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
- Rate limiting (TODO)
- Request timeouts

## Performance

- Connection pooling for database
- Proper indexing on database tables
- Configurable timeouts for HTTP operations
- Graceful shutdown handling