package handler

import (
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
	"github.com/arabian9ts/sweeTest/app/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Repository repository.UserRepository
}

func NewAuthHandler(repository repository.UserRepository) (*AuthHandler, error) {
	return &AuthHandler{Repository: repository}, nil
}

func (handler *AuthHandler) StudentAuthHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := util.Decode(ctx.Request)

		if err != nil {
			ctx.AbortWithStatus(401)
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		if claims["role"] != "student" {
			ctx.AbortWithStatus(401)
			return
		}

		id64 := int64(claims["id"].(float64))
		student, err := handler.Repository.GetStudentById(id64)

		if err != nil {
			ctx.AbortWithStatus(401)
			return
		}

		ctx.Set("currentUser", student)
	}
}

func (handler *AuthHandler) AssistantAuthHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := util.Decode(ctx.Request)

		if err != nil {
			ctx.AbortWithStatus(401)
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		if claims["role"] != "assistant" {
			ctx.AbortWithStatus(401)
			return
		}

		id64 := int64(claims["id"].(float64))
		assistant, err := handler.Repository.GetAssistantById(id64)

		if err != nil {
			ctx.AbortWithStatus(401)
			return
		}

		ctx.Set("currentUser", assistant)
	}
}

func (handler *AuthHandler) TeacherAuthHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := util.Decode(ctx.Request)

		if err != nil {
			ctx.AbortWithStatus(401)
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		if claims["role"] != "teacher" {
			ctx.AbortWithStatus(401)
			return
		}

		id64 := int64(claims["id"].(float64))
		teacher, err := handler.Repository.GetTeacherById(id64)

		if err != nil {
			ctx.AbortWithStatus(401)
			return
		}

		ctx.Set("currentUser", teacher)
	}
}

func (handler *AuthHandler) AdminAuthHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := util.Decode(ctx.Request)

		if err != nil {
			ctx.AbortWithStatus(401)
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		if claims["role"] != "admin" {
			ctx.AbortWithStatus(401)
			return
		}

		id64 := int64(claims["id"].(float64))
		admin, err := handler.Repository.GetAdminById(id64)

		if err != nil {
			ctx.AbortWithStatus(401)
			return
		}

		ctx.Set("currentUser", admin)
	}
}
