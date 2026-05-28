package handlers

import (
	"io"
	"net/http"

	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/services"
	"github.com/gin-gonic/gin"
)

type SSEHandler struct {
	broker *services.Broker
}

func NewSSEHandler(b *services.Broker) *SSEHandler {
	return &SSEHandler{broker: b}
}

func (h *SSEHandler) StreamKitchenEvents(c *gin.Context) {
	establishmentID := c.Query("establishment_id")
	if establishmentID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "establishment_id é obrigatório"})
		return
	}

	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	client := &services.Client{
		EstablishmentID: establishmentID,
		MessageChan:     make(chan string, 10),
	}
	h.broker.Register(client)

	defer h.broker.Unregister(client)

	c.Stream(func(w io.Writer) bool {
		select {
		case payload, ok := <-client.MessageChan:
			if !ok {
				return false
			}
			c.SSEvent("message", payload)
			return true
		case <-c.Request.Context().Done():
			return false
		}
	})
}
