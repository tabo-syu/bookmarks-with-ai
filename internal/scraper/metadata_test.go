package scraper

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetMetadata(t *testing.T) {
	tests := []struct {
		name        string
		html        string
		wantTitle   string
		wantDesc    string
		wantFavicon string
	}{
		{
			name: "complete metadata",
			html: `
				<!DOCTYPE html>
				<html>
				<head>
					<title>Test Title</title>
					<meta name="description" content="Test Description">
					<link rel="icon" href="/favicon.ico">
				</head>
				<body>Test content</body>
				</html>
			`,
			wantTitle:   "Test Title",
			wantDesc:    "Test Description",
			wantFavicon: "/favicon.ico",
		},
		{
			name: "only title",
			html: `
				<!DOCTYPE html>
				<html>
				<head>
					<title>Test Title</title>
				</head>
				<body>Test content</body>
				</html>
			`,
			wantTitle:   "Test Title",
			wantDesc:    "",
			wantFavicon: "",
		},
		{
			name: "og metadata",
			html: `
				<!DOCTYPE html>
				<html>
				<head>
					<meta property="og:title" content="OG Title">
					<meta property="og:description" content="OG Description">
					<link rel="shortcut icon" href="/favicon.png">
				</head>
				<body>Test content</body>
				</html>
			`,
			wantTitle:   "OG Title",
			wantDesc:    "OG Description",
			wantFavicon: "/favicon.png",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create test server
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "text/html")
				w.Write([]byte(tt.html))
			}))
			defer ts.Close()

			// Create scraper with timeout
			scraper := NewScraper(5 * time.Second)

			// Extract metadata
			metadata, err := scraper.GetMetadata(context.Background(), ts.URL)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantTitle, metadata.Title)
			assert.Equal(t, tt.wantDesc, metadata.Description)

			// For favicon, we only check if it's set when expected
			if tt.wantFavicon != "" {
				assert.NotEmpty(t, metadata.FaviconURL)
			}
		})
	}
}

func TestGetMetadataError(t *testing.T) {
	scraper := NewScraper(5 * time.Second)
	tests := []struct {
		name    string
		url     string
		wantErr bool
	}{
		{
			name:    "invalid url",
			url:     "invalid-url",
			wantErr: true,
		},
		{
			name:    "non-existent domain",
			url:     "http://non-existent-domain.test",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := scraper.GetMetadata(context.Background(), tt.url)
			assert.Error(t, err)
		})
	}
}
