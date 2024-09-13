package main

import (
	"inventory-management-system/handlers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func main() {
	var err error

	// Connect to the database
	db, err = sqlx.Connect("postgres", "user=admin password=password dbname=inventory sslmode=disable")

	if err != nil {
		log.Fatalln(err)
	}

	// Initialize Gin
	r := gin.Default()

	// Create a new ProductHandler
	productHandler := &handlers.ProductHandler{DB: db}

	// Products routes
	r.GET("/products", productHandler.GetProducts)
	r.GET("/products/:id", productHandler.GetProductByID)
	r.POST("/products", productHandler.CreateProduct)
	r.PUT("/products/:id", productHandler.UpdateProduct)
	r.DELETE("/products/:id", productHandler.DeleteProduct)

	r.Run(":8080")
}
