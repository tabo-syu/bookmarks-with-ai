# Bookmarks Application

A full-stack application for managing website bookmarks with automatic metadata scraping.

## Project Structure

The project is divided into two main parts:

### Backend (/backend)

A Go backend service providing:
- RESTful API for bookmark management
- Automatic metadata extraction (title, description, favicon)
- PostgreSQL database storage
- OpenAPI specification

### Frontend (/frontend)

A Next.js frontend application providing:
- Modern, responsive user interface
- Real-time bookmark management
- Automatic metadata display
- TypeScript and React Query integration

## Getting Started

### Backend Setup

1. Navigate to the backend directory:
```bash
cd backend
```

2. Start the database:
```bash
docker-compose up -d postgres
```

3. Apply database migrations:
```bash
docker exec -i bookmarks_db psql -U postgres -d bookmarks_db < migrations/001_create_bookmarks_table.sql
```

4. Start the server:
```bash
go run cmd/server/main.go
```

The backend API will be available at http://localhost:8081/api

### Frontend Setup

1. Navigate to the frontend directory:
```bash
cd frontend
```

2. Install dependencies:
```bash
npm install
```

3. Generate API client code:
```bash
npm run generate-api
```

4. Start the development server:
```bash
npm run dev
```

The frontend application will be available at http://localhost:3000

## Development

### Backend

- `cmd/server`: Main application entry point
- `internal/api`: HTTP handlers and routing
- `internal/models`: Data models
- `internal/scraper`: Webpage metadata scraping
- `internal/storage`: Database operations
- `migrations`: SQL migration files

### Frontend

- `src/components`: React components
- `src/api`: Generated API client code
- `src/app`: Next.js application files

## API Documentation

The API is documented using OpenAPI 3.0. The specification is available at `backend/openapi.yaml`.

## Error Handling

Both frontend and backend implement proper error handling:

### Backend
- Appropriate HTTP status codes
- Structured error responses
- Input validation
- Database error handling

### Frontend
- Loading states
- Error messages
- Form validation
- Network error handling

## Security

- CORS configuration
- Input sanitization
- Prepared SQL statements
- Request timeouts
- Secure HTTP headers