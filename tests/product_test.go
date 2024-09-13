package tests

import (
	"errors"
	"inventory-management-system/handlers"
	"inventory-management-system/mocks"
	"inventory-management-system/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
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

func TestGetProductById(t *testing.T) {

	gin.SetMode(gin.TestMode)
	r := gin.Default()

	mockDB := new(mocks.MockDB)

	mockProduct := models.Product{ID: 1, Name: "Product 1", Description: "Description 1", Price: 100, StockQty: 10}

	mockDB.On("Get", mock.Anything, "SELECT * FROM products WHERE id = $1", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		dest := args.Get(0).(*models.Product)
		*dest = mockProduct
	})

	productHandler := handlers.ProductHandler{DB: mockDB}

	r.GET("/products/:id", productHandler.GetProductByID)

	req, _ := http.NewRequest(http.MethodGet, "/products/1", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Product 1")

	mockDB.AssertExpectations(t)
}

func TestUpdateProduct(t *testing.T) {

    gin.SetMode(gin.TestMode)

    r := gin.Default()

    // Crie um mock de banco de dados usando sqlmock
	db, mock, err := sqlmock.New()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

	// Layout correspondente à string de data/hora
	layout := time.RFC3339

	// Converter strings para time.Time
	createdAt, err := time.Parse(layout, "2021-09-01T00:00:00Z")
	if err != nil {
		t.Fatalf("an error '%s' was not expected when parsing created_at", err)
	}
	updatedAt, err := time.Parse(layout, "2021-09-01T00:00:00Z")
	if err != nil {
		t.Fatalf("an error '%s' was not expected when parsing updated_at", err)
	}

    // Configure o mock para a consulta SQL
    mock.ExpectQuery("UPDATE products SET name = \\$1, description = \\$2, price = \\$3, stock_quantity = \\$4 WHERE id = \\$5 RETURNING \\*").
        WithArgs("Product Updated", "Updated Description", 150.0, 5, 1).
        WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "price", "stock_quantity", "created_at", "updated_at"}).
		AddRow(1, "Product Updated", "Updated Description", 150.0, 5, createdAt, updatedAt))

    productHandler := handlers.ProductHandler{DB: sqlxDB}
    r.PUT("/products/:id", productHandler.UpdateProduct)

    productJSON := `{"id":1, "name":"Product Updated","description":"Updated Description","price":150.0,"stock_quantity":5}`
    req, _ := http.NewRequest(http.MethodPut, "/products/1", strings.NewReader(productJSON))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)

    // Verifique se todas as expectativas foram atendidas
    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("there were unfulfilled expectations: %s", err)
    }
}