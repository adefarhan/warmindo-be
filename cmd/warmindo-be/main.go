package main

import (
	"log"

	"github.com/adefarhan/warmindo-be/internal/delivery/http"
	"github.com/adefarhan/warmindo-be/internal/entity/product"
	"github.com/adefarhan/warmindo-be/internal/entity/user"
	"github.com/adefarhan/warmindo-be/internal/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Koneksi ke database PostgreSQL
	dsn := "host=localhost user=postgres password=postgre dbname=warmindo port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	// Migrate model ke database
	db.AutoMigrate(&product.Product{}, &user.User{})

	// Inisialisasi router Gin
	router := gin.Default()

	// Buat instance handler produk dengan dependensi repository yang diinisialisasi dengan DB
	productRepo := product.NewProductRepository(db)
	productUseCase := usecase.NewProductUseCase(productRepo)
	productHandler := http.NewProductHandler(productUseCase)

	userRepo := user.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := http.NewUserHandler(userUseCase)

	// Routes
	router.POST("/products", productHandler.CreateProduct)
	router.GET("/products", productHandler.GetProducts)
	router.GET("products/:productId", productHandler.GetProduct)
	router.PUT("/products/:productId", productHandler.UpdateProduct)
	router.DELETE("/products/:productId", productHandler.DeleteProduct)

	router.POST("/users", userHandler.CreateUser)
	router.GET("/users", userHandler.GetUsers)
	router.GET("users/:userId", userHandler.GetUser)
	router.PUT("/users/:userId", userHandler.UpdateUser)
	router.DELETE("/users/:userId", userHandler.DeleteUser)

	// Mulai server
	router.Run(":8080")
}
