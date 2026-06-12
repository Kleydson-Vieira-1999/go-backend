package menu

import (
	"log/slog"
	"net/http"

	menuModel "github.com/Kleydson-Vieira-1999/resturant-orders-backend/models/menu"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MenuHandler struct {
	DB *gorm.DB
}

func NewMenuHandler(db *gorm.DB) *MenuHandler {
	return &MenuHandler{DB: db}
}

func (h *MenuHandler) ListAllMenus(c *gin.Context) {
	userID, _ := c.Get("userID")
	var menus []menuModel.Menu

	// Lista apenas os menus das lojas que pertencem ao usuário
	h.DB.Joins("JOIN stores ON stores.id = menus.store_id").
		Where("stores.user_id = ?", userID).Find(&menus)

	c.JSON(http.StatusOK, gin.H{"menus": menus})
}

func (h *MenuHandler) ListAllByStoreMenus(c *gin.Context) {
	storeID := c.Param("storeID")
	var menus []menuModel.Menu

	err := h.DB.Where("store_id = ?", storeID).Find(&menus).Error
	if err != nil {
		slog.Error("error: ", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"menus": menus})
}

func (h *MenuHandler) GetMenuByID(c *gin.Context) {
	var menu menuModel.Menu
	menuID := c.Param("id")

	// Busca o menu garantindo que ele pertença a uma loja do usuário logado
	h.DB.Where("menus.id = ?", menuID).First(&menu)

	if menu.ID == uuid.Nil {
		slog.Error("Menu não encontrado ", "menuID", menuID)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
		return
	}

	c.JSON(200, gin.H{"menu": menu})
}

func (h *MenuHandler) CreateMenu(c *gin.Context) {
	storeID := c.Param("storeID")
	var menu menuModel.Menu
	err := c.ShouldBindJSON(&menu)
	if err != nil {
		slog.Error("Payload inválido na criação do menu", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
		return
	}

	parsedUUID, parseErr := uuid.Parse(storeID)
	if parseErr != nil {
		slog.Error("storeID inválido no CreateMenu", "error", parseErr)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "storeID inválido: " + parseErr.Error()})
		return
	}

	menu.StoreID = parsedUUID

	err = h.DB.Create(&menu).Error
	if err != nil {
		slog.Error("Erro ao criar menu no banco", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"menu": menu})
}

func (h *MenuHandler) UpdateMenu(c *gin.Context) {
	userID, _ := c.Get("userID")
	menuID := c.Param("id")

	var menu menuModel.Menu
	// Verifica se o menu existe e pertence ao usuário
	if err := h.DB.Joins("JOIN stores ON stores.id = menus.store_id").
		Where("stores.user_id = ? AND menus.id = ?", userID, menuID).
		First(&menu).Error; err != nil {
		slog.Warn("Menu não encontrado ou sem permissão")
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
		return
	}

	var input struct {
		Name     string `json:"name"`
		IsActive *bool  `json:"is_active"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		slog.Error("Dados inválidos", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
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

	if err := h.DB.Model(&menu).Updates(updates).Error; err != nil {
		slog.Warn("Erro ao atualizar menu ")
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"menu": menu})
}

func (h *MenuHandler) DeleteMenu(c *gin.Context) {
	userID, _ := c.Get("userID")
	menuID := c.Param("id")

	result := h.DB.Joins("JOIN stores ON stores.id = menus.store_id").
		Where("stores.user_id = ? AND menus.id = ?", userID, menuID).
		Delete(&menuModel.Menu{})

	if result.RowsAffected == 0 {
		slog.Warn("Tentativa de remover menu inexistente ou sem permissão", "menuID", menuID, "userID", userID)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Menu removido com sucesso"})
}

func (h *MenuHandler) AddProductToMenu(c *gin.Context) {
	userID, _ := c.Get("userID")
	menuID := c.Param("id")
	productID := c.Param("productID")

	// 1. Validar se o menu pertence ao usuário
	var menu menuModel.Menu
	if err := h.DB.Joins("JOIN stores ON stores.id = menus.store_id").
		Where("stores.user_id = ? AND menus.id = ?", userID, menuID).
		First(&menu).Error; err != nil {
		slog.Warn("Menu não encontrado ao adicionar produto", "menuID", menuID, "userID", userID)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
		return
	}

	// 2. Criar a associação na tabela menu_products
	menuProduct := menuModel.MenuProduct{
		MenuID:    uuid.MustParse(menuID),
		ProductID: uuid.MustParse(productID),
	}

	if err := h.DB.Create(&menuProduct).Error; err != nil {
		slog.Error("Erro ao adicionar produto ao cardápio", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Produto adicionado ao cardápio com sucesso"})
}

func (h *MenuHandler) UpdateAvailableProductInMenu(c *gin.Context) {
	userID, _ := c.Get("userID")
	menuID := c.Param("id")
	productID := c.Param("productID")

	// 1. Validar se o menu pertence ao usuário
	var menu menuModel.Menu
	if err := h.DB.Joins("JOIN stores ON stores.id = menus.store_id").
		Where("stores.user_id = ? AND menus.id = ?", userID, menuID).
		First(&menu).Error; err != nil {
		slog.Warn("Menu não encontrado ao atualizar disponibilidade", "menuID", menuID, "userID", userID)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
		return
	}

	var input struct {
		IsAvailable bool `json:"is_available"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		slog.Error("Payload inválido no UpdateAvailableProductInMenu", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
		return
	}

	// 2. Atualizar a disponibilidade na tabela menu_products
	result := h.DB.Model(&menuModel.MenuProduct{}).
		Where("menu_id = ? AND product_id = ?", menuID, productID).
		Update("is_available", input.IsAvailable)

	if result.Error != nil {
		slog.Error("Erro ao atualizar disponibilidade do produto no cardápio", "error", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
		return
	}

	if result.RowsAffected == 0 {
		slog.Warn("Associação menu-produto não encontrada", "menuID", menuID, "productID", productID)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Disponibilidade atualizada com sucesso"})
}

func (h *MenuHandler) RemoveProductFromMenu(c *gin.Context) {
	userID, _ := c.Get("userID")
	menuID := c.Param("id")
	productID := c.Param("productID")

	// Validar se o menu pertence ao usuário antes de remover o link
	var menu menuModel.Menu
	if err := h.DB.Joins("JOIN stores ON stores.id = menus.store_id").
		Where("stores.user_id = ? AND menus.id = ?", userID, menuID).
		First(&menu).Error; err != nil {
		slog.Warn("Menu não encontrado ao remover produto", "menuID", menuID, "userID", userID)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
		return
	}

	result := h.DB.Where("menu_id = ? AND product_id = ?", menuID, productID).
		Delete(&menuModel.MenuProduct{})

	if result.RowsAffected == 0 {
		slog.Warn("Associação menu-produto não encontrada", "menuID", menuID, "productID", productID)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Produto removido do cardápio com sucesso"})
}
