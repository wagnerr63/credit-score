package auth

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

type ValidateListPersonBelongingsRequestDTO struct {
	UserID int    `json:"user_id"`
	Token  string `json:"token"`
	Role   string `json:"role"`
}

func (u *authUseCases) ValidateListPersonBelongingsRequest(data ValidateListPersonBelongingsRequestDTO) error {
	token, err := jwt.Parse(data.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return errors.New("Unexpected signing method"), nil
		}
		return hmacSampleSecret, nil
	})

	if err != nil {
		return errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return errors.New("invalid token")
	}

	userById, _ := u.repositories.User.FindByID(claims["user_id"].(int))
	if userById.ID == 0 {
		return errors.New("invalid user")
	}

	if userById.Role == "master" || userById.Role == "credit_institution" || userById.Role == "credit_analyst" {
		return nil
	}

	return errors.New("unauthorized")

}
