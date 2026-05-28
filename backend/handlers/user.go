package handlers

import (
	"net/http"

	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/models/user"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserHandler struct {
	DB *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{DB: db}
}

func (h *UserHandler) SingIn(c *gin.Context) {
	var input struct {
		Name        string `json:"name"`
		Email       string `json:"email"`
		SSOProvider string `json:"sso_provider"`
		SSOID       string `json:"sso_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Payload JSON inválido"})
		return
	}

	user := user.User{
		Name:        input.Name,
		Email:       input.Email,
		SSOProvider: input.SSOProvider,
		SSOID:       input.SSOID,
	}

	if err := h.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "sso_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"updated_at"}),
	}).Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar banco de dados: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":      user.ID,
		"message": "Autenticado com sucesso",
	})
}

func (h *UserHandler) SingInWithJWT(c *gin.Context) {
	userID, _ := c.Get("userID")
	var user user.User

	h.DB.First(&user, "id = ?", userID)

	if user.ID == uuid.Nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
