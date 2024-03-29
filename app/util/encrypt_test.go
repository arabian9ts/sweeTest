package util

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptPassword(t *testing.T) {
	t.Run("len(password)=0, success", func(t *testing.T) {
		password := ""
		digest := EncryptPassword(password)

		assert.Equal(t, 60, len(digest))
	})

	t.Run("len(password)=8, success", func(t *testing.T) {
		password := "abchefgh"
		digest := EncryptPassword(password)

		assert.Equal(t, 60, len(digest))
	})

	t.Run("len(password)=100, success", func(t *testing.T) {
		password := strings.Repeat("a", 100)
		digest := EncryptPassword(password)

		assert.Equal(t, 60, len(digest))
	})
}
