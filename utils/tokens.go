package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Used to sign the JWT
var signingKey = []byte("Gokedex")

// Generate a JWT from the user id and username, signed with a 30 day exp date
func GenerateJWT(id int64, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"id":       id,
		"exp":      time.Now().UTC().AddDate(0, 1, 0).Unix(),
		"nbf":      time.Date(2024, 06, 01, 12, 0, 0, 0, time.UTC).Unix(),
	})

	tokenString, err := token.SignedString(signingKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Validates a JWT token and returns valid, user ID and error
func ValidateJWT(tokenString string) (bool, int64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return signingKey, nil
	})

	if err != nil {
		return false, 0, err
	}

	if !token.Valid {
		return false, 0, nil
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		expireAtClaim, _ := claims.GetExpirationTime()

		if expireAtClaim.After(time.Now().UTC()) {

			return true, int64(claims["id"].(float64)), nil
		}
	}

	return false, 0, nil
}
