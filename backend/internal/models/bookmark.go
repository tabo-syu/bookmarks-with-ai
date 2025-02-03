package models

import "time"

// Bookmark represents a stored bookmark with metadata
type Bookmark struct {
	ID          int64     `json:"id" db:"id"`
	URL         string    `json:"url" db:"url"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	FaviconURL  string    `json:"favicon_url" db:"favicon_url"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// CreateBookmarkRequest represents the request body for creating a bookmark
type CreateBookmarkRequest struct {
	URL string `json:"url"`
}

// BookmarkResponse represents the response for bookmark endpoints
type BookmarkResponse struct {
	Bookmark *Bookmark `json:"bookmark,omitempty"`
	Error    string    `json:"error,omitempty"`
}

// BookmarksResponse represents the response for listing bookmarks
type BookmarksResponse struct {
	Bookmarks []Bookmark `json:"bookmarks"`
	Error     string     `json:"error,omitempty"`
}

// DeleteResponse represents the response for delete operation
type DeleteResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}
