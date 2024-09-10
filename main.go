package main

import (
	"log"
	"net/http"

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

	r.GET("/", func(context *gin.Context){
		context.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})

	r.Run(":8080")
}