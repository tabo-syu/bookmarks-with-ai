# Bookmarks Application Architecture

## System Overview

The Bookmarks Application is a web-based system that allows users to save and manage website bookmarks by entering URLs. The system automatically retrieves and stores website metadata to provide rich bookmark information.

### Key Features
- URL submission via form input
- Automatic metadata retrieval (title, description, favicon, etc.)
- Persistent storage of bookmark data
- Display of saved bookmarks with metadata

## Backend Architecture (Go)

### Components
1. **HTTP Server**
   - Built using standard Go `net/http` package
   - RESTful API endpoints for bookmark operations
   - CORS middleware for frontend communication

2. **Metadata Scraper**
   - HTTP client for fetching webpage content
   - HTML parser for extracting metadata (title, description, og:tags)
   - Favicon retrieval and processing

3. **Database Layer**
   - SQL database interface
   - Repository pattern for data access
   - Models for bookmark and metadata

### Directory Structure
```
backend/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── api/
│   │   ├── handlers/
│   │   └── routes.go
│   ├── models/
│   │   └── bookmark.go
│   ├── scraper/
│   │   └── metadata.go
│   └── storage/
│       └── repository.go
└── go.mod
```

## Frontend Architecture (TypeScript + Next.js)

### Components
1. **Page Components**
   - Home page with bookmark form and list
   - Individual bookmark view/edit pages

2. **UI Components**
   - Bookmark submission form
   - Bookmark card/list components
   - Loading states and error handling

3. **State Management**
   - React Query for server state
   - Form state handling
   - Loading and error states

### Directory Structure
```
frontend/
├── src/
│   ├── app/
│   │   ├── page.tsx
│   │   └── layout.tsx
│   ├── components/
│   │   ├── BookmarkForm/
│   │   └── BookmarkList/
│   └── lib/
│       ├── api/
│       └── types/
├── package.json
└── tsconfig.json
```

## Database Design

### Tables

#### bookmarks
```sql
CREATE TABLE bookmarks (
    id SERIAL PRIMARY KEY,
    url TEXT NOT NULL,
    title TEXT,
    description TEXT,
    favicon_url TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
```

## API Design

### Endpoints

#### POST /api/bookmarks
Create a new bookmark
```json
Request:
{
    "url": "https://example.com"
}

Response:
{
    "id": 1,
    "url": "https://example.com",
    "title": "Example Website",
    "description": "This is an example website",
    "favicon_url": "https://example.com/favicon.ico",
    "created_at": "2025-02-04T00:12:46Z",
    "updated_at": "2025-02-04T00:12:46Z"
}
```

#### GET /api/bookmarks
Retrieve all bookmarks
```json
Response:
{
    "bookmarks": [
        {
            "id": 1,
            "url": "https://example.com",
            "title": "Example Website",
            "description": "This is an example website",
            "favicon_url": "https://example.com/favicon.ico",
            "created_at": "2025-02-04T00:12:46Z",
            "updated_at": "2025-02-04T00:12:46Z"
        }
    ]
}
```

#### GET /api/bookmarks/:id
Retrieve a specific bookmark
```json
Response:
{
    "id": 1,
    "url": "https://example.com",
    "title": "Example Website",
    "description": "This is an example website",
    "favicon_url": "https://example.com/favicon.ico",
    "created_at": "2025-02-04T00:12:46Z",
    "updated_at": "2025-02-04T00:12:46Z"
}
```

#### DELETE /api/bookmarks/:id
Delete a bookmark
```json
Response:
{
    "success": true
}
```

## Security Considerations

1. Input Validation
   - URL validation and sanitization
   - Maximum length constraints
   - Content type verification

2. Rate Limiting
   - API endpoint rate limiting
   - Scraper request throttling

3. Error Handling
   - Graceful error responses
   - Logging and monitoring
   - Timeout handling for scraper

## Performance Considerations

1. Caching
   - Response caching for frequently accessed bookmarks
   - Metadata caching to prevent unnecessary scraping

2. Database Optimization
   - Proper indexing
   - Connection pooling
   - Query optimization

3. Frontend Optimization
   - Image optimization
   - Lazy loading
   - Bundle size optimization