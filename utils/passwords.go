package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// Hash a plain string password
func HashPassword(password string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return "", err
	}

	return string(result), nil
}

// Check if received password is the same as the stored password from the DB
func ComparePassword(dbPassword string, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))
	if err != nil {
		return false, err
	}

	return true, nil
}
