package product

import (
	"net/http"

	productModel "github.com/Kleydson-Vieira-1999/resturant-orders-backend/models/product"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductHandler struct {
	DB *gorm.DB
}

func NewProductHandler(db *gorm.DB) *ProductHandler {
	return &ProductHandler{DB: db}
}

func (h *ProductHandler) ListProductsByStore(c *gin.Context) {
	storeID := c.Param("storeID")
	var products []productModel.Product

	if err := h.DB.Where("store_id = ?", storeID).Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar produtos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	productID := c.Param("id")
	var product productModel.Product

	if err := h.DB.First(&product, "id = ?", productID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	userID, _ := c.Get("userID")
	var product productModel.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Payload inválido: " + err.Error()})
		return
	}

	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userID inválido"})
		return
	}

	parsedUserID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userID inválido"})
		return
	}

	product.UserID = parsedUserID

	if err := h.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar produto"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"product": product})
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	userID, _ := c.Get("userID")
	productID := c.Param("id")
	var product productModel.Product

	// Segurança: Garante que o produto pertence a uma loja do usuário
	if err := h.DB.Joins("JOIN stores ON stores.id = products.store_id").
		Where("stores.user_id = ? AND products.id = ?", userID, productID).
		First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado ou sem permissão"})
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	if err := h.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar produto"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Produto atualizado", "product": product})
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	userID, _ := c.Get("userID")
	productID := c.Param("id")

	result := h.DB.Joins("JOIN stores ON stores.id = products.store_id").
		Where("stores.user_id = ? AND products.id = ?", userID, productID).
		Delete(&productModel.Product{})

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado ou sem permissão"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Produto removido com sucesso"})
}
