package main

import (
	"fmt"
	"inventory-management-system/handlers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func main() {

    // Ler variáveis de ambiente
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("POSTGRES_USER")
    dbPassword := os.Getenv("POSTGRES_PASSWORD")
    dbName := os.Getenv("POSTGRES_DB")

	fmt.Println(dbHost, dbPort, dbUser, dbPassword, dbName)

    // Construir a string de conexão
    connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        dbHost, dbPort, dbUser, dbPassword, dbName)

	var err error

	// Connect to the database
	db, err = sqlx.Connect("postgres", connStr)

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
