package core

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Routing struct {
	Router *gin.Engine
	StandardResponse *Response
}

type Database struct {
	DB *gorm.DB
	Migrator *Migrator
}

var routing *Routing
var database *Database

func InitModule() (*Routing, *Database) {
	loggerJson := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(loggerJson)

	if err := godotenv.Load(); err != nil {
		slog.Info("Aviso: .env não encontrado, usando variáveis de sistema")
	}

	db := Connect()

	routing = &Routing{Router: Setup(loggerJson), StandardResponse: NewResponse()};
	database = &Database{DB: db, Migrator: &Migrator{}}

	return routing, database
}

func StartApplication() {
	err := database.Migrator.RunAll(database.DB)
	if err != nil {
		slog.Warn(fmt.Sprintf("Falha ao rodar migrações: %v", err))
	}

	sqlDB, _ := database.DB.DB()
	defer sqlDB.Close()

	slog.Info("Servidor iniciado em http://localhost:8080")
	routing.Router.Run(":8080")
}