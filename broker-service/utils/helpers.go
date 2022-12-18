package utils

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	AppName = "qLib"
)

func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

func ValidatePassword(userPassword string, hashed []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashed, []byte(userPassword))
	return err == nil
}

// GenerateUUID returns a new v4 UUID.
func GenerateUUID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	return id.String(), nil
}
