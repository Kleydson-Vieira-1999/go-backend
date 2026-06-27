package auth

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func writeDummySecretFile() error {
	secretContent := `{
		"web": {
			"client_id": "test-client-id",
			"project_id": "test-project",
			"auth_uri": "https://accounts.google.com/o/oauth2/auth",
			"token_uri": "https://oauth2.googleapis.com/token",
			"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
			"client_secret": "test-client-secret",
			"redirect_uris": ["http://localhost:8080/callback"]
		}
	}`
	return os.WriteFile("client_secret_apps.googleusercontent.com.json", []byte(secretContent), 0644)
}

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

func TestGoogleAuthGetAuthURL(t *testing.T) {
	err := writeDummySecretFile()
	if err != nil {
		t.Fatalf("failed to write dummy secret file: %v", err)
	}
	defer os.Remove("client_secret_apps.googleusercontent.com.json")

	_ = setupTestDB(t)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	handler := NewGoogleAuthHandler()
	handler.GetAuthURL(c)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}

func TestPostAuthToken_InvalidJSON(t *testing.T) {
	err := writeDummySecretFile()
	if err != nil {
		t.Fatalf("failed to write dummy secret file: %v", err)
	}
	defer os.Remove("client_secret_apps.googleusercontent.com.json")

	_ = setupTestDB(t)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest("POST", "/api/google/auth/token", bytes.NewBufferString("invalid json"))
	c.Request.Header.Set("Content-Type", "application/json")

	handler := NewGoogleAuthHandler()
	handler.PostAuthToken(c)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected status 500, got %d", w.Code)
	}
}
