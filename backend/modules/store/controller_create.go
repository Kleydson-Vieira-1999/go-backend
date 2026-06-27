package store

import (
	"log/slog"
	"net/http"

	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/core"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateStore(c *gin.Context) {
	userID, _ := c.Get("userID")

	store := Store{}
	err := c.ShouldBindJSON(&store)
	if err != nil {
		slog.Error("Payload inválido na criação da loja", "error", err)
		core.NewResponse().Error(c)
		return
	}

	uidStr, ok := userID.(string)
	if !ok {
		slog.Error("userID inválido no contexto")
		core.NewResponse().Error(c)
		return
	}

	parsedUUID, parseErr := uuid.Parse(uidStr)
	if parseErr != nil {
		slog.Error("Erro ao parsear UUID do usuário", "error", parseErr)
		core.NewResponse().Error(c)
		return
	}

	store.UserID = parsedUUID

	err = DB.Create(&store).Error
	if err != nil {
		slog.Error("Erro ao criar loja no banco", "error", err)
		core.NewResponse().Error(c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"store": store})
}
