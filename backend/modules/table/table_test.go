package table

import (
	"testing"
	"time"

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

	DB = db
	return db
}

func TestTableTableName(t *testing.T) {
	t.Run("Table name", func(t *testing.T) {
		tbl := Table{}
		if tbl.TableName() != "tables" {
			t.Errorf("expected 'tables', got '%s'", tbl.TableName())
		}
	})

	t.Run("TableSession name", func(t *testing.T) {
		ts := TableSession{}
		if ts.TableName() != "table_sessions" {
			t.Errorf("expected 'table_sessions', got '%s'", ts.TableName())
		}
	})
}

func TestTableCRUD(t *testing.T) {
	db := setupTestDB(t)

	storeID := uuid.New()

	// 1. Create Table
	tableID := uuid.New()
	tbl := Table{
		ID:         tableID,
		StoreID:    storeID,
		Identifier: "Table 1",
		IsActive:   true,
		UpdatedAt:  time.Now().Format(time.RFC3339),
	}

	if err := db.Create(&tbl).Error; err != nil {
		t.Fatalf("failed to create table: %v", err)
	}

	// 2. Read Table (Get by ID)
	var fetchedTable Table
	if err := db.First(&fetchedTable, "id = ?", tableID).Error; err != nil {
		t.Fatalf("failed to find table: %v", err)
	}

	if fetchedTable.Identifier != "Table 1" {
		t.Errorf("expected identifier 'Table 1', got '%s'", fetchedTable.Identifier)
	}

	// 3. Update Table
	fetchedTable.Identifier = "Table 1 Updated"
	fetchedTable.IsActive = false

	if err := db.Save(&fetchedTable).Error; err != nil {
		t.Fatalf("failed to update table: %v", err)
	}

	var updatedTable Table
	if err := db.First(&updatedTable, "id = ?", tableID).Error; err != nil {
		t.Fatalf("failed to fetch updated table: %v", err)
	}

	if updatedTable.IsActive {
		t.Errorf("expected table to be inactive after update")
	}

	// 4. Delete Table
	if err := db.Delete(&updatedTable).Error; err != nil {
		t.Fatalf("failed to delete table: %v", err)
	}
}

func TestTableSessionCRUD(t *testing.T) {
	db := setupTestDB(t)

	storeID := uuid.New()
	tableID := uuid.New()

	// 1. Create Table Session
	sessionID := uuid.New()
	now := time.Now()
	session := TableSession{
		ID:        sessionID,
		StoreID:   storeID,
		TableID:   tableID,
		Status:    "active",
		OpenedAt:  now,
		UpdatedAt: now,
	}

	if err := db.Create(&session).Error; err != nil {
		t.Fatalf("failed to create table session: %v", err)
	}

	// 2. Read Table Session
	var fetchedSession TableSession
	if err := db.First(&fetchedSession, "id = ?", sessionID).Error; err != nil {
		t.Fatalf("failed to find table session: %v", err)
	}

	if fetchedSession.Status != "active" {
		t.Errorf("expected status 'active', got '%s'", fetchedSession.Status)
	}

	// 3. Update Table Session (Close Session)
	closedTime := time.Now().Add(30 * time.Minute)
	fetchedSession.Status = "closed"
	fetchedSession.ClosedAt = &closedTime
	fetchedSession.UpdatedAt = time.Now()

	if err := db.Save(&fetchedSession).Error; err != nil {
		t.Fatalf("failed to update session: %v", err)
	}

	var updatedSession TableSession
	if err := db.First(&updatedSession, "id = ?", sessionID).Error; err != nil {
		t.Fatalf("failed to fetch updated session: %v", err)
	}

	if updatedSession.Status != "closed" {
		t.Errorf("expected status 'closed', got '%s'", updatedSession.Status)
	}
}
