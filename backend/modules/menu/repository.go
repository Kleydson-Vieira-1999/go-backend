package menu



type MenuRepository struct {
}

func (h *MenuRepository) GetMenuByID(menuID string) (Menu, error) {
	var menu Menu
	err := DB.Where("id = ?", menuID).First(&menu).Error
	return menu, err
}

func GetMenuProductsByMenuID(menuID string) ([]MenuProduct, error) {
	var menuProducts []MenuProduct
	err := DB.Where("menu_id = ?", menuID).Find(&menuProducts).Error
	return menuProducts, err
}

