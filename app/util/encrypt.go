package util

import (
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}

	return string(hash)
}
