package user

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	db.Exec(`CREATE TABLE users (
		id text PRIMARY KEY,
		name text NOT NULL,
		email text NOT NULL UNIQUE,
		picture text,
		sso_provider text NOT NULL,
		sso_id text NOT NULL UNIQUE,
		created_at datetime,
		updated_at datetime
	)`)

	DB = db
	return db
}

func TestSingInWithJWT(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db := setupTestDB(t)

	r := gin.New()
	r.GET("/singin-jwt", func(c *gin.Context) {
		userIDStr := c.Query("userID")
		if userIDStr != "" {
			u, err := uuid.Parse(userIDStr)
			if err == nil {
				c.Set("userID", u.String())
			}
		}
		c.Next()
	}, SingInWithJWT)

	// Create a test user in DB
	testUser := User{
		ID:          uuid.New(),
		Name:        "Alice",
		Email:       "alice@example.com",
		SSOProvider: "google",
		SSOID:       "google-oauth2|alice123",
	}
	if err := db.Create(&testUser).Error; err != nil {
		t.Fatalf("failed to seed user: %v", err)
	}

	t.Run("User Found", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/singin-jwt?userID="+testUser.ID.String(), nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d. Response: %s", w.Code, w.Body.String())
		}
	})

	t.Run("User Not Found", func(t *testing.T) {
		randomID := uuid.New().String()
		req, _ := http.NewRequest(http.MethodGet, "/singin-jwt?userID="+randomID, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusInternalServerError {
			t.Errorf("expected status 500 (standard response error), got %d. Response: %s", w.Code, w.Body.String())
		}
	})
}
