package main

import (
	"score-service/repositories"
	"score-service/usecases"

	"github.com/gin-gonic/gin"
)

func main() {
	repositories := repositories.New()

	usecases := usecases.New(usecases.Options{
		Repositories: repositories,
	})

	r := gin.Default()

	r.GET("/person/:document_id/belongings", func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		token = token[7:]

		if token == "undefined" || token == "" {
			c.JSON(401, gin.H{
				"error": "unauthorized",
			})
			return
		}

		DocumentID := c.Param("document_id")
		person, err := usecases.Person.FindByDocumentID(DocumentID, token)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"person": person})
	})

	r.Run(":8081")
}
