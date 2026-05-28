package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/features"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			log.Println("error: Cabeçalho Authorization é obrigatório")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Ocorreu um erro"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			log.Println("error: Formato do cabeçalho Authorization inválido")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Ocorreu um erro"})
			c.Abort()
			return
		}

		tokenString := parts[1]
		tokenBuilder := features.NewTokenBuilder()
		token, err := tokenBuilder.ValidateUserToken(tokenString)
		if err != nil {
			if err == jwt.ErrTokenExpired {
				log.Println("error: Token expirado")
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Ocorreu um erro"})
				c.Abort()
				return
			}
		}

		c.Set("userID", token.UserID)
		c.Set("email", token.UserEmail)
		c.Next()
	}
}
