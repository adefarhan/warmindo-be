package http

import (
	"log"
	"net/http"

	"github.com/adefarhan/warmindo-be/internal/entity/product"
	"github.com/adefarhan/warmindo-be/internal/usecase"
	"github.com/adefarhan/warmindo-be/response"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	useCase *usecase.ProductUseCase
}

func NewProductHandler(useCase *usecase.ProductUseCase) *ProductHandler {
	return &ProductHandler{useCase: useCase}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var request product.Product

	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Printf("Bad request: %s", err.Error())
		errorResponse := response.NewErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	product, err := h.useCase.CreateProduct(request)
	if err != nil {
		log.Printf("Failed to get product: %s", err)
		errorResponse := response.NewErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	successResponse := response.NewSuccessResponse(http.StatusCreated, product)

	c.JSON(http.StatusCreated, successResponse)
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, err := h.useCase.GetProducts()
	if err != nil {
		log.Printf("Failed to get products: %s", err.Error())
		errorResponse := response.NewErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	successResponse := response.NewSuccessResponse(http.StatusOK, products)

	c.JSON(http.StatusOK, successResponse)
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	productId := c.Param("productId")

	product, err := h.useCase.GetProduct(productId)
	if err != nil {
		log.Printf("Failed to get product: %s", err.Error())
		errorResponse := response.NewErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	var successResponse response.Response

	if product.ID == "" {
		successResponse = response.NewSuccessResponse(http.StatusOK, nil)
		c.JSON(http.StatusOK, successResponse)
		return
	}

	successResponse = response.NewSuccessResponse(http.StatusOK, product)

	c.JSON(http.StatusOK, successResponse)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	productId := c.Param("productId")

	var request product.Product

	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Printf("Bad request: %s", err.Error())
		errorResponse := response.NewErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	product, err := h.useCase.UpdateProduct(productId, request)
	if err != nil {
		if err.Error() == "product not found" {
			log.Printf("Product with id %s not found", request.ID)
			errorResponse := response.NewErrorResponse(http.StatusNotFound, err.Error())
			c.JSON(http.StatusNotFound, errorResponse)
			return
		}

		log.Printf("Failed to update product: %s", err.Error())
		errorResponse := response.NewErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	successResponse := response.NewSuccessResponse(http.StatusOK, product)

	c.JSON(http.StatusOK, successResponse)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	productId := c.Param("productId")

	product, err := h.useCase.DeleteProduct(productId)
	if err != nil {
		if err.Error() == "product not found" {
			log.Printf("Product with id %s not found", productId)
			errorResponse := response.NewErrorResponse(http.StatusNotFound, err.Error())
			c.JSON(http.StatusNotFound, errorResponse)
			return
		}

		log.Printf("Failed to delete product: %s", err.Error())
		errorResponse := response.NewErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	successResponse := response.NewSuccessResponse(http.StatusOK, product)

	c.JSON(http.StatusOK, successResponse)
}
