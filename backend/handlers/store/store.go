package store

import (
	"net/http"

	storeModel "github.com/Kleydson-Vieira-1999/resturant-orders-backend/models/store"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StoreHandler struct {
	DB *gorm.DB
}

func NewStoreHandler(db *gorm.DB) *StoreHandler {
	return &StoreHandler{DB: db}
}

func (h *StoreHandler) ListAllStores(c *gin.Context) {
	var stores []storeModel.Store
	userID, _ := c.Get("userID")

	err := h.DB.Where("user_id = ?", userID).Find(&stores).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar lojas: " + err.Error()})
		return
	}
	// store := storeModel.Store{}
	c.JSON(http.StatusOK, gin.H{"stores": stores})
}

func (h *StoreHandler) GetStoreByID(c *gin.Context) {
	userID, _ := c.Get("userID")
	storeID := c.Param("id")

	store := storeModel.Store{}

	h.DB.Where("user_id = ? AND id = ?", userID, storeID).First(&store)

	if store.ID == uuid.Nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Loja não encontrada"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"store": store})
}

func (h *StoreHandler) CreateStore(c *gin.Context) {
	userID, _ := c.Get("userID")

	store := storeModel.Store{}
	err := c.ShouldBindJSON(&store)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Payload Invalido: " + err.Error()})
		return
	}

	uidStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userID inválido no contexto"})
		return
	}

	parsedUUID, parseErr := uuid.Parse(uidStr)
	if parseErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userID inválido: " + parseErr.Error()})
		return
	}

	store.UserID = parsedUUID

	err = h.DB.Create(&store).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar loja: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"store": store})
}

func (h *StoreHandler) UpdateStore(c *gin.Context) {
	userID, _ := c.Get("userID")
	storeID := c.Param("id")

	store := storeModel.Store{}
	h.DB.Where("user_id = ? AND id = ?", userID, storeID).First(&store)

	if store.ID == uuid.Nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Loja não encontrada"})
		return
	}

	err := c.ShouldBindJSON(&store)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Payload Invalido: " + err.Error()})
		return
	}

	err = h.DB.Save(&store).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar loja: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"store": store})
}
