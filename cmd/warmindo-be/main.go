package main

import (
	"fmt"
	"log"
	"os"

	"github.com/adefarhan/warmindo-be/internal/delivery/http"
	"github.com/adefarhan/warmindo-be/internal/entity/customer"
	"github.com/adefarhan/warmindo-be/internal/entity/order"
	orderdetail "github.com/adefarhan/warmindo-be/internal/entity/order_detail"
	"github.com/adefarhan/warmindo-be/internal/entity/product"
	"github.com/adefarhan/warmindo-be/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Koneksi ke database PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", os.Getenv("HOST_DB"), os.Getenv("USER_DB"), os.Getenv("PASSWORD_DB"), os.Getenv("NAME_DB"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	// Migrate model ke database
	db.AutoMigrate(&product.Product{}, &customer.Customer{}, &order.Order{}, &orderdetail.OrderDetail{})

	// Inisialisasi router Gin
	router := gin.Default()

	// Buat instance handler produk dengan dependensi repository yang diinisialisasi dengan DB
	productRepo := product.NewProductRepository(db)
	productUseCase := usecase.NewProductUseCase(productRepo)
	productHandler := http.NewProductHandler(productUseCase)

	customerRepo := customer.NewCustomerRepository(db)
	customerUseCase := usecase.NewCustomerUseCase(customerRepo)
	customerHandler := http.NewCustomerHandler(customerUseCase)

	orderRepo := order.NewOrderRepository(db)
	orderUseCase := usecase.NewOrderUseCase(orderRepo, customerRepo)
	orderHandler := http.NewOrderHandler(orderUseCase)

	orderDetailRepo := orderdetail.NewOrderDetailRepository(db)
	orderDetailUseCase := usecase.NewOrderDetailUseCase(orderDetailRepo, productRepo, orderRepo)
	orderDetailHandler := http.NewOrderDetailHandler(orderDetailUseCase)

	// Routes
	router.POST("/products", productHandler.CreateProduct)
	router.GET("/products", productHandler.GetProducts)
	router.GET("products/:productId", productHandler.GetProduct)
	router.PUT("/products/:productId", productHandler.UpdateProduct)
	router.DELETE("/products/:productId", productHandler.DeleteProduct)

	router.POST("/customers", customerHandler.CreateCustomer)
	router.GET("/customers", customerHandler.GetCustomers)
	router.GET("customers/:customerId", customerHandler.GetCustomer)
	router.PUT("/customers/:customerId", customerHandler.UpdateCustomer)
	router.DELETE("/customers/:customerId", customerHandler.DeleteCustomer)

	router.POST("orders", orderHandler.CreateOrder)

	router.POST("/orders-detail/:orderId", orderDetailHandler.CreateOrderDetail)

	// Mulai server
	router.Run(":8080")
}
