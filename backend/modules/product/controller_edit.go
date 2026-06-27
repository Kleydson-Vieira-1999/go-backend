package product

import (
	"log/slog"
	"net/http"

	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/core"
	"github.com/gin-gonic/gin"
)

func UpdateProduct(c *gin.Context) {
	userID, _ := c.Get("userID")
	productID := c.Param("id")
	var product Product

	// Segurança: Garante que o produto pertence a uma loja do usuário
	if err := DB.Joins("JOIN stores ON stores.id = products.store_id").
		Where("stores.user_id = ? AND products.id = ?", userID, productID).
		First(&product).Error; err != nil {
		slog.Error("Tentativa de atualizar produto inexistente ou sem permissão", "userID", userID, "productID", productID)
		core.NewResponse().Error(c)
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		slog.Error("Dados inválidos na atualização do produto", "error", err)
		core.NewResponse().Error(c)
		return
	}

	if err := DB.Save(&product).Error; err != nil {
		slog.Error("Erro ao salvar produto", "error", err)
		core.NewResponse().Error(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Produto atualizado", "product": product})
}
