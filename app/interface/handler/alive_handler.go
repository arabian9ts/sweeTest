package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
