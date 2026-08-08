package menu

import (
	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/core"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	protected := r.Group("/action/api")
	protected.Use(core.NewAuthMiddleware().AuthUser())
	{
		protected.GET("/menus", ListAllMenusByUser)
		protected.GET("/menus/:menuID", GetMenuByID)
		protected.GET("/menus/s/:storeID", ListAllByStoreMenus)
		protected.POST("/menus/:storeID", CreateMenu)
		protected.PATCH("/menus/:storeID", UpdateMenu)
		protected.POST("/menus/p/:menuID/:productID", AddProductToMenu)
		protected.PATCH("/menus/p/:menuID/:productID", UpdateAvailableProductInMenu)
		protected.DELETE("/menus/p/:menuID/:productID", RemoveProductFromMenu)
	}

	unprotected := r.Group("/public/api")
	{
		unprotected.GET("/menu/:menuID", GetMenuByID)
	}
}
