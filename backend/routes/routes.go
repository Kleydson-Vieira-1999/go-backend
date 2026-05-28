package routes

import (
	"os"
	"time"

	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/handlers"
	handlersAuth "github.com/Kleydson-Vieira-1999/resturant-orders-backend/handlers/auth"
	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/handlers/menu"
	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/handlers/product"
	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/handlers/store"
	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/middlewares"
	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB) *gin.Engine {
	broker := services.NewBroker()
	go broker.Start()

	if os.Getenv("ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/health", func(c *gin.Context) { c.String(200, "API está online!") })

	googleAuthHandler := handlersAuth.NewGoogleAuthHandler(db)
	r.GET("/api/google/auth/url", googleAuthHandler.GetAuthURL)
	r.POST("/api/google/auth/token", googleAuthHandler.PostAuthToken)

	userHandler := handlers.NewUserHandler(db)
	userAuthMiddleware := middlewares.NewAuthMiddleware()

	protected := r.Group("/action/api")
	protected.Use(userAuthMiddleware.AuthUser())
	{
		protected.GET("/google/auth", userHandler.SingInWithJWT)

		storeHandler := store.NewStoreHandler(db)
		protected.GET("/stores", storeHandler.ListAllStores)
		protected.GET("/stores/:id", storeHandler.GetStoreByID)
		protected.POST("/stores", storeHandler.CreateStore)
		protected.PATCH("/stores/:id", storeHandler.UpdateStore)

		menuHandler := menu.NewMenuHandler(db)
		protected.GET("/menus", menuHandler.ListAllMenus)
		protected.GET("/menus/:id", menuHandler.GetMenuByID)
		protected.GET("/menus/s/:storeID", menuHandler.ListAllByStoreMenus)
		protected.POST("/menus/:storeID", menuHandler.CreateMenu)
		protected.POST("/menus/p/:id/*productID", menuHandler.AddProductToMenu)
		protected.DELETE("/menus/p/:id/*productID", menuHandler.RemoveProductFromMenu)

		productHandler := product.NewProductHandler(db)
		protected.GET("/products/s/:storeID", productHandler.ListProductsByStore)
		protected.GET("/products/:id", productHandler.GetProductByID)
		protected.POST("/products/s/:storeID", productHandler.CreateProduct)
		protected.PATCH("/products/:id", productHandler.UpdateProduct)
		protected.DELETE("/products/:id", productHandler.DeleteProduct)
	}

	r.POST("/establesh/singin", userHandler.SingIn)

	sseHandler := handlers.NewSSEHandler(broker)
	r.GET("/api/stream/kitchen", sseHandler.StreamKitchenEvents)

	orderHandler := handlers.NewOrderHandler(broker, db)
	r.POST("/api/orders", orderHandler.CreateOrder)

	return r
}
