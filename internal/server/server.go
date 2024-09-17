package server

import (
	"inventory-management-system/internal/routes"
	"inventory-management-system/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)


type Server struct {
	config *config.Config
	router *gin.Engine
	db	 *sqlx.DB
}

func NewServer(config *config.Config, dbConnection *sqlx.DB) *Server {
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