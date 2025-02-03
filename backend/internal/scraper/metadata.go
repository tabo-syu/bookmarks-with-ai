package scraper

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/net/html"
)

// Metadata represents the scraped information from a webpage
type Metadata struct {
	Title       string
	Description string
	FaviconURL  string
}

// Scraper handles webpage metadata extraction
type Scraper struct {
	client *http.Client
}

// NewScraper creates a new metadata scraper with configured timeout
func NewScraper(timeout time.Duration) *Scraper {
	return &Scraper{
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

// GetMetadata fetches and extracts metadata from the given URL
func (s *Scraper) GetMetadata(ctx context.Context, urlStr string) (*Metadata, error) {
	// Validate URL
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %w", err)
	}

	if !strings.HasPrefix(parsedURL.Scheme, "http") {
		return nil, errors.New("URL must start with http:// or https://")
	}

	// Create request with context
	req, err := http.NewRequestWithContext(ctx, "GET", urlStr, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set user agent to avoid being blocked
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; BookmarksBot/1.0)")

	// Perform request
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch URL: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Parse HTML
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %w", err)
	}

	// Extract metadata
	metadata := &Metadata{}
	metadata.extractMetadata(doc, parsedURL)

	// If favicon not found in metadata, try default location
	if metadata.FaviconURL == "" {
		metadata.FaviconURL = s.findDefaultFavicon(parsedURL)
	}

	return metadata, nil
}

// extractMetadata traverses the HTML tree to find metadata
func (m *Metadata) extractMetadata(n *html.Node, baseURL *url.URL) {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "title":
			if m.Title == "" && n.FirstChild != nil {
				m.Title = n.FirstChild.Data
			}
		case "meta":
			var name, content string
			for _, attr := range n.Attr {
				switch attr.Key {
				case "name", "property":
					name = attr.Val
				case "content":
					content = attr.Val
				}
			}
			switch name {
			case "description", "og:description":
				if m.Description == "" {
					m.Description = content
				}
			case "og:title":
				if m.Title == "" {
					m.Title = content
				}
			}
		case "link":
			var rel, href string
			for _, attr := range n.Attr {
				switch attr.Key {
				case "rel":
					rel = attr.Val
				case "href":
					href = attr.Val
				}
			}
			if rel == "icon" || rel == "shortcut icon" {
				if href != "" {
					faviconURL, err := baseURL.Parse(href)
					if err == nil {
						m.FaviconURL = faviconURL.String()
					}
				}
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		m.extractMetadata(c, baseURL)
	}
}

// findDefaultFavicon attempts to find favicon at the default location
func (s *Scraper) findDefaultFavicon(baseURL *url.URL) string {
	defaultFaviconURL := *baseURL
	defaultFaviconURL.Path = "/favicon.ico"

	// Try to fetch favicon
	req, err := http.NewRequest("HEAD", defaultFaviconURL.String(), nil)
	if err != nil {
		return ""
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return defaultFaviconURL.String()
	}

	return ""
}
