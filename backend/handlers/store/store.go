package store

import (
	"log/slog"
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
		slog.Error("Erro ao buscar lojas do usuário", "error", err, "userID", userID)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"stores": stores})
}

func (h *StoreHandler) GetStoreByID(c *gin.Context) {
	userID, _ := c.Get("userID")
	storeID := c.Param("id")

	store := storeModel.Store{}

	h.DB.Where("user_id = ? AND id = ?", userID, storeID).First(&store)

	if store.ID == uuid.Nil {
		slog.Warn("Loja não encontrada", "storeID", storeID, "userID", userID)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"store": store})
}

func (h *StoreHandler) CreateStore(c *gin.Context) {
	userID, _ := c.Get("userID")

	store := storeModel.Store{}
	err := c.ShouldBindJSON(&store)
	if err != nil {
		slog.Error("Payload inválido na criação da loja", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
		return
	}

	uidStr, ok := userID.(string)
	if !ok {
		slog.Error("userID inválido no contexto")
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
		return
	}

	parsedUUID, parseErr := uuid.Parse(uidStr)
	if parseErr != nil {
		slog.Error("Erro ao parsear UUID do usuário", "error", parseErr)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
		return
	}

	store.UserID = parsedUUID

	err = h.DB.Create(&store).Error
	if err != nil {
		slog.Error("Erro ao criar loja no banco", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
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
		slog.Warn("Tentativa de atualizar loja inexistente", "storeID", storeID, "userID", userID)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
		return
	}

	err := c.ShouldBindJSON(&store)
	if err != nil {
		slog.Error("Payload inválido na atualização da loja", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
		return
	}

	err = h.DB.Save(&store).Error
	if err != nil {
		slog.Error("Erro ao atualizar loja", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"store": store})
}
