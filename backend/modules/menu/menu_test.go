package menu

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/product"
	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/store"
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
		store_template_id text,
		name text NOT NULL,
		picture text,
		type text,
		description text,
		is_active boolean,
		created_at datetime,
		updated_at datetime
	)`)

	db.Exec(`CREATE TABLE products (
		id text PRIMARY KEY,
		user_id text NOT NULL,
		store_id text,
		name text NOT NULL,
		description text,
		cost_price integer,
		price integer,
		image_base64 text,
		created_at datetime,
		updated_at datetime
	)`)

	db.Exec(`CREATE TABLE menus (
		id text PRIMARY KEY,
		store_id text NOT NULL,
		name text NOT NULL,
		is_active boolean,
		updated_at text
	)`)

	db.Exec(`CREATE TABLE menu_products (
		menu_id text NOT NULL,
		product_id text NOT NULL,
		is_available boolean,
		PRIMARY KEY (menu_id, product_id)
	)`)

	DB = db
	return db
}

func setupRouter(userID uuid.UUID) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	r.Use(func(c *gin.Context) {
		if userID != uuid.Nil {
			c.Set("userID", userID.String())
		}
		c.Next()
	})

	// Use distinct path patterns to avoid wildcard parameter naming conflicts in Gin router
	r.GET("/menus/list/all", ListAllMenus)
	r.GET("/stores/:storeID/menus", ListAllByStoreMenus)
	r.GET("/menus/get/:menuID", GetMenuByID)
	r.POST("/stores/:storeID/menus", CreateMenu)
	r.PATCH("/menus/update/:id", UpdateMenu)
	r.DELETE("/menus/delete/:id", DeleteMenu)
	r.POST("/menus/assoc/:menuID/products/:productID", AddProductToMenu)
	r.PATCH("/menus/assoc/:menuID/products/:productID", UpdateAvailableProductInMenu)
	r.DELETE("/menus/assoc/:menuID/products/:productID", RemoveProductFromMenu)

	return r
}

func TestListAllMenus(t *testing.T) {
	db := setupTestDB(t)
	userID := uuid.New()

	// Create store
	store1 := store.Store{ID: uuid.New(), UserID: userID, Name: "User Store 1"}
	db.Create(&store1)

	// Create menu
	menu1 := Menu{ID: uuid.New(), StoreID: store1.ID, Name: "Menu 1", IsActive: true}
	db.Create(&menu1)

	r := setupRouter(userID)
	req := httptest.NewRequest(http.MethodGet, "/menus/list/all", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestListAllByStoreMenus(t *testing.T) {
	db := setupTestDB(t)
	storeID := uuid.New()

	// Create menu
	m := Menu{ID: uuid.New(), StoreID: storeID, Name: "Store Menu", IsActive: true}
	db.Create(&m)

	r := setupRouter(uuid.New())
	req := httptest.NewRequest(http.MethodGet, "/stores/"+storeID.String()+"/menus", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestGetMenuByID(t *testing.T) {
	db := setupTestDB(t)
	userID := uuid.New()
	menuID := uuid.New()
	storeID := uuid.New()

	// Create store owned by userID
	s := store.Store{ID: storeID, UserID: userID, Name: "User Store"}
	db.Create(&s)

	m := Menu{ID: menuID, StoreID: storeID, Name: "Menu Get", IsActive: true}
	db.Create(&m)

	r := setupRouter(userID)

	req := httptest.NewRequest(http.MethodGet, "/menus/get/"+menuID.String(), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d. Body: %s", w.Code, w.Body.String())
	}
}

func TestCreateMenu(t *testing.T) {
	_ = setupTestDB(t)
	storeID := uuid.New()

	r := setupRouter(uuid.New())

	body := map[string]interface{}{
		"name": "New Menu",
	}
	bodyBytes, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/stores/"+storeID.String()+"/menus", bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected 201, got %d. Body: %s", w.Code, w.Body.String())
	}
}

func TestUpdateMenu(t *testing.T) {
	db := setupTestDB(t)
	userID := uuid.New()
	s := store.Store{ID: uuid.New(), UserID: userID, Name: "Test Store"}
	db.Create(&s)

	m := Menu{ID: uuid.New(), StoreID: s.ID, Name: "Old Name", IsActive: true}
	db.Create(&m)

	r := setupRouter(userID)

	activeVal := false
	body := map[string]interface{}{
		"name":      "New Name",
		"is_active": &activeVal,
	}
	bodyBytes, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPatch, "/menus/update/"+m.ID.String(), bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d. Body: %s", w.Code, w.Body.String())
	}
}

func TestDeleteMenu(t *testing.T) {
	db := setupTestDB(t)
	userID := uuid.New()
	s := store.Store{ID: uuid.New(), UserID: userID, Name: "Test Store"}
	db.Create(&s)

	m := Menu{ID: uuid.New(), StoreID: s.ID, Name: "To Delete", IsActive: true}
	db.Create(&m)

	r := setupRouter(userID)
	req := httptest.NewRequest(http.MethodDelete, "/menus/delete/"+m.ID.String(), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestAddProductToMenu(t *testing.T) {
	db := setupTestDB(t)
	userID := uuid.New()
	s := store.Store{ID: uuid.New(), UserID: userID, Name: "Test Store"}
	db.Create(&s)

	m := Menu{ID: uuid.New(), StoreID: s.ID, Name: "Menu with Product", IsActive: true}
	db.Create(&m)

	p := product.Product{
		ID:        uuid.New(),
		UserID:    userID,
		Name:      "Product 1",
		CostPrice: 100,
		Price:     150,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	db.Create(&p)

	r := setupRouter(userID)

	req := httptest.NewRequest(http.MethodPost, "/menus/assoc/"+m.ID.String()+"/products/"+p.ID.String(), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d. Body: %s", w.Code, w.Body.String())
	}
}

func TestUpdateAvailableProductInMenu(t *testing.T) {
	db := setupTestDB(t)
	userID := uuid.New()
	s := store.Store{ID: uuid.New(), UserID: userID, Name: "Test Store"}
	db.Create(&s)

	m := Menu{ID: uuid.New(), StoreID: s.ID, Name: "Menu with Product", IsActive: true}
	db.Create(&m)

	p := product.Product{
		ID:        uuid.New(),
		UserID:    userID,
		Name:      "Product 1",
		CostPrice: 100,
		Price:     150,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	db.Create(&p)

	menuProduct := MenuProduct{
		MenuID:      m.ID,
		ProductID:   p.ID,
		IsAvailable: true,
	}
	db.Create(&menuProduct)

	r := setupRouter(userID)

	body := map[string]bool{"is_available": false}
	bodyBytes, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPatch, "/menus/assoc/"+m.ID.String()+"/products/"+p.ID.String(), bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d. Body: %s", w.Code, w.Body.String())
	}
}

func TestRemoveProductFromMenu(t *testing.T) {
	db := setupTestDB(t)
	userID := uuid.New()
	s := store.Store{ID: uuid.New(), UserID: userID, Name: "Test Store"}
	db.Create(&s)

	m := Menu{ID: uuid.New(), StoreID: s.ID, Name: "Menu with Product", IsActive: true}
	db.Create(&m)

	p := product.Product{
		ID:        uuid.New(),
		UserID:    userID,
		Name:      "Product 1",
		CostPrice: 100,
		Price:     150,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	db.Create(&p)

	menuProduct := MenuProduct{
		MenuID:      m.ID,
		ProductID:   p.ID,
		IsAvailable: true,
	}
	db.Create(&menuProduct)

	r := setupRouter(userID)

	req := httptest.NewRequest(http.MethodDelete, "/menus/assoc/"+m.ID.String()+"/products/"+p.ID.String(), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d. Body: %s", w.Code, w.Body.String())
	}
}
