package main

import (
	"geekbakes/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// public route
	r.POST("/login", handlers.Login)

	// auth route
	protected := r.Group("/")
	protected.Use(handlers.AuthMiddleware())
	{
		protected.POST("/invoices", handlers.CreateInvoice)
		protected.GET("/invoices", handlers.GetInvoices)
	}

	r.Run(":8080") // Start the server on port 8080
}
