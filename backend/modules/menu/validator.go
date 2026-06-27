package menu

import (
	"log/slog"


	"github.com/google/uuid"
)

func VerifyMenuId(userID, menuID string) (*Menu, bool) {
	var menu Menu
	// Verifica se o menu existe e pertence ao usuário
	if err := DB.Joins("JOIN stores ON stores.id = menus.store_id").
		Where("stores.user_id = ? AND menus.id = ?", userID, menuID).
		First(&menu).Error; err != nil {
		slog.Warn("Menu não encontrado ou tentativa de ação sem permissão")
		return &Menu{}, true
	}

	if menu.ID == uuid.Nil {
		slog.Error("Menu não encontrado ", "menuID", menuID)
		return &Menu{}, true
	}
	return &menu, false
}
