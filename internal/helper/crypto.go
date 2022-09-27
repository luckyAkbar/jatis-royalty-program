package helper

import (
	"errors"
	"strings"

	"github.com/luckyAkbar/jatis-royalty-program/internal/config"

	"github.com/kumparan/go-utils"
	"github.com/mattheath/base62"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// HashString encrypt given text
func HashString(text string) (string, error) {
	bt, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bt), nil
}

func IsHashedStringMatch(plain, cipher []byte) bool {
	err := bcrypt.CompareHashAndPassword(cipher, plain)
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return false
	}
	if err != nil {
		logrus.Error(err)
		return false
	}
	return true
}

func GenerateToken(uniqueID int64) string {
	sb := strings.Builder{}

	encodedID := base62.EncodeInt64(uniqueID)
	sb.WriteString(encodedID)
	sb.WriteString("_")

	randString := utils.GenerateRandomAlphanumeric(config.DefaultTokenLength)
	sb.WriteString(randString)

	return sb.String()
}