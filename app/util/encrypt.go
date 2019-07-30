package util

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func EncryptPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}

	return string(hash)
}

func EncodeStudent(student *model.Student) string {
	jwtConfig := config.GetJwtConfig()

	expires := jwtConfig.ExpiresTime()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = student.ID
	claims["role"] = "student"
	claims["student_no"] = student.StudentNo
	claims["exp"] = expires.Unix()

	tokenString, err := token.SignedString([]byte(jwtConfig.Secret))
	if err != nil {
		return ""
	}

	return tokenString
}

func EncodeAssistant(assistant *model.Assistant) string {
	jwtConfig := config.GetJwtConfig()

	expires := jwtConfig.ExpiresTime()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = assistant.ID
	claims["role"] = "assistant"
	claims["student_no"] = assistant.StudentNo
	claims["exp"] = expires.Unix()

	tokenString, err := token.SignedString([]byte(jwtConfig.Secret))
	if err != nil {
		return ""
	}

	return tokenString
}

func EncodeTeacher(teacher *model.Teacher) string {
	jwtConfig := config.GetJwtConfig()

	expires := jwtConfig.ExpiresTime()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = teacher.ID
	claims["role"] = "teacher"
	claims["exp"] = expires.Unix()

	tokenString, err := token.SignedString([]byte(jwtConfig.Secret))
	if err != nil {
		return ""
	}

	return tokenString
}

func EncodeAdmin(admin *model.Admin) string {
	jwtConfig := config.GetJwtConfig()

	expires := jwtConfig.ExpiresTime()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = admin.ID
	claims["role"] = "admin"
	claims["exp"] = expires.Unix()

	tokenString, err := token.SignedString([]byte(jwtConfig.Secret))
	if err != nil {
		return ""
	}

	return tokenString
}

func Decode(req *http.Request) (*jwt.Token, error) {
	token, err := request.ParseFromRequest(req, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
		b := []byte(config.GetJwtConfig().Secret)

		return b, nil
	})

	return token, err
}
