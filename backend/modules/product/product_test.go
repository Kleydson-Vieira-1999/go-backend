package product_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/menu"
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
		is_active boolean
	)`)

	db.Exec(`CREATE TABLE menu_products (
		menu_id text NOT NULL,
		product_id text NOT NULL,
		is_available boolean,
		PRIMARY KEY (menu_id, product_id)
	)`)

	product.DB = db
	return db
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	// Helper middleware to mock userID injection
	r.Use(func(c *gin.Context) {
		userID := c.GetHeader("X-User-ID")
		if userID != "" {
			c.Set("userID", userID)
		}
		c.Next()
	})

	r.GET("/menus/:menuID/products", product.ListProductsByMenu)
	r.GET("/products", product.ListAllProductsByUser)
	r.GET("/products/:id", product.GetProductByID)
	r.POST("/products", product.CreateProduct)
	r.PATCH("/products/:id", product.UpdateProduct)
	r.DELETE("/products/:id", product.DeleteProduct)

	return r
}

func TestListProductsByMenu(t *testing.T) {
	db := setupTestDB(t)
	r := setupRouter()

	menuID := uuid.New()
	productID1 := uuid.New()
	productID2 := uuid.New()

	p1 := product.Product{
		ID:        productID1,
		UserID:    uuid.New(),
		Name:      "Product 1",
		Price:     1000,
		CostPrice: 500,
	}
	p2 := product.Product{
		ID:        productID2,
		UserID:    uuid.New(),
		Name:      "Product 2",
		Price:     1500,
		CostPrice: 800,
	}

	db.Create(&p1)
	db.Create(&p2)

	menuProd1 := menu.MenuProduct{
		MenuID:      menuID,
		ProductID:   productID1,
		IsAvailable: true,
	}
	menuProd2 := menu.MenuProduct{
		MenuID:      menuID,
		ProductID:   productID2,
		IsAvailable: false,
	}

	db.Create(&menuProd1)
	db.Create(&menuProd2)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/menus/"+menuID.String()+"/products", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d. Body: %s", w.Code, w.Body.String())
	}

	var resp map[string][]product.ProductWithAvailability
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	products := resp["products"]
	if len(products) != 2 {
		t.Fatalf("expected 2 products, got %d", len(products))
	}
}

func TestListAllProductsByUser(t *testing.T) {
	db := setupTestDB(t)
	r := setupRouter()

	userID := uuid.New()
	p1 := product.Product{
		ID:        uuid.New(),
		UserID:    userID,
		Name:      "User Product 1",
		Price:     1000,
		CostPrice: 500,
	}
	p2 := product.Product{
		ID:        uuid.New(),
		UserID:    uuid.New(), // different user
		Name:      "Other Product",
		Price:     1500,
		CostPrice: 800,
	}

	db.Create(&p1)
	db.Create(&p2)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products", nil)
	req.Header.Set("X-User-ID", userID.String())
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var resp map[string][]product.Product
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	products := resp["products"]
	if len(products) != 1 {
		t.Fatalf("expected 1 product, got %d", len(products))
	}
}

func TestGetProductByID(t *testing.T) {
	db := setupTestDB(t)
	r := setupRouter()

	productID := uuid.New()
	p := product.Product{
		ID:        productID,
		UserID:    uuid.New(),
		Name:      "Special Product",
		Price:     2000,
		CostPrice: 1000,
	}
	db.Create(&p)

	t.Run("Found", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/products/"+productID.String(), nil)
		r.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", w.Code)
		}
	})
}

func TestCreateProduct(t *testing.T) {
	_ = setupTestDB(t)
	r := setupRouter()

	userID := uuid.New()
	newProduct := product.Product{
		Name:      "New Created Product",
		Price:     1200,
		CostPrice: 600,
	}

	body, _ := json.Marshal(newProduct)

	t.Run("Success", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-User-ID", userID.String())
		r.ServeHTTP(w, req)

		if w.Code != http.StatusCreated {
			t.Fatalf("expected status 201, got %d. Body: %s", w.Code, w.Body.String())
		}
	})
}

func TestUpdateProduct(t *testing.T) {
	db := setupTestDB(t)
	r := setupRouter()

	userID := uuid.New()
	storeID := uuid.New()
	productID := uuid.New()

	// Create store owned by user
	s := store.Store{
		ID:     storeID,
		UserID: userID,
		Name:   "User's Store",
	}
	db.Create(&s)

	// Create product linked to store
	p := product.Product{
		ID:        productID,
		UserID:    userID,
		Name:      "Original Product",
		Price:     1000,
		CostPrice: 500,
	}
	db.Create(&p)
	db.Exec("UPDATE products SET store_id = ? WHERE id = ?", storeID, productID)

	t.Run("Success", func(t *testing.T) {
		updatedData := product.Product{
			ID:        productID,
			UserID:    userID,
			Name:      "Updated Product Name",
			Price:     1200,
			CostPrice: 600,
		}
		body, _ := json.Marshal(updatedData)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PATCH", "/products/"+productID.String(), bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-User-ID", userID.String())
		r.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("expected status 200, got %d. Body: %s", w.Code, w.Body.String())
		}
	})
}

func TestDeleteProduct(t *testing.T) {
	db := setupTestDB(t)
	r := setupRouter()

	userID := uuid.New()
	storeID := uuid.New()
	productID := uuid.New()

	// Create store owned by user
	s := store.Store{
		ID:     storeID,
		UserID: userID,
		Name:   "User's Store",
	}
	db.Create(&s)

	// Create product linked to store
	p := product.Product{
		ID:        productID,
		UserID:    userID,
		Name:      "Product to delete",
		Price:     1000,
		CostPrice: 500,
	}
	db.Create(&p)
	db.Exec("UPDATE products SET store_id = ? WHERE id = ?", storeID, productID)

	t.Run("Success", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/products/"+productID.String(), nil)
		req.Header.Set("X-User-ID", userID.String())
		r.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d. Body: %s", w.Code, w.Body.String())
		}
	})
}
