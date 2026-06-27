package order

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/core"
	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/store"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateOrder(c *gin.Context) {
	var order Order
	if err := c.ShouldBindJSON(&order); err != nil {
		slog.Error("Erro ao decodificar JSON do pedido", "error", err)
		core.NewResponse().Error(c)
		return
	}

	if order.SessionID == uuid.Nil {
		slog.Warn("Tentativa de criar pedido sem session_id")
		core.NewResponse().Error(c)
		return
	}

	if order.ID == uuid.Nil {
		order.ID = uuid.New()
	}
	order.Status = "pending"

	if err := DB.Create(&order).Error; err != nil {
		slog.Error("Erro ao salvar pedido no banco", "error", err)
		core.NewResponse().Error(c)
		return
	}

	orderJSON, _ := json.Marshal(order)
	store.Br.Broadcast(order.SessionID.String(), string(orderJSON))

	c.JSON(http.StatusCreated, gin.H{"message": "Pedido criado e enviado para a cozinha"})
}
