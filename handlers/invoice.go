package handlers

import (
	"geekbakes/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// create invoice
func CreateInvoice(c *gin.Context) {
	var newInvoice models.Invoice

	// Bind JSON to invoice
	if err := c.ShouldBindJSON(&newInvoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// append invoice to model
	models.Invoices = append(models.Invoices, newInvoice)

	// generate invoice
	c.JSON(http.StatusOK, gin.H{"invoice": newInvoice})
}

// retrieve invoice
func GetInvoices(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"invoices": models.Invoices})
}
