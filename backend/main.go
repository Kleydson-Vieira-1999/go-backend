package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/database"
	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/models"
	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/models/codes"
	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/models/menu"
	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/models/product"
	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/models/store"
	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/models/table"
	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/models/user"
	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/routes"

	"github.com/joho/godotenv"
)

func main() {
	loggerJson := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(loggerJson)

	if err := godotenv.Load(); err != nil {
		slog.Info("Aviso: .env não encontrado, usando variáveis de sistema")
	}

	db, err := database.Connect()
	if err != nil {
		slog.Warn(fmt.Sprintf("Falha no banco:", err))
	}

	slog.Info("Executando migrações...")
	err = db.AutoMigrate(
		&user.User{},
		&store.Store{},
		&store.StoreTemplate{},
		&product.Product{},
		&product.ProductTemplate{},
		&menu.MenuTemplate{},
		&menu.MenuProductTemplate{},
		&menu.Menu{},
		&menu.MenuProduct{},
		&codes.KitchenCode{},
		&codes.WaiterCode{},
		&table.Table{},
		&table.TableSession{},
		&models.Order{},
		&models.OrderItem{},
		&store.StoreBalance{},
	)
	if err != nil {
		slog.Warn("Falha na migração:", err)
	}

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	router := routes.Setup(db)

	slog.Info("Servidor iniciado em http://localhost:8080")
	router.Run(":8080")
}
