package core

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_SSL"))

	var db *gorm.DB
	var err error
	for i := 1; i <= 5; i++ {
		slog.Info(fmt.Sprintf("Tentando conectar ao banco... (Tentativa %d/5)", i))
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			slog.Info("Conectado ao banco de dados com sucesso!")
			return db // Fixed bug: return db instead of nil
		}

		slog.Error(fmt.Sprintf("Erro na tentativa %d: %v. Aguardando...", i, err))
		time.Sleep(2 * time.Second)
	}

	slog.Warn(fmt.Sprintf("Falha crônica no banco após 5 tentativas: %v", err))
	return db
}
