package menu

import (
	"log/slog"
	"net/http"

	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/core"
	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/modules/product"
	"github.com/gin-gonic/gin"
)

func ListAllMenusByUser(c *gin.Context) {
	userID, _ := c.Get("userID")
	var menus []Menu

	// Lista apenas os menus das lojas que pertencem ao usuário
	DB.Joins("JOIN stores ON stores.id = menus.store_id").
		Where("stores.user_id = ?", userID).Find(&menus)

	c.JSON(http.StatusOK, gin.H{"menus": menus})
}

func ListAllByStoreMenus(c *gin.Context) {
	// userID := c.GetString("userID")
	storeID := c.Param("storeID")

	// h.verifyStoreId(userID, storeID)

	var menus []Menu

	err := DB.Where("store_id = ?", storeID).Find(&menus).Error
	if err != nil {
		slog.Error("error: ", "error", err)
		core.NewResponse().Error(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"menus": menus})
}

func GetMenuByID(c *gin.Context) {
	userID := c.GetString("userID")

	menuID := c.Param("menuID")
	menu, hasError := VerifyMenuId(userID, menuID)
	if hasError {
		core.NewResponse().Error(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"menu": menu})
}

func GetMenuByIDWithoutUserID(c *gin.Context) {
	menuID := c.Param("id")
	var menu Menu

	err := DB.Where("id = ?", menuID).First(&menu).Error
	if err != nil {
		slog.Error("", "error", err)
		core.NewResponse().Error(c)
		return
	}

	var products []product.ProductWithAvailability
	err = DB.Table("products").
		Joins("JOIN menu_products ON menu_products.product_id = products.id").
		Where("menu_products.menu_id = ?", menuID).
		Select("products.*, menu_products.is_available"). // Seleciona os campos desejados
		Find(&products).Error

	if err != nil {
		slog.Error("Erro ao buscar produtos por menu", "error", err, "menuID", menuID)
		core.NewResponse().Error(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"menu":     menu,
		"products": products,
	})
}
