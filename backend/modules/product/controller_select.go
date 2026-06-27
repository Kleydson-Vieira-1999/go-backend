package product

import (
	"log/slog"
	"net/http"

	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/core"
	"github.com/gin-gonic/gin"
)

func ListProductsByMenu(c *gin.Context) {
	menuID := c.Param("menuID")
	var products []ProductWithAvailability

	err := DB.Table("products").
		Joins("JOIN menu_products ON menu_products.product_id = products.id").
		Where("menu_products.menu_id = ?", menuID).
		Select("products.*, menu_products.is_available"). // Seleciona os campos desejados
		Find(&products).Error
	if err != nil {
		slog.Error("Erro ao buscar produtos por menu", "error", err, "menuID", menuID)
		core.NewResponse().Error(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}

func ListAllProductsByUser(c *gin.Context) {
	userID, _ := c.Get("userID")
	var products []Product

	if err := DB.Where("user_id = ?", userID).Find(&products).Error; err != nil {
		slog.Error("Erro ao buscar produtos por usuário", "error", err, "userID", userID)
		core.NewResponse().Error(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}

func GetProductByID(c *gin.Context) {
	productID := c.Param("id")
	var product Product

	if err := DB.First(&product, "id = ?", productID).Error; err != nil {
		slog.Error("Produto não encontrado", "error", err, "id", productID)
		core.NewResponse().Error(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}
