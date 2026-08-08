package product

import (
	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/core"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	protected := r.Group("/action/api")
	protected.Use(core.NewAuthMiddleware().AuthUser())
	{
		protected.GET("/products", ListAllProductsByUser)
		protected.GET("/products/:id", GetProductByID)
		protected.GET("/products/bymenu/:menuID", ListProductsByMenu)
	}
}