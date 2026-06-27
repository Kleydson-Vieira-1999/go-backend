package user

import (
	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/core"
)

func SetupRoutes() {
	protected := Router.Group("/action/api")
	protected.Use(core.NewAuthMiddleware().AuthUser())
	{
		protected.GET("/google/auth", SingInWithJWT)
	}
}
