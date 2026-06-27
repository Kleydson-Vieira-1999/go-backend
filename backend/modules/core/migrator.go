package core

import (
	"fmt"
	"log/slog"

	"gorm.io/gorm"
)

type Migrator struct {
}

type MigrationFunc func(db *gorm.DB) error

var registry []MigrationFunc

func (h *Migrator) Register(fn MigrationFunc) {
	registry = append(registry, fn)
}

func (h *Migrator) RunAll(db *gorm.DB) error {
	slog.Info(fmt.Sprintf("Executando %d migrações registradas...", len(registry)))
	for _, fn := range registry {
		if err := fn(db); err != nil {
			return err
		}
	}
	slog.Info("Migrações concluídas com sucesso.")
	return nil
}
