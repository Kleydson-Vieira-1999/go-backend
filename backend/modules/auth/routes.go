package auth

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	googleAuth := NewGoogleAuthHandler()
	auths := r.Group("/api/google/auth")
	auths.GET("/url", googleAuth.GetAuthURL)
	auths.POST("/token", googleAuth.PostAuthToken)
}
