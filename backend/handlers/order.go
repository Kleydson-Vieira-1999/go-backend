package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/models"
	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderHandler struct {
	broker *services.Broker
	db     *gorm.DB
}

func NewOrderHandler(b *services.Broker, db *gorm.DB) *OrderHandler {
	return &OrderHandler{broker: b, db: db}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		slog.Error("Erro ao decodificar JSON do pedido", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Ocorreu um erro"})
		return
	}

	if order.SessionID == uuid.Nil {
		slog.Warn("Tentativa de criar pedido sem session_id")
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Ocorreu um erro"})
		return
	}

	if order.ID == uuid.Nil {
		order.ID = uuid.New()
	}
	order.Status = "pending"

	if err := h.db.Create(&order).Error; err != nil {
		slog.Error("Erro ao salvar pedido no banco", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
		return
	}

	orderJSON, _ := json.Marshal(order)
	h.broker.Broadcast(order.SessionID.String(), string(orderJSON))

	c.JSON(http.StatusCreated, gin.H{"message": "Pedido criado e enviado para a cozinha"})
}
