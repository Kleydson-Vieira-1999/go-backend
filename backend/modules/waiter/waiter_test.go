package waiter

import (
	"testing"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	// Manually create tables to avoid PostgreSQL-specific DEFAULT gen_random_uuid() syntax in SQLite
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

	db.Exec(`CREATE TABLE waiter_codes (
		id text PRIMARY KEY,
		store_id text NOT NULL,
		code text NOT NULL,
		label text NOT NULL,
		is_active boolean,
		updated_at text
	)`)

	DB = db
	return db
}

func TestWaiterCodeCRUD(t *testing.T) {
	db := setupTestDB(t)

	// 1. Create a Store ID
	storeID := uuid.New()

	// 2. Create a WaiterCode
	waiterCodeID := uuid.New()
	wCode := WaiterCode{
		ID:        waiterCodeID,
		StoreID:   storeID,
		Code:      "W1234",
		Label:     "Table Waiter 1",
		IsActive:  true,
		UpdatedAt: "2026-06-27T14:12:49-03:00",
	}

	if err := db.Create(&wCode).Error; err != nil {
		t.Fatalf("failed to create waiter code: %v", err)
	}

	// 3. Read WaiterCode (Get by ID)
	var fetched WaiterCode
	if err := db.First(&fetched, "id = ?", waiterCodeID).Error; err != nil {
		t.Fatalf("failed to find waiter code: %v", err)
	}

	if fetched.Code != "W1234" {
		t.Errorf("expected code to be W1234, got %s", fetched.Code)
	}
	if !fetched.IsActive {
		t.Errorf("expected IsActive to be true")
	}

	// 4. Update WaiterCode
	fetched.Label = "Updated Label"
	fetched.IsActive = false
	if err := db.Save(&fetched).Error; err != nil {
		t.Fatalf("failed to update waiter code: %v", err)
	}

	var updated WaiterCode
	if err := db.First(&updated, "id = ?", waiterCodeID).Error; err != nil {
		t.Fatalf("failed to find updated waiter code: %v", err)
	}
	if updated.Label != "Updated Label" {
		t.Errorf("expected updated label to be 'Updated Label', got %s", updated.Label)
	}

	// 5. Delete WaiterCode
	if err := db.Delete(&updated).Error; err != nil {
		t.Fatalf("failed to delete waiter code: %v", err)
	}
}

func TestWaiterCodeTableName(t *testing.T) {
	wCode := WaiterCode{}
	if wCode.TableName() != "waiter_codes" {
		t.Errorf("expected TableName to be 'waiter_codes', got '%s'", wCode.TableName())
	}
}
