package store

import (
	"log/slog"
	"net/http"

	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/core"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UpdateStore(c *gin.Context) {
	userID, _ := c.Get("userID")
	storeID := c.Param("id")

	store := Store{}
	DB.Where("user_id = ? AND id = ?", userID, storeID).First(&store)

	if store.ID == uuid.Nil {
		slog.Warn("Tentativa de atualizar loja inexistente", "storeID", storeID, "userID", userID)
		core.NewResponse().Error(c)
		return
	}

	err := c.ShouldBindJSON(&store)
	if err != nil {
		slog.Error("Payload inválido na atualização da loja", "error", err)
		core.NewResponse().Error(c)
		return
	}

	err = DB.Save(&store).Error
	if err != nil {
		slog.Error("Erro ao atualizar loja", "error", err)
		core.NewResponse().Error(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"store": store})
}
