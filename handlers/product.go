package handlers

import (
	"inventory-management-system/db"
	"inventory-management-system/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	DB db.Database
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	var products []models.Product

	error := h.DB.Select(&products, "SELECT * FROM products")

	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch products", "error": error.Error()})
	}

	c.JSON(http.StatusOK, products)

}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	var product models.Product

	// Get the ID from the URL
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID is required"})
		return
	}

	error := h.DB.Get(&product, "SELECT * FROM products WHERE id = $1", id)

	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch product", "error": error.Error()})
		return 
	}

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product models.Product

	// Bind the request body to the product struct
	error := c.BindJSON(&product)

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body", "error": error.Error()})
		return
	}

	query := "INSERT INTO products (name, description, price, stock_quantity) VALUES ($1, $2, $3, $4) RETURNING *"

	row := h.DB.QueryRow(query, product.Name, product.Description, product.Price, product.StockQty)

	error = row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.StockQty, &product.CreatedAt, &product.UpdatedAt)

	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create product", "error": error.Error()})
		return 
	}

	c.JSON(http.StatusCreated, product)

}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	var product models.Product

	// Bind the request body to the product struct
	error := c.BindJSON(&product)

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body", "error": error.Error()})
		return
	}

	query := "UPDATE products SET name = $1, description = $2, price = $3, stock_qty = $4 WHERE id = $5 RETURNING *"

	row := h.DB.QueryRow(query, product.Name, product.Description, product.Price, product.StockQty, product.ID)

	error = row.Scan(&product.ID)

	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update product", "error": error.Error()})
		return 
	}

	c.JSON(http.StatusOK, product)

}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	// Get the ID from the URL
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID is required"})
		return
	}

	query := "DELETE FROM products WHERE id = $1"

	_, error := h.DB.Exec(query, id)

	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete product", "error": error.Error()})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}