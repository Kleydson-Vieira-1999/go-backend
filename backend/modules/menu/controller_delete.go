package menu

import (
	"log/slog"
	"net/http"

	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/core"
	"github.com/gin-gonic/gin"
)

func DeleteMenu(c *gin.Context) {
	userID, _ := c.Get("userID")
	menuID := c.Param("id")

	result := DB.Where("id = ? AND store_id IN (SELECT id FROM stores WHERE user_id = ?)", menuID, userID).
		Delete(&Menu{})

	if result.RowsAffected == 0 {
		slog.Warn("Tentativa de remover menu inexistente ou sem permissão", "menuID", menuID, "userID", userID)
		core.NewResponse().Error(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Menu removido com sucesso"})
}

func RemoveProductFromMenu(c *gin.Context) {
	userID := c.GetString("userID")
	menuID := c.Param("menuID")
	productID := c.Param("productID")

	// Validar se o menu pertence ao usuário antes de remover o link
	_, hasError := VerifyMenuId(userID, menuID)
	if hasError {
		core.NewResponse().Error(c)
		return
	}

	result := DB.Where("menu_id = ? AND product_id = ?", menuID, productID).
		Delete(&MenuProduct{})

	if result.RowsAffected == 0 {
		slog.Warn("Associação menu-produto não encontrada", "menuID", menuID, "productID", productID)
		core.NewResponse().Error(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Produto removido do cardápio com sucesso"})
}
