package http

import (
	"log"
	"net/http"

	orderdetail "github.com/adefarhan/warmindo-be/internal/entity/order_detail"
	"github.com/adefarhan/warmindo-be/internal/usecase"
	"github.com/adefarhan/warmindo-be/response"
	"github.com/gin-gonic/gin"
)

type OrderDetailHandler struct {
	useCase *usecase.OrderDetailUseCase
}

func NewOrderDetailHandler(useCase *usecase.OrderDetailUseCase) *OrderDetailHandler {
	return &OrderDetailHandler{useCase: useCase}
}

func (u *OrderDetailHandler) CreateOrderDetail(c *gin.Context) {
	orderId := c.Param("orderId")

	var request []orderdetail.OrderDetail

	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Printf("Bad request: %s", err.Error())
		errorResponse := response.NewErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	orderDetails, err := u.useCase.CreateOrderDetail(orderId, request)
	if err != nil {
		if err.Error() == "order not found" || err.Error() == "product not have stock available" {
			errorResponse := response.NewErrorResponse(http.StatusNotFound, err.Error())
			c.JSON(http.StatusNotFound, errorResponse)
			return
		}
		log.Printf("Failed to create customer: %s", err)
		errorResponse := response.NewErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	successResponse := response.NewSuccessResponse(http.StatusCreated, orderDetails)

	c.JSON(http.StatusCreated, successResponse)
}
