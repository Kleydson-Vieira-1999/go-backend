package store

import (
	"log/slog"
	"net/http"

	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/core"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ListAllStores(c *gin.Context) {
	var stores []Store
	userID, _ := c.Get("userID")

	err := DB.Where("user_id = ?", userID).Find(&stores).Error
	if err != nil {
		slog.Error("Erro ao buscar lojas do usuário", "error", err, "userID", userID)
		core.NewResponse().Error(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{"stores": stores})
}

func GetStoreByID(c *gin.Context) {
	userID, _ := c.Get("userID")
	storeID := c.Param("id")

	store := Store{}

	DB.Where("user_id = ? AND id = ?", userID, storeID).First(&store)

	if store.ID == uuid.Nil {
		slog.Warn("Loja não encontrada", "storeID", storeID, "userID", userID)
		core.NewResponse().Error(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"store": store})
}
