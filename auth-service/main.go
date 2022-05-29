package main

import (
	"auth-service/entities"
	"auth-service/repositories"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var hmacSampleSecret []byte

type ValidateRequestDTO struct {
	Action string `json:"action"`
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
}

func main() {
	repositories := repositories.New()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/auth", func(c *gin.Context) {
		var user entities.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		userByEmail, err := repositories.User.FindByEmail(user.Email)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if userByEmail.Password != user.Password {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid credentials"})
			return
		}

		// Token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id":   userByEmail.ID,
			"user_role": userByEmail.Role,
			"nbf":       time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		})

		tokenString, err := token.SignedString(hmacSampleSecret)

		c.JSON(http.StatusOK, gin.H{
			"token": tokenString,
		})
	})

	r.POST("/auth/validateRequest", func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")

		tokenString = tokenString[7:] // remove Bearer from token string

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return hmacSampleSecret, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		var validateRequestDTO ValidateRequestDTO
		if err := c.ShouldBindJSON(&validateRequestDTO); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		if validateRequestDTO.Action == "listPersonDebts" && validateRequestDTO.Role != "master" && validateRequestDTO.UserID != int(claims["user_id"].(float64)) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{})

	})

	r.Run()
}
