package product

import (
	"log/slog"
	"net/http"

	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/core"
	"github.com/gin-gonic/gin"
)

func DeleteProduct(c *gin.Context) {
	userID, _ := c.Get("userID")
	productID := c.Param("id")

	result := DB.Where("id = ? AND store_id IN (SELECT id FROM stores WHERE user_id = ?)", productID, userID).
		Delete(&Product{})

	if result.RowsAffected == 0 {
		slog.Error("Tentativa de remover produto inexistente ou sem permissão", "userID", userID, "productID", productID)
		core.NewResponse().Error(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Produto removido com sucesso"})
}
