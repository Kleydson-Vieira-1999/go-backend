package user

import (
	"log/slog"
	"net/http"

	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/core"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SingInWithJWT(c *gin.Context) {
	userID, _ := c.Get("userID")
	var user User

	DB.First(&user, "id = ?", userID)

	if user.ID == uuid.Nil {
		slog.Warn("Usuário não encontrado")
		core.NewResponse().Error(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
