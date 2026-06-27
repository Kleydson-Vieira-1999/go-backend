package order

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/store"
	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/table"
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

	db.Exec(`CREATE TABLE stores (
		id text PRIMARY KEY,
		user_id text NOT NULL,
		name text NOT NULL,
		picture text,
		type text,
		description text,
		is_active boolean,
		created_at datetime,
		updated_at datetime
	)`)

	db.Exec(`CREATE TABLE tables (
		id text PRIMARY KEY,
		store_id text NOT NULL,
		identifier text NOT NULL,
		is_active boolean,
		updated_at text
	)`)

	db.Exec(`CREATE TABLE table_sessions (
		id text PRIMARY KEY,
		store_id text NOT NULL,
		table_id text NOT NULL,
		status text NOT NULL,
		opened_at datetime,
		closed_at datetime,
		updated_at datetime
	)`)

	db.Exec(`CREATE TABLE waiter_codes (
		id text PRIMARY KEY,
		store_id text NOT NULL,
		code text NOT NULL,
		label text NOT NULL,
		is_active boolean,
		updated_at text
	)`)

	db.Exec(`CREATE TABLE orders (
		id text PRIMARY KEY,
		session_id text NOT NULL,
		waiter_code_id text,
		status text NOT NULL,
		notes text,
		created_at datetime,
		updated_at datetime
	)`)

	DB = db
	return db
}

func TestCreateOrder_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db := setupTestDB(t)

	broker := store.NewBroker()
	go broker.Start()

	storeID := uuid.New()
	st := store.Store{
		ID:        storeID,
		UserID:    uuid.New(),
		Name:      "Test Store",
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	db.Create(&st)

	tableID := uuid.New()
	tbl := table.Table{
		ID:         tableID,
		StoreID:    storeID,
		Identifier: "Table 1",
		IsActive:   true,
		UpdatedAt:  time.Now().Format(time.RFC3339),
	}
	db.Create(&tbl)

	sessionID := uuid.New()
	sess := table.TableSession{
		ID:        sessionID,
		StoreID:   storeID,
		TableID:   tableID,
		Status:    "active",
		OpenedAt:  time.Now(),
		UpdatedAt: time.Now(),
	}
	db.Create(&sess)

	client := &store.Client{
		EstablishmentID: sessionID.String(),
		MessageChan:     make(chan string, 1),
	}
	broker.Register(client)

	r := gin.New()
	r.POST("/api/orders", CreateOrder)

	reqPayload := map[string]interface{}{
		"session_id": sessionID.String(),
		"notes":      "Sem cebola",
	}
	body, _ := json.Marshal(reqPayload)

	req, _ := http.NewRequest(http.MethodPost, "/api/orders", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d. Body: %s", w.Code, w.Body.String())
	}

	var dbOrder Order
	if err := db.First(&dbOrder).Error; err != nil {
		t.Fatalf("Order not found in database: %v", err)
	}
}
