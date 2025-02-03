# Bookmarks Application

A full-stack application for managing website bookmarks with automatic metadata scraping.

## Project Structure

The project is organized into three main parts:

### API Specification (Root)

- `openapi.yaml`: OpenAPI 3.0 specification defining the REST API interface
  - Used by backend for API implementation
  - Used by frontend to generate type-safe API client

### Backend (/backend)

A Go backend service providing:
- RESTful API implementation following OpenAPI spec
- Automatic metadata extraction (title, description, favicon)
- PostgreSQL database storage
- Comprehensive error handling

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

3. Generate API client code (uses root openapi.yaml):
```bash
npm run generate-api
```

4. Start the development server:
```bash
npm run dev
```

The frontend application will be available at http://localhost:3000

## Development

### API First Development

1. Update the OpenAPI specification (`openapi.yaml`) in the root directory
2. Implement the changes in the backend
3. Generate the frontend API client to get updated types and endpoints

### Backend Development

- `cmd/server`: Main application entry point
- `internal/api`: HTTP handlers and routing
- `internal/models`: Data models
- `internal/scraper`: Webpage metadata scraping
- `internal/storage`: Database operations
- `migrations`: SQL migration files

### Frontend Development

- `src/components`: React components
- `src/api`: Generated API client code
- `src/app`: Next.js application files

## Documentation

- Root `README.md`: Project overview and setup
- `backend/README.md`: Detailed backend documentation
- `frontend/README.md`: Detailed frontend documentation
- `openapi.yaml`: API specification and documentation

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