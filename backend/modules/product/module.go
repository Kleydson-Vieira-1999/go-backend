package product

import (
	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/core"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitModule(routing *core.Routing, database *core.Database) {
	DB = database.DB
	database.Migrator.Register(RunMigrations)
	SetupRoutes(routing.Router)
}
