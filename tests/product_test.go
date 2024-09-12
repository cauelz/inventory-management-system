package tests

import (
	"errors"
	"inventory-management-system/handlers"
	"inventory-management-system/mocks"
	"inventory-management-system/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllProducts(t *testing.T) {
	
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	mockDB := new(mocks.MockDB)

	mockProducts := []models.Product{
		{ID: 1, Name: "Product 1", Description: "Description 1", Price: 100, StockQty: 10},
		{ID: 2, Name: "Product 2", Description: "Description 2", Price: 200, StockQty: 20},
	}

	mockDB.On("Select", mock.Anything, "SELECT * FROM products", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		dest := args.Get(0).(*[]models.Product)
		*dest = mockProducts
	})

	productHandler := handlers.ProductHandler{DB: mockDB}

	r.GET("/products", productHandler.GetProducts)

	// Simula a requisição HTTP GET
	req, _ := http.NewRequest(http.MethodGet, "/products", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Verifica o código de status e o corpo da resposta
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Product 1")
	assert.Contains(t, w.Body.String(), "Product 2")

	// Verifica se as expectativas do mock foram atendidas
	mockDB.AssertExpectations(t)

}

func TestErrorWhenGetAllProducts(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	mockDB := new(mocks.MockDB)

	mockDB.On("Select", mock.Anything, "SELECT * FROM products", mock.Anything).Return(errors.New("produtos não encontrados"))

	productHandler := handlers.ProductHandler{DB: mockDB}

	r.GET("/products", productHandler.GetProducts)

	req, _ := http.NewRequest(http.MethodGet, "/products", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "Could not fetch products")

	// Verifica se as expectativas do mock foram atendidas
	mockDB.AssertExpectations(t)
}
