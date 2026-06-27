package store

import (
	"bytes"
	"encoding/json"
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
		t.Fatalf("Failed to open test database: %v", err)
	}

	db.Exec(`CREATE TABLE stores (
		id text PRIMARY KEY,
		user_id text NOT NULL,
		store_template_id text,
		name text NOT NULL,
		picture text,
		type text,
		description text,
		is_active boolean,
		created_at datetime,
		updated_at datetime
	)`)

	db.Exec(`CREATE TABLE store_balance (
		id text PRIMARY KEY,
		store_id text NOT NULL UNIQUE,
		current_balance integer,
		total_profit integer,
		updated_at datetime
	)`)

	DB = db
	return db
}

func setupTestRouter(userID string) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	// Middleware to set userID in context
	r.Use(func(c *gin.Context) {
		if userID != "" {
			c.Set("userID", userID)
		}
		c.Next()
	})

	r.GET("/stores", ListAllStores)
	r.GET("/stores/:id", GetStoreByID)
	r.POST("/stores", CreateStore)
	r.PATCH("/stores/:id", UpdateStore)

	return r
}

func TestListAllStores(t *testing.T) {
	db := setupTestDB(t)

	userID := uuid.New()
	otherUserID := uuid.New()

	store1 := Store{
		ID:     uuid.New(),
		UserID: userID,
		Name:   "Store 1",
	}
	store2 := Store{
		ID:     uuid.New(),
		UserID: userID,
		Name:   "Store 2",
	}
	storeOther := Store{
		ID:     uuid.New(),
		UserID: otherUserID,
		Name:   "Other Store",
	}

	db.Create(&store1)
	db.Create(&store2)
	db.Create(&storeOther)

	t.Run("Success - List all stores for user", func(t *testing.T) {
		r := setupTestRouter(userID.String())
		req, _ := http.NewRequest(http.MethodGet, "/stores", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("expected 200, got %d", w.Code)
		}

		var resp map[string][]Store
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		stores := resp["stores"]
		if len(stores) != 2 {
			t.Errorf("expected 2 stores, got %d", len(stores))
		}
	})

	t.Run("Error - Database error", func(t *testing.T) {
		err := db.Migrator().DropTable(&Store{})
		if err != nil {
			t.Fatalf("failed to drop table: %v", err)
		}

		r := setupTestRouter(userID.String())
		req, _ := http.NewRequest(http.MethodGet, "/stores", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusInternalServerError {
			t.Errorf("expected 500, got %d", w.Code)
		}
	})
}

func TestGetStoreByID(t *testing.T) {
	db := setupTestDB(t)

	userID := uuid.New()
	storeID := uuid.New()

	store := Store{
		ID:     storeID,
		UserID: userID,
		Name:   "My Store",
	}
	db.Create(&store)

	t.Run("Success - Get store by ID", func(t *testing.T) {
		r := setupTestRouter(userID.String())
		req, _ := http.NewRequest(http.MethodGet, "/stores/"+storeID.String(), nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("expected 200, got %d", w.Code)
		}

		var resp map[string]Store
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		if resp["store"].ID != storeID {
			t.Errorf("expected store ID %v, got %v", storeID, resp["store"].ID)
		}
	})
}

func TestCreateStore(t *testing.T) {
	_ = setupTestDB(t)

	userID := uuid.New()

	t.Run("Success - Create store", func(t *testing.T) {
		r := setupTestRouter(userID.String())
		newStore := map[string]interface{}{
			"name":        "Brand New Store",
			"description": "Store Description",
		}
		body, _ := json.Marshal(newStore)
		req, _ := http.NewRequest(http.MethodPost, "/stores", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusCreated {
			t.Errorf("expected 201, got %d. Response: %s", w.Code, w.Body.String())
		}
	})

	t.Run("Error - Invalid payload JSON", func(t *testing.T) {
		r := setupTestRouter(userID.String())
		req, _ := http.NewRequest(http.MethodPost, "/stores", bytes.NewBufferString("{invalid-json}"))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusInternalServerError {
			t.Errorf("expected 500, got %d", w.Code)
		}
	})
}

func TestUpdateStore(t *testing.T) {
	db := setupTestDB(t)

	userID := uuid.New()
	storeID := uuid.New()

	store := Store{
		ID:     storeID,
		UserID: userID,
		Name:   "Original Name",
	}
	db.Create(&store)

	t.Run("Success - Update store", func(t *testing.T) {
		r := setupTestRouter(userID.String())
		updatedPayload := map[string]interface{}{
			"name": "Updated Name",
		}
		body, _ := json.Marshal(updatedPayload)
		req, _ := http.NewRequest(http.MethodPatch, "/stores/"+storeID.String(), bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("expected 200, got %d", w.Code)
		}
	})
}
