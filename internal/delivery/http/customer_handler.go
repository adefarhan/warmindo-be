package http

import (
	"log"
	"net/http"

	"github.com/adefarhan/warmindo-be/internal/entity/customer"
	"github.com/adefarhan/warmindo-be/internal/usecase"
	"github.com/adefarhan/warmindo-be/response"
	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	useCase *usecase.CustomerUseCase
}

func NewCustomerHandler(useCase *usecase.CustomerUseCase) *CustomerHandler {
	return &CustomerHandler{useCase: useCase}
}

func (u *CustomerHandler) CreateCustomer(c *gin.Context) {
	var request customer.Customer

	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Printf("Bad request: %s", err.Error())
		errorResponse := response.NewErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	customer, err := u.useCase.CreateCustomer(request)
	if err != nil {
		log.Printf("Failed to create customer: %s", err)
		errorResponse := response.NewErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	successResponse := response.NewSuccessResponse(http.StatusCreated, customer)

	c.JSON(http.StatusCreated, successResponse)
}

func (u *CustomerHandler) GetCustomers(c *gin.Context) {
	customers, err := u.useCase.GetCustomers()
	if err != nil {
		log.Printf("Failed to get customers: %s", err)
		errorResponse := response.NewErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	successResponse := response.NewSuccessResponse(http.StatusOK, customers)

	c.JSON(http.StatusOK, successResponse)
}

func (u *CustomerHandler) GetCustomer(c *gin.Context) {
	customerId := c.Param("customerId")

	customer, err := u.useCase.GetCustomer(customerId)
	if err != nil {
		log.Printf("Failed to get customer: %s", err)
		errorResponse := response.NewErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	var successResponse response.Response

	if customer.ID == "" {
		successResponse = response.NewSuccessResponse(http.StatusOK, nil)
		c.JSON(http.StatusOK, successResponse)
		return
	}

	successResponse = response.NewSuccessResponse(http.StatusOK, customer)

	c.JSON(http.StatusOK, successResponse)
}

func (u *CustomerHandler) UpdateCustomer(c *gin.Context) {
	customerId := c.Param("customerId")

	var request customer.Customer

	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Printf("Bad request: %s", err.Error())
		errorResponse := response.NewErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	customer, err := u.useCase.UpdateCustomer(customerId, request)
	if err != nil {
		if err.Error() == "customer not found" {
			log.Printf("Customer with id %s not found", customerId)
			errorResponse := response.NewErrorResponse(http.StatusNotFound, err.Error())
			c.JSON(http.StatusNotFound, errorResponse)
			return
		}

		log.Printf("Failed to update user: %s", err.Error())
		errorResponse := response.NewErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	successResponse := response.NewSuccessResponse(http.StatusOK, customer)

	c.JSON(http.StatusOK, successResponse)
}

func (u *CustomerHandler) DeleteCustomer(c *gin.Context) {
	customerId := c.Param("customerId")

	customer, err := u.useCase.DeletCustomer(customerId)
	if err != nil {
		if err.Error() == "customer not found" {
			log.Printf("Customer with id %s not found", customerId)
			errorResponse := response.NewErrorResponse(http.StatusNotFound, err.Error())
			c.JSON(http.StatusNotFound, errorResponse)
			return
		}

		log.Printf("Failed to de;ete user: %s", err.Error())
		errorResponse := response.NewErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	successResponse := response.NewSuccessResponse(http.StatusOK, customer)

	c.JSON(http.StatusOK, successResponse)
}
