package main

import (
	"fmt"
	"inventory-management-system/internal/server"
	"inventory-management-system/pkg/config"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var db *sqlx.DB

func main() {

	var err error
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalln(err)
		return
	}

    // Ler variáveis de ambiente
    dbHost := viper.GetString("DB_HOST")
    dbPort := viper.GetString("DB_PORT")
    dbUser := viper.GetString("DB_USER")
    dbPassword := viper.GetString("DB_PASSWORD")
    dbName := viper.GetString("DB_NAME")

	fmt.Println(dbHost, dbPort, dbUser, dbPassword, dbName)

    // Construir a string de conexão
    connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	// Connect to the database
	db, err = sqlx.Connect("postgres", connStr)

	if err != nil {
		log.Fatalln(err)
	}

	// Initialize Gin
	r := server.NewServer(cfg, db)

	err = r.Run()

	if err != nil {
		log.Fatalln(err)
	}
}
