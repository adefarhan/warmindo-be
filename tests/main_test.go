package tests

import (
	"log"

	"github.com/adefarhan/warmindo-be/internal/delivery/http"
	"github.com/adefarhan/warmindo-be/internal/entity/product"
	"github.com/adefarhan/warmindo-be/internal/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	Router *gin.Engine
)

func SetupTest() {
	DB = setUpTestDB()

	gin.SetMode(gin.ReleaseMode)

	Router = gin.Default()

	// Buat instance handler produk dengan dependensi repository yang diinisialisasi dengan DB
	productRepo := product.NewProductRepository(DB)
	productUseCase := usecase.NewProductUseCase(productRepo)
	productHandler := http.NewProductHandler(productUseCase)

	// Routes
	Router.GET("/products", productHandler.GetProducts)
	Router.POST("/products", productHandler.CreateProduct)
	Router.GET("products/:productId", productHandler.GetProduct)
	Router.PUT("/products/:productId", productHandler.UpdateProduct)
	Router.DELETE("/products/:productId", productHandler.DeleteProduct)
}

func setUpTestDB() *gorm.DB {
	// Koneksi ke database PostgreSQL
	dsn := "host=localhost user=postgres password=postgre dbname=warmindo_test port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	// Migrate model ke database
	db.AutoMigrate(&product.Product{})

	return db
}
