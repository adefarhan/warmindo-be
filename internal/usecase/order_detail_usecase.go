package usecase

import (
	"errors"
	"log"
	"time"

	"github.com/adefarhan/warmindo-be/internal/entity/order"
	orderdetail "github.com/adefarhan/warmindo-be/internal/entity/order_detail"
	"github.com/adefarhan/warmindo-be/internal/entity/product"
	"github.com/google/uuid"
)

type OrderDetailUseCase struct {
	repository        orderdetail.OrderDetailRepository
	productRepository product.ProductRepository
	orderRepository   order.OrderRepository
}

func NewOrderDetailUseCase(repository orderdetail.OrderDetailRepository, productRepository product.ProductRepository, orderRepository order.OrderRepository) *OrderDetailUseCase {
	return &OrderDetailUseCase{repository: repository, productRepository: productRepository, orderRepository: orderRepository}
}

// check order id ada, check product ada & stocknya, hitung harga, kurangi stock, save
func (uc *OrderDetailUseCase) CreateOrderDetail(orderId string, orderDetails []orderdetail.OrderDetail) ([]orderdetail.OrderDetail, error) {
	order, err := uc.orderRepository.GetOrder(orderId)
	if err != nil {
		return orderDetails, err
	}

	if order.ID == "" {
		log.Printf("Order with id %s not found", orderId)
		return orderDetails, errors.New("order not found")
	}

	var totalPrice float64

	for _, orderDetail := range orderDetails {
		// check ada product dan stocknya
		product, err := uc.productRepository.GetProduct(orderDetail.ProductID)
		if err != nil {
			return orderDetails, err
		}

		if product.ID == "" {
			log.Printf("Product with id %s not found", orderDetail.ProductID)
			return orderDetails, errors.New("product not found")
		}

		if product.Stock-orderDetail.Amount < 0 {
			log.Printf("Product with id %s not have stock available", orderDetail.ProductID)
			return orderDetails, errors.New("product not have stock available")
		}
		// kurangi stock
		product.Stock = product.Stock - orderDetail.Amount

		// save product
		err = uc.productRepository.SaveProduct(product)
		if err != nil {
			return orderDetails, err
		}

		// create order detail
		orderDetail.ID = uuid.NewString()
		orderDetail.OrderID = orderId

		err = uc.repository.CreateOrderDetail(orderDetail)
		if err != nil {
			return orderDetails, err
		}

		// hitung harga
		totalPrice += float64(orderDetail.Amount) * product.Price
	}

	order.TotalPrice = totalPrice
	order.Status = "Waiting Payment"
	timeNow := time.Now()
	order.UpdatedAt = &timeNow

	err = uc.orderRepository.SaveOrder(order)
	if err != nil {
		return orderDetails, err
	}

	return orderDetails, nil
}
