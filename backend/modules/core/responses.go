package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
}

func NewResponse() *Response {
	return &Response{}
}

func (r *Response) Error(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro"})
}
