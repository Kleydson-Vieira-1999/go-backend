package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/features"
	"github.com/Kleydson-Vieira-1999/resturant-orders-backend/models/user"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gorm.io/gorm"
)

type GoogleUser struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
}

type GoogleAuthHandler struct {
	DB     *gorm.DB
	config *oauth2.Config
}

func NewGoogleAuthHandler(db *gorm.DB) *GoogleAuthHandler {
	data, err := os.ReadFile("client_secret_apps.googleusercontent.com.json")
	if err != nil {
		slog.Warn(fmt.Sprintf("Error Configuração do Google não encontrada no servidor %v", err))
		log.Fatal("Ocorreu um erro")
	}

	configData, err := google.ConfigFromJSON(data,
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
	)
	if err != nil {
		slog.Warn(fmt.Sprintf("Erro ao processar credenciais do Google %v", err))
		log.Fatal("Ocorreu um erro")
	}

	return &GoogleAuthHandler{DB: db, config: configData}
}

func (h *GoogleAuthHandler) GetAuthURL(c *gin.Context) {
	url := h.config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	c.JSON(http.StatusOK, gin.H{"url": url})
}

func (h *GoogleAuthHandler) PostAuthToken(c *gin.Context) {
	var input struct {
		State string `json:"state"`
		Code  string `json:"code"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		slog.Warn(fmt.Sprintf("Paylod JSON inválido %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
		return
	}

	ctx := context.Background()
	token, err := h.config.Exchange(ctx, input.Code)
	if err != nil {
		slog.Error("Erro ao trocar código por token Google", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
		return
	}

	client := h.config.Client(ctx, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		slog.Warn(fmt.Sprintf("Erro ao buscar userinfo: %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
		return
	}
	defer resp.Body.Close()

	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Warn(fmt.Sprintf("Erro ao ler resposta do Google: %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
		return
	}

	var userEstableshment user.User
	var googleUser GoogleUser
	if err := json.Unmarshal(responseBytes, &googleUser); err != nil {
		slog.Warn(fmt.Sprintf("Erro ao processar dados do perfil: %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
		return
	}

	h.DB.Where("sso_id = ?", googleUser.ID).First(&userEstableshment)

	if userEstableshment.ID == uuid.Nil {

		userEstableshment = user.User{
			SSOProvider: "google",
			SSOID:       googleUser.ID,
			Email:       googleUser.Email,
			Name:        googleUser.Name,
			Picture:     googleUser.Picture,
		}

		if err := h.DB.Create(&userEstableshment).Error; err != nil {
			slog.Warn(fmt.Sprintf("Erro ao cadastrar dados do usuario: %v", err))
			c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
			return
		}
	}

	tokenBuilder := features.NewTokenBuilder()
	tokenString, err := tokenBuilder.CreateUserToken(
		userEstableshment.ID.String(),
		userEstableshment.Email)
	if err != nil {
		slog.Warn(fmt.Sprintf("Erro ao criar token %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro"})
		return
	}

	returnedUser := user.User{
		Name:    userEstableshment.Name,
		Email:   userEstableshment.Email,
		Picture: userEstableshment.Picture}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"data":  returnedUser,
	})
}
