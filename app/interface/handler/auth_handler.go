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
			ctx.JSON(401, "Invalid Token")
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		if claims["role"] != "student" {
			ctx.JSON(401, "Invalid Token")
			return
		}

		id64 := int64(claims["id"].(float64))
		student, err := handler.Repository.GetStudentById(id64)

		if err != nil {
			ctx.JSON(401, "Invalid Token")
			return
		}

		ctx.Set("currentUser", student)
	}
}

func (handler *AuthHandler) AssistantAuthHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := util.Decode(ctx.Request)

		if err != nil {
			ctx.JSON(401, "Invalid Token")
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		if claims["role"] != "assistant" {
			ctx.JSON(401, "Invalid Token")
			return
		}

		id64 := int64(claims["id"].(float64))
		assistant, err := handler.Repository.GetAssistantById(id64)

		if err != nil {
			ctx.JSON(401, "Invalid Token")
			return
		}

		ctx.Set("currentUser", assistant)
	}
}

func (handler *AuthHandler) TeacherAuthHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := util.Decode(ctx.Request)

		if err != nil {
			ctx.JSON(401, "Invalid Token")
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		if claims["role"] != "teacher" {
			ctx.JSON(401, "Invalid Token")
			return
		}

		id64 := int64(claims["id"].(float64))
		teacher, err := handler.Repository.GetTeacherById(id64)

		if err != nil {
			ctx.JSON(401, "Invalid Token")
			return
		}

		ctx.Set("currentUser", teacher)
	}
}

func (handler *AuthHandler) AdminAuthHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := util.Decode(ctx.Request)

		if err != nil {
			ctx.JSON(401, "Invalid Token")
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		if claims["role"] != "admin" {
			ctx.JSON(401, "Invalid Token")
			return
		}

		id64 := int64(claims["id"].(float64))
		admin, err := handler.Repository.GetAdminById(id64)

		if err != nil {
			ctx.JSON(401, "Invalid Token")
			return
		}

		ctx.Set("currentUser", admin)
	}
}
