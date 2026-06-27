package menu

import (
	"log/slog"
	"net/http"

	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/core"
	"github.com/gin-gonic/gin"
)

func UpdateMenu(c *gin.Context) {
	userID := c.GetString("userID")

	menuID := c.Param("id")

	menu, hasError := VerifyMenuId(userID, menuID)
	if hasError {
		core.NewResponse().Error(c)
		return
	}

	var input struct {
		Name     string `json:"name"`
		IsActive *bool  `json:"is_active"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		slog.Error("Dados inválidos", "error", err)
		core.NewResponse().Error(c)
		return
	}

	// Atualiza apenas os campos enviados
	updates := make(map[string]interface{})
	if input.Name != "" {
		updates["name"] = input.Name
	}
	if input.IsActive != nil {
		updates["is_active"] = *input.IsActive
	}

	if err := DB.Model(&menu).Updates(updates).Error; err != nil {
		slog.Warn("Erro ao atualizar menu ")
		core.NewResponse().Error(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"menu": menu})
}

func UpdateAvailableProductInMenu(c *gin.Context) {
	userID := c.GetString("userID")
	menuID := c.Param("menuID")
	productID := c.Param("productID")

	// 1. Validar se o menu pertence ao usuário
	_, hasError := VerifyMenuId(userID, menuID)
	if hasError {
		core.NewResponse().Error(c)
		return
	}

	var input struct {
		IsAvailable bool `json:"is_available"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		slog.Error("Payload inválido no UpdateAvailableProductInMenu", "error", err)
		core.NewResponse().Error(c)
		return
	}

	// 2. Atualizar a disponibilidade na tabela menu_products
	result := DB.Model(&MenuProduct{}).
		Where("menu_id = ? AND product_id = ?", menuID, productID).
		Update("is_available", input.IsAvailable)

	if result.Error != nil {
		slog.Error("Erro ao atualizar disponibilidade do produto no cardápio", "error", result.Error)
		core.NewResponse().Error(c)
		return
	}

	if result.RowsAffected == 0 {
		slog.Warn("Associação menu-produto não encontrada", "menuID", menuID, "productID", productID)
		core.NewResponse().Error(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Disponibilidade atualizada com sucesso"})
}
