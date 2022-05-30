package auth

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

type ValidateListPersonDebtsRequestDTO struct {
	UserID int    `json:"user_id"`
	Token  string `json:"token"`
	Role   string `json:"role"`
}

var hmacSampleSecret []byte

func ValidateListPersonDebtsRequest(data ValidateListPersonDebtsRequestDTO) error {
	token, err := jwt.Parse(data.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return errors.New("Unexpected signing method"), nil
		}
		return hmacSampleSecret, nil
	})

	if err != nil {
		return errors.New("invalid token 1")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return errors.New("invalid token 2")
	}

	if data.Role != "master" && data.UserID != int(claims["user_id"].(float64)) {
		return errors.New("unauthorized")
	}

	return nil
}
