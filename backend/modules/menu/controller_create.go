package menu

import (
	"log/slog"
	"net/http"

	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/core"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateMenu(c *gin.Context) {
	storeID := c.Param("storeID")
	var menu Menu
	err := c.ShouldBindJSON(&menu)
	if err != nil {
		slog.Error("Payload inválido na criação do menu", "error", err)
		core.NewResponse().Error(c)
		return
	}

	parsedUUID, parseErr := uuid.Parse(storeID)
	if parseErr != nil {
		slog.Error("storeID inválido no CreateMenu", "error", parseErr)
		core.NewResponse().Error(c)
		return
	}

	menu.StoreID = parsedUUID

	err = DB.Create(&menu).Error
	if err != nil {
		slog.Error("Erro ao criar menu no banco", "error", err)
		core.NewResponse().Error(c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"menu": menu})
}

func AddProductToMenu(c *gin.Context) {
	userID := c.GetString("userID")
	menuID := c.Param("menuID")
	productID := c.Param("productID")

	// 1. Validar se o menu pertence ao usuário
	_, hasError := VerifyMenuId(userID, menuID)
	if hasError {
		core.NewResponse().Error(c)
		return
	}

	// 2. Criar a associação na tabela menu_products
	menuProduct := MenuProduct{
		MenuID:    uuid.MustParse(menuID),
		ProductID: uuid.MustParse(productID),
	}

	if err := DB.Create(&menuProduct).Error; err != nil {
		slog.Error("Erro ao adicionar produto ao cardápio", "error", err)
		core.NewResponse().Error(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Produto adicionado ao cardápio com sucesso"})
}
