package product

import (
	"log/slog"
	"net/http"

	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/core"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateProduct(c *gin.Context) {
	userID, _ := c.Get("userID")
	var product Product

	if err := c.ShouldBindJSON(&product); err != nil {
		slog.Error("Erro ao decodificar JSON do produto", "error", err)
		core.NewResponse().Error(c)
		return
	}

	userIDStr, ok := userID.(string)
	if !ok {
		slog.Error("userID ausente ou inválido no contexto")
		core.NewResponse().Error(c)
		return
	}

	parsedUserID, err := uuid.Parse(userIDStr)
	if err != nil {
		slog.Error("Erro ao converter userID para UUID", "error", err)
		core.NewResponse().Error(c)
		return
	}

	product.UserID = parsedUserID

	if err := DB.Create(&product).Error; err != nil {
		slog.Error("Erro ao persistir produto no banco", "error", err)
		core.NewResponse().Error(c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"product": product})
}
