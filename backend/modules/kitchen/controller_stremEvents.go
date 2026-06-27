package kitchen

import (
	"io"
	"log/slog"

	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/core"
	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/store"
	"github.com/gin-gonic/gin"
)

func StreamKitchenEvents(c *gin.Context) {
	establishmentID := c.Query("establishment_id")
	if establishmentID == "" {
		slog.Warn("Tentativa de abrir SSE sem establishment_id")
		core.NewResponse().Error(c)
		return
	}

	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	client := &store.Client{
		EstablishmentID: establishmentID,
		MessageChan:     make(chan string, 10),
	}
	store.Br.Register(client)

	defer store.Br.Unregister(client)

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
