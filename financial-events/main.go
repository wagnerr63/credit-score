package main

import (
	"financial-events/repositories"

	"github.com/gin-gonic/gin"
)

func main() {
	repositories := repositories.New()

	r := gin.Default()

	r.GET("/documentIDConsult/:documentID/last", func(c *gin.Context) {
		documentID := c.Param("documentID")
		lastDocumentIDConsult, err := repositories.DocumentIDConsultRepository.FindLastConsultByDocumentID(documentID)

		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, lastDocumentIDConsult)

	})

	r.GET("/financialEvent/:documentID", func(c *gin.Context) {
		documentID := c.Param("documentID")
		financialEvent, err := repositories.FinancialEventRepository.ListByDocumentID(documentID)

		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, financialEvent)

	})

	r.GET("/creditCardTransaction/:documentID/last", func(c *gin.Context) {
		documentID := c.Param("documentID")
		creditCardTransaction, err := repositories.CreditCardTransactionRepository.FindLastByDocumentID(documentID)

		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, creditCardTransaction)

	})

	r.Run(":8082")
}
