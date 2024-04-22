package tests

import (
	"log"

	"github.com/adefarhan/warmindo-be/internal/delivery/http"
	"github.com/adefarhan/warmindo-be/internal/entity/customer"
	"github.com/adefarhan/warmindo-be/internal/entity/product"
	"github.com/adefarhan/warmindo-be/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Router *gin.Engine
)

func SetupTest() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := setUpTestDB()

	gin.SetMode(gin.ReleaseMode)

	Router = gin.Default()

	// Buat instance handler produk dengan dependensi repository yang diinisialisasi dengan DB
	productRepo := product.NewProductRepository(db)
	productUseCase := usecase.NewProductUseCase(productRepo)
	productHandler := http.NewProductHandler(productUseCase)

	customerRepo := customer.NewCustomerRepository(db)
	customerUseCase := usecase.NewCustomerUseCase(customerRepo)
	customerHandler := http.NewCustomerHandler(customerUseCase)

	// Routes
	Router.GET("/products", productHandler.GetProducts)
	Router.POST("/products", productHandler.CreateProduct)
	Router.GET("products/:productId", productHandler.GetProduct)
	Router.PUT("/products/:productId", productHandler.UpdateProduct)
	Router.DELETE("/products/:productId", productHandler.DeleteProduct)

	Router.POST("/customers", customerHandler.CreateCustomer)
	Router.GET("/customers", customerHandler.GetCustomers)
	Router.GET("customers/:customerId", customerHandler.GetCustomer)
	Router.PUT("/customers/:customerId", customerHandler.UpdateCustomer)
	Router.DELETE("/customers/:customerId", customerHandler.DeleteCustomer)
}

func setUpTestDB() *gorm.DB {
	// Koneksi ke database PostgreSQL
	dsn := "host=localhost user=postgres password=postgre dbname=warmindo_test port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	// Migrate model ke database
	db.AutoMigrate(&product.Product{}, &customer.Customer{})

	return db
}
