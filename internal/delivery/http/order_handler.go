package http

import (
	"log"
	"net/http"

	"github.com/adefarhan/warmindo-be/internal/entity/order"
	"github.com/adefarhan/warmindo-be/internal/usecase"
	"github.com/adefarhan/warmindo-be/response"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	useCase *usecase.OrderUseCase
}

func NewOrderHandler(usecase *usecase.OrderUseCase) *OrderHandler {
	return &OrderHandler{useCase: usecase}
}

func (u *OrderHandler) CreateOrder(c *gin.Context) {
	var request order.Order

	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Printf("Bad request: %s", err.Error())
		errorResponse := response.NewErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	order, err := u.useCase.CreateOrder(request)
	if err != nil {
		log.Printf("Failed to create order: %s", err)
		errorResponse := response.NewErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	successResponse := response.NewSuccessResponse(http.StatusCreated, order)

	c.JSON(http.StatusCreated, successResponse)
}
