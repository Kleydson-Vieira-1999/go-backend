package menu

import (
	"log"
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

	c.JSON(200, gin.H{"menus": menus})
}

func (h *MenuHandler) ListAllByStoreMenus(c *gin.Context) {
	storeID := c.Param("storeID")
	var menus []menuModel.Menu

	err := h.DB.Where("store_id = ?", storeID).Find(&menus).Error
	if err != nil {
		log.Println("error: " + err.Error())
		c.JSON(200, gin.H{"menus": menus})
		return
	}

	c.JSON(200, gin.H{"menus": menus})
}

func (h *MenuHandler) GetMenuByID(c *gin.Context) {
	userID, _ := c.Get("userID")
	var menu menuModel.Menu
	menuID := c.Param("id")

	// Busca o menu garantindo que ele pertença a uma loja do usuário logado
	h.DB.Joins("JOIN stores ON stores.id = menus.store_id").Where("stores.user_id = ? AND menus.id = ?", userID, menuID).First(&menu)

	if menu.ID == uuid.Nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Menu não encontrado"})
		return
	}

	c.JSON(200, gin.H{"menu": menu})
}

func (h *MenuHandler) CreateMenu(c *gin.Context) {
	storeID := c.Param("storeID")
	var menu menuModel.Menu
	err := c.ShouldBindJSON(&menu)
	if err != nil {
		log.Println("error: Paylod Invalido " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Payload Invalido: " + err.Error()})
		return
	}

	parsedUUID, parseErr := uuid.Parse(storeID)
	if parseErr != nil {
		log.Println("error: storeID inválido " + parseErr.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "storeID inválido: " + parseErr.Error()})
		return
	}

	menu.StoreID = parsedUUID

	err = h.DB.Create(&menu).Error
	if err != nil {
		log.Println("error: Erro ao criar menu " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar menu: " + err.Error()})
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
		c.JSON(http.StatusNotFound, gin.H{"error": "Menu não encontrado ou sem permissão"})
		return
	}

	var input struct {
		Name     string `json:"name"`
		IsActive *bool  `json:"is_active"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar menu"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Menu atualizado com sucesso", "menu": menu})
}

func (h *MenuHandler) DeleteMenu(c *gin.Context) {
	userID, _ := c.Get("userID")
	menuID := c.Param("id")

	result := h.DB.Joins("JOIN stores ON stores.id = menus.store_id").
		Where("stores.user_id = ? AND menus.id = ?", userID, menuID).
		Delete(&menuModel.Menu{})

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Menu não encontrado ou sem permissão"})
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
		c.JSON(http.StatusNotFound, gin.H{"error": "Menu não encontrado ou sem permissão"})
		return
	}

	// 2. Criar a associação na tabela menu_products
	menuProduct := menuModel.MenuProduct{
		MenuID:    uuid.MustParse(menuID),
		ProductID: uuid.MustParse(productID),
	}

	if err := h.DB.Create(&menuProduct).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao adicionar produto ao cardápio (talvez já esteja adicionado)"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Produto adicionado ao cardápio com sucesso"})
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
		c.JSON(http.StatusNotFound, gin.H{"error": "Menu não encontrado ou sem permissão"})
		return
	}

	result := h.DB.Where("menu_id = ? AND product_id = ?", menuID, productID).
		Delete(&menuModel.MenuProduct{})

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Associação não encontrada"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Produto removido do cardápio com sucesso"})
}
