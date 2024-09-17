package server

import (
	"inventory-management-system/db"
	"inventory-management-system/internal/routes"
	"inventory-management-system/pkg/config"

	"github.com/gin-gonic/gin"
)


type Server struct {
	config *config.Config
	router *gin.Engine
	db	 db.Database
}

func NewServer(config *config.Config, dbConnection db.Database) *Server {
	return &Server{
		config: config,
		db: dbConnection,
		router: gin.Default(),
	}
}

func (server *Server) Run() error {
	routes.RegisterRoutes(server.router, server.db)
	return server.router.Run(server.config.ServerAddress)
}