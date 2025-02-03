package api

import (
	"net/http"

	"bookmarks-go/internal/api/handlers"
	"bookmarks-go/internal/storage"

	"github.com/gorilla/mux"
)

// CORSMiddleware adds CORS headers to responses
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// SetupRoutes configures all API routes and middleware
func SetupRoutes(repo storage.Repository) *mux.Router {
	r := mux.NewRouter()

	// Create handlers
	bookmarkHandler := handlers.NewBookmarkHandler(repo)

	// API routes
	api := r.PathPrefix("/api").Subrouter()
	api.Use(CORSMiddleware)

	// Bookmark routes
	bookmarks := api.PathPrefix("/bookmarks").Subrouter()
	bookmarks.HandleFunc("", bookmarkHandler.CreateBookmark).Methods("POST")
	bookmarks.HandleFunc("", bookmarkHandler.ListBookmarks).Methods("GET")
	bookmarks.HandleFunc("/{id:[0-9]+}", bookmarkHandler.GetBookmark).Methods("GET")
	bookmarks.HandleFunc("/{id:[0-9]+}", bookmarkHandler.DeleteBookmark).Methods("DELETE")

	// Add OPTIONS method for CORS preflight requests
	bookmarks.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {}).Methods("OPTIONS")
	bookmarks.HandleFunc("/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {}).Methods("OPTIONS")

	return r
}
