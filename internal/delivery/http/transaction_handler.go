package http

import (
	"log"
	"net/http"

	"github.com/adefarhan/warmindo-be/internal/entity/transaction"
	"github.com/adefarhan/warmindo-be/internal/usecase"
	"github.com/adefarhan/warmindo-be/response"
	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	useCase *usecase.TransactionUseCase
}

func NewTransactionHandler(useCase *usecase.TransactionUseCase) *TransactionHandler {
	return &TransactionHandler{useCase: useCase}
}

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	var request transaction.Transaction

	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Printf("Bad request: %s", err.Error())
		errorResponse := response.NewErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	transaction, err := h.useCase.CreateTransaction(request)
	if err != nil {
		if err.Error() == "order not found" {
			errorResponse := response.NewErrorResponse(http.StatusNotFound, err.Error())
			c.JSON(http.StatusNotFound, errorResponse)
			return
		}
		errorResponse := response.NewErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	successResponse := response.NewSuccessResponse(http.StatusCreated, transaction)

	c.JSON(http.StatusCreated, successResponse)
}
