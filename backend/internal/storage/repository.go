package storage

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"bookmarks-go/internal/models"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	ErrNotFound = errors.New("bookmark not found")
	ErrDatabase = errors.New("database error")
)

// Repository defines the interface for bookmark storage operations
type Repository interface {
	CreateBookmark(ctx context.Context, bookmark *models.Bookmark) error
	GetBookmark(ctx context.Context, id int64) (*models.Bookmark, error)
	ListBookmarks(ctx context.Context) ([]models.Bookmark, error)
	DeleteBookmark(ctx context.Context, id int64) error
}

// PostgresRepository implements Repository interface for PostgreSQL
type PostgresRepository struct {
	db *sqlx.DB
}

// NewPostgresRepository creates a new PostgreSQL repository
func NewPostgresRepository(db *sqlx.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

// CreateBookmark inserts a new bookmark into the database
func (r *PostgresRepository) CreateBookmark(ctx context.Context, bookmark *models.Bookmark) error {
	query := `
		INSERT INTO bookmarks (url, title, description, favicon_url, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`

	now := time.Now().UTC()
	bookmark.CreatedAt = now
	bookmark.UpdatedAt = now

	err := r.db.QueryRowxContext(
		ctx,
		query,
		bookmark.URL,
		bookmark.Title,
		bookmark.Description,
		bookmark.FaviconURL,
		bookmark.CreatedAt,
		bookmark.UpdatedAt,
	).Scan(&bookmark.ID)

	if err != nil {
		return errors.New("failed to create bookmark: " + err.Error())
	}

	return nil
}

// GetBookmark retrieves a bookmark by ID
func (r *PostgresRepository) GetBookmark(ctx context.Context, id int64) (*models.Bookmark, error) {
	bookmark := &models.Bookmark{}
	query := `
		SELECT id, url, title, description, favicon_url, created_at, updated_at
		FROM bookmarks
		WHERE id = $1`

	err := r.db.GetContext(ctx, bookmark, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, errors.New("failed to get bookmark: " + err.Error())
	}

	return bookmark, nil
}

// ListBookmarks retrieves all bookmarks
func (r *PostgresRepository) ListBookmarks(ctx context.Context) ([]models.Bookmark, error) {
	var bookmarks []models.Bookmark
	query := `
		SELECT id, url, title, description, favicon_url, created_at, updated_at
		FROM bookmarks
		ORDER BY created_at DESC`

	err := r.db.SelectContext(ctx, &bookmarks, query)
	if err != nil {
		return nil, errors.New("failed to list bookmarks: " + err.Error())
	}

	return bookmarks, nil
}

// DeleteBookmark removes a bookmark by ID
func (r *PostgresRepository) DeleteBookmark(ctx context.Context, id int64) error {
	query := `DELETE FROM bookmarks WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return errors.New("failed to delete bookmark: " + err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.New("failed to get rows affected: " + err.Error())
	}

	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}
