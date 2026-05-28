package handlers

import (
	"encoding/json"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao processar o JSON da requisição"})
		return
	}

	if order.SessionID == uuid.Nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "session_id é obrigatório"})
		return
	}

	if order.ID == uuid.Nil {
		order.ID = uuid.New()
	}
	order.Status = "pending"

	if err := h.db.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar pedido"})
		return
	}

	orderJSON, _ := json.Marshal(order)
	h.broker.Broadcast(order.SessionID.String(), string(orderJSON))

	c.JSON(http.StatusCreated, gin.H{"message": "Pedido criado e enviado para a cozinha"})
}
