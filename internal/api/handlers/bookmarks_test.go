package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"bookmarks-go/internal/models"
	"bookmarks-go/internal/storage"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRepository is a mock implementation of storage.Repository
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreateBookmark(ctx context.Context, bookmark *models.Bookmark) error {
	args := m.Called(ctx, bookmark)
	return args.Error(0)
}

func (m *MockRepository) GetBookmark(ctx context.Context, id int64) (*models.Bookmark, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Bookmark), args.Error(1)
}

func (m *MockRepository) ListBookmarks(ctx context.Context) ([]models.Bookmark, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.Bookmark), args.Error(1)
}

func (m *MockRepository) DeleteBookmark(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestCreateBookmark(t *testing.T) {
	mockRepo := new(MockRepository)
	handler := NewBookmarkHandler(mockRepo)

	tests := []struct {
		name           string
		requestBody    interface{}
		setupMock      func()
		expectedStatus int
		expectedError  string
	}{
		{
			name: "successful creation",
			requestBody: models.CreateBookmarkRequest{
				URL: "https://example.com",
			},
			setupMock: func() {
				mockRepo.On("CreateBookmark", mock.Anything, mock.MatchedBy(func(b *models.Bookmark) bool {
					return b.URL == "https://example.com"
				})).Return(nil)
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "invalid request body",
			requestBody:    "invalid",
			setupMock:      func() {},
			expectedStatus: http.StatusBadRequest,
			expectedError:  "Invalid request body\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()

			body, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest("POST", "/bookmarks", bytes.NewBuffer(body))
			w := httptest.NewRecorder()

			handler.CreateBookmark(w, req)

			resp := w.Result()
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			if tt.expectedError != "" {
				bodyBytes, _ := io.ReadAll(resp.Body)
				assert.Equal(t, tt.expectedError, string(bodyBytes))
			} else {
				var response models.BookmarkResponse
				json.NewDecoder(resp.Body).Decode(&response)
				assert.NotNil(t, response.Bookmark)
				assert.Equal(t, "https://example.com", response.Bookmark.URL)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestGetBookmark(t *testing.T) {
	mockRepo := new(MockRepository)
	handler := NewBookmarkHandler(mockRepo)

	bookmark := &models.Bookmark{
		ID:        1,
		URL:       "https://example.com",
		Title:     "Example",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tests := []struct {
		name           string
		bookmarkID     string
		setupMock      func()
		expectedStatus int
		expectedError  string
	}{
		{
			name:       "successful retrieval",
			bookmarkID: "1",
			setupMock: func() {
				mockRepo.On("GetBookmark", mock.Anything, int64(1)).Return(bookmark, nil)
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:       "not found",
			bookmarkID: "999",
			setupMock: func() {
				mockRepo.On("GetBookmark", mock.Anything, int64(999)).Return(nil, storage.ErrNotFound)
			},
			expectedStatus: http.StatusNotFound,
			expectedError:  "Bookmark not found\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()

			req := httptest.NewRequest("GET", "/bookmarks/"+tt.bookmarkID, nil)
			req = mux.SetURLVars(req, map[string]string{"id": tt.bookmarkID})
			w := httptest.NewRecorder()

			handler.GetBookmark(w, req)

			resp := w.Result()
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			if tt.expectedError != "" {
				bodyBytes, _ := io.ReadAll(resp.Body)
				assert.Equal(t, tt.expectedError, string(bodyBytes))
			} else {
				var response models.BookmarkResponse
				json.NewDecoder(resp.Body).Decode(&response)
				assert.Equal(t, bookmark.ID, response.Bookmark.ID)
				assert.Equal(t, bookmark.URL, response.Bookmark.URL)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestListBookmarks(t *testing.T) {
	mockRepo := new(MockRepository)
	handler := NewBookmarkHandler(mockRepo)

	bookmarks := []models.Bookmark{
		{
			ID:  1,
			URL: "https://example1.com",
		},
		{
			ID:  2,
			URL: "https://example2.com",
		},
	}

	tests := []struct {
		name           string
		setupMock      func()
		expectedStatus int
	}{
		{
			name: "successful listing",
			setupMock: func() {
				mockRepo.On("ListBookmarks", mock.Anything).Return(bookmarks, nil)
			},
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()

			req := httptest.NewRequest("GET", "/bookmarks", nil)
			w := httptest.NewRecorder()

			handler.ListBookmarks(w, req)

			resp := w.Result()
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			var response models.BookmarksResponse
			json.NewDecoder(resp.Body).Decode(&response)
			assert.Equal(t, len(bookmarks), len(response.Bookmarks))

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestDeleteBookmark(t *testing.T) {
	mockRepo := new(MockRepository)
	handler := NewBookmarkHandler(mockRepo)

	tests := []struct {
		name           string
		bookmarkID     string
		setupMock      func()
		expectedStatus int
		expectedError  string
	}{
		{
			name:       "successful deletion",
			bookmarkID: "1",
			setupMock: func() {
				mockRepo.On("DeleteBookmark", mock.Anything, int64(1)).Return(nil)
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:       "not found",
			bookmarkID: "999",
			setupMock: func() {
				mockRepo.On("DeleteBookmark", mock.Anything, int64(999)).Return(storage.ErrNotFound)
			},
			expectedStatus: http.StatusNotFound,
			expectedError:  "Bookmark not found\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()

			req := httptest.NewRequest("DELETE", "/bookmarks/"+tt.bookmarkID, nil)
			req = mux.SetURLVars(req, map[string]string{"id": tt.bookmarkID})
			w := httptest.NewRecorder()

			handler.DeleteBookmark(w, req)

			resp := w.Result()
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			if tt.expectedError != "" {
				bodyBytes, _ := io.ReadAll(resp.Body)
				assert.Equal(t, tt.expectedError, string(bodyBytes))
			} else {
				var response models.DeleteResponse
				json.NewDecoder(resp.Body).Decode(&response)
				assert.True(t, response.Success)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
