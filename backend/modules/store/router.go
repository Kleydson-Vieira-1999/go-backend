package store

import (
	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/core"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	protected := r.Group("/action/api")
	protected.Use(core.NewAuthMiddleware().AuthUser())
	{
		protected.GET("/stores", ListAllStores)
		protected.GET("/stores/:id", GetStoreByID)
		protected.POST("/stores", CreateStore)
		protected.PATCH("/stores/:id", UpdateStore)
	}
}
