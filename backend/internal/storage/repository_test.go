package storage

import (
	"context"
	"testing"

	"bookmarks-go/internal/models"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
)

type RepositoryTestSuite struct {
	suite.Suite
	db         *sqlx.DB
	repository Repository
}

func (s *RepositoryTestSuite) SetupSuite() {
	connStr := "host=localhost port=5433 user=postgres password=postgres dbname=bookmarks_test_db sslmode=disable"
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		s.T().Fatalf("Failed to connect to test database: %v", err)
	}
	s.db = db
	s.repository = NewPostgresRepository(db)

	// Create tables
	schema := `
		CREATE TABLE IF NOT EXISTS bookmarks (
			id SERIAL PRIMARY KEY,
			url TEXT NOT NULL,
			title TEXT,
			description TEXT,
			favicon_url TEXT,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL
		);
	`
	_, err = db.Exec(schema)
	if err != nil {
		s.T().Fatalf("Failed to create test tables: %v", err)
	}
}

func (s *RepositoryTestSuite) TearDownSuite() {
	if s.db != nil {
		_, err := s.db.Exec("DROP TABLE IF EXISTS bookmarks")
		if err != nil {
			s.T().Errorf("Failed to drop test tables: %v", err)
		}
		s.db.Close()
	}
}

func (s *RepositoryTestSuite) SetupTest() {
	_, err := s.db.Exec("TRUNCATE TABLE bookmarks RESTART IDENTITY")
	if err != nil {
		s.T().Fatalf("Failed to truncate test tables: %v", err)
	}
}

func (s *RepositoryTestSuite) TestCreateBookmark() {
	bookmark := &models.Bookmark{
		URL:         "https://example.com",
		Title:       "Example",
		Description: "An example website",
		FaviconURL:  "https://example.com/favicon.ico",
	}

	err := s.repository.CreateBookmark(context.Background(), bookmark)
	s.NoError(err)
	s.NotZero(bookmark.ID)
	s.NotZero(bookmark.CreatedAt)
	s.NotZero(bookmark.UpdatedAt)
}

func (s *RepositoryTestSuite) TestGetBookmark() {
	// Create a bookmark first
	bookmark := &models.Bookmark{
		URL:         "https://example.com",
		Title:       "Example",
		Description: "An example website",
		FaviconURL:  "https://example.com/favicon.ico",
	}
	err := s.repository.CreateBookmark(context.Background(), bookmark)
	s.NoError(err)

	// Test getting the bookmark
	retrieved, err := s.repository.GetBookmark(context.Background(), bookmark.ID)
	s.NoError(err)
	s.Equal(bookmark.URL, retrieved.URL)
	s.Equal(bookmark.Title, retrieved.Title)
	s.Equal(bookmark.Description, retrieved.Description)
	s.Equal(bookmark.FaviconURL, retrieved.FaviconURL)
}

func (s *RepositoryTestSuite) TestGetBookmarkNotFound() {
	_, err := s.repository.GetBookmark(context.Background(), 999)
	s.Equal(ErrNotFound, err)
}

func (s *RepositoryTestSuite) TestListBookmarks() {
	// Create multiple bookmarks
	bookmarks := []models.Bookmark{
		{URL: "https://example1.com", Title: "Example 1"},
		{URL: "https://example2.com", Title: "Example 2"},
		{URL: "https://example3.com", Title: "Example 3"},
	}

	for i := range bookmarks {
		err := s.repository.CreateBookmark(context.Background(), &bookmarks[i])
		s.NoError(err)
	}

	// Test listing bookmarks
	list, err := s.repository.ListBookmarks(context.Background())
	s.NoError(err)
	s.Len(list, len(bookmarks))
}

func (s *RepositoryTestSuite) TestDeleteBookmark() {
	// Create a bookmark first
	bookmark := &models.Bookmark{
		URL:   "https://example.com",
		Title: "Example",
	}
	err := s.repository.CreateBookmark(context.Background(), bookmark)
	s.NoError(err)

	// Test deleting the bookmark
	err = s.repository.DeleteBookmark(context.Background(), bookmark.ID)
	s.NoError(err)

	// Verify it's deleted
	_, err = s.repository.GetBookmark(context.Background(), bookmark.ID)
	s.Equal(ErrNotFound, err)
}

func (s *RepositoryTestSuite) TestDeleteBookmarkNotFound() {
	err := s.repository.DeleteBookmark(context.Background(), 999)
	s.Equal(ErrNotFound, err)
}

func TestRepositorySuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}
