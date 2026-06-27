package user

import (
	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/core"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var Router *gin.Engine
var DB *gorm.DB

func InitModule(routing *core.Routing, database *core.Database) {
	DB = database.DB
	Router = routing.Router
	database.Migrator.Register(RunMigrations)
	SetupRoutes()
}
