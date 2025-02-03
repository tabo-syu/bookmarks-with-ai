package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"bookmarks-go/internal/models"
	"bookmarks-go/internal/scraper"
	"bookmarks-go/internal/storage"

	"github.com/gorilla/mux"
)

// BookmarkHandler handles bookmark-related HTTP requests
type BookmarkHandler struct {
	repo    storage.Repository
	scraper *scraper.Scraper
}

// NewBookmarkHandler creates a new bookmark handler
func NewBookmarkHandler(repo storage.Repository) *BookmarkHandler {
	return &BookmarkHandler{
		repo:    repo,
		scraper: scraper.NewScraper(10 * time.Second),
	}
}

// CreateBookmark handles the creation of a new bookmark
func (h *BookmarkHandler) CreateBookmark(w http.ResponseWriter, r *http.Request) {
	var req models.CreateBookmarkRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Fetch metadata
	metadata, err := h.scraper.GetMetadata(r.Context(), req.URL)
	if err != nil {
		http.Error(w, "Failed to fetch metadata: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Create bookmark
	bookmark := &models.Bookmark{
		URL:         req.URL,
		Title:       metadata.Title,
		Description: metadata.Description,
		FaviconURL:  metadata.FaviconURL,
	}

	if err := h.repo.CreateBookmark(r.Context(), bookmark); err != nil {
		http.Error(w, "Failed to create bookmark: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.BookmarkResponse{Bookmark: bookmark})
}

// GetBookmark handles retrieving a single bookmark
func (h *BookmarkHandler) GetBookmark(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid bookmark ID", http.StatusBadRequest)
		return
	}

	bookmark, err := h.repo.GetBookmark(r.Context(), id)
	if err != nil {
		if err == storage.ErrNotFound {
			http.Error(w, "Bookmark not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to get bookmark: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.BookmarkResponse{Bookmark: bookmark})
}

// ListBookmarks handles retrieving all bookmarks
func (h *BookmarkHandler) ListBookmarks(w http.ResponseWriter, r *http.Request) {
	bookmarks, err := h.repo.ListBookmarks(r.Context())
	if err != nil {
		http.Error(w, "Failed to list bookmarks: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.BookmarksResponse{Bookmarks: bookmarks})
}

// DeleteBookmark handles deleting a bookmark
func (h *BookmarkHandler) DeleteBookmark(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid bookmark ID", http.StatusBadRequest)
		return
	}

	err = h.repo.DeleteBookmark(r.Context(), id)
	if err != nil {
		if err == storage.ErrNotFound {
			http.Error(w, "Bookmark not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to delete bookmark: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.DeleteResponse{Success: true})
}
