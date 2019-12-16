package handler

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
)

type AliveHandler struct{}

func NewAliveHandler() (*AliveHandler, error) {
	return &AliveHandler{}, nil
}

func (handler *AliveHandler) Echo(response string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, response)
	}
}
