package main

import (
	"debts-service/repositories"
	"debts-service/usecases"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	repositories := repositories.New()

	usecases := usecases.New(usecases.Options{
		Repositories: repositories,
	})

	r := gin.Default()
	r.GET("/personDebts/:personID", func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		// remove Bearer from token
		token = token[7:]

		if token == "undefined" {
			c.JSON(401, gin.H{
				"error": "unauthorized",
			})
			return
		}

		personID, _ := strconv.Atoi(c.Params.ByName("personID"))

		personDebts, err := usecases.Person.GetPersonDebtsByPersonID(personID, token)
		if err != nil {
			c.JSON(400, err.Error())
			return
		}

		c.JSON(200, gin.H{
			"personDebts": personDebts,
		})
	})
	r.Run(":3333")
}
