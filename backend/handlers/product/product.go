package product

import (
	"log/slog"
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

func (h *ProductHandler) ListProductsByMenu(c *gin.Context) {
	menuID := c.Param("menuID")
	var products []productModel.ProductWithAvailability

	err := h.DB.Table("products").
		Joins("JOIN menu_products ON menu_products.product_id = products.id").
		Where("menu_products.menu_id = ?", menuID).
		Select("products.*, menu_products.is_available"). // Seleciona os campos desejados
		Find(&products).Error
	if err != nil {
		slog.Error("Erro ao buscar produtos por menu", "error", err, "menuID", menuID)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}

func (h *ProductHandler) ListAllProductsByUser(c *gin.Context) {
	userID, _ := c.Get("userID")
	var products []productModel.Product

	if err := h.DB.Where("user_id = ?", userID).Find(&products).Error; err != nil {
		slog.Error("Erro ao buscar produtos por usuário", "error", err, "userID", userID)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	productID := c.Param("id")
	var product productModel.Product

	if err := h.DB.First(&product, "id = ?", productID).Error; err != nil {
		slog.Error("Produto não encontrado", "error", err, "id", productID)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	userID, _ := c.Get("userID")
	var product productModel.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		slog.Error("Erro ao decodificar JSON do produto", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
		return
	}

	userIDStr, ok := userID.(string)
	if !ok {
		slog.Error("userID ausente ou inválido no contexto")
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
		return
	}

	parsedUserID, err := uuid.Parse(userIDStr)
	if err != nil {
		slog.Error("Erro ao converter userID para UUID", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
		return
	}

	product.UserID = parsedUserID

	if err := h.DB.Create(&product).Error; err != nil {
		slog.Error("Erro ao persistir produto no banco", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
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
		slog.Error("Tentativa de atualizar produto inexistente ou sem permissão", "userID", userID, "productID", productID)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		slog.Error("Dados inválidos na atualização do produto", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
		return
	}

	if err := h.DB.Save(&product).Error; err != nil {
		slog.Error("Erro ao salvar produto", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
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
		slog.Error("Tentativa de remover produto inexistente ou sem permissão", "userID", userID, "productID", productID)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Produto removido com sucesso"})
}
