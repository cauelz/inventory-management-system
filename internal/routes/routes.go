package routes

import (
	"inventory-management-system/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)


func RegisterRoutes(router *gin.Engine, db *sqlx.DB) {
    productHandler := &handlers.ProductHandler{DB: db}

    router.GET("/products", productHandler.GetProducts)
    router.GET("/products/:id", productHandler.GetProductByID)
    router.POST("/products", productHandler.CreateProduct)
    router.PUT("/products/:id", productHandler.UpdateProduct)
    router.DELETE("/products/:id", productHandler.DeleteProduct)
}