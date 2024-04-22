package http

import (
	"log"
	"net/http"

	"github.com/adefarhan/warmindo-be/internal/entity/user"
	"github.com/adefarhan/warmindo-be/internal/usecase"
	"github.com/adefarhan/warmindo-be/response"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	useCase *usecase.UserUseCase
}

func NewUserHandler(useCase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{useCase: useCase}
}

func (u *UserHandler) CreateUser(c *gin.Context) {
	var request user.User

	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Printf("Bad request: %s", err.Error())
		errorResponse := response.NewErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	user, err := u.useCase.CreateUser(request)
	if err != nil {
		log.Printf("Failed to create user: %s", err)
		errorResponse := response.NewErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	successResponse := response.NewSuccessResponse(http.StatusCreated, user)

	c.JSON(http.StatusCreated, successResponse)
}

func (u *UserHandler) GetUsers(c *gin.Context) {
	users, err := u.useCase.GetUsers()
	if err != nil {
		log.Printf("Failed to get users: %s", err)
		errorResponse := response.NewErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	successResponse := response.NewSuccessResponse(http.StatusOK, users)

	c.JSON(http.StatusOK, successResponse)
}

func (u *UserHandler) GetUser(c *gin.Context) {
	userId := c.Param("userId")

	user, err := u.useCase.GetUser(userId)
	if err != nil {
		log.Printf("Failed to create user: %s", err)
		errorResponse := response.NewErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	var successResponse response.Response

	if user.ID == "" {
		successResponse = response.NewSuccessResponse(http.StatusOK, nil)
		c.JSON(http.StatusOK, successResponse)
		return
	}

	successResponse = response.NewSuccessResponse(http.StatusOK, user)

	c.JSON(http.StatusOK, successResponse)
}

func (u *UserHandler) UpdateUser(c *gin.Context) {
	userId := c.Param("userId")

	var request user.User

	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Printf("Bad request: %s", err.Error())
		errorResponse := response.NewErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	user, err := u.useCase.UpdateUser(userId, request)
	if err != nil {
		if err.Error() == "user not found" {
			log.Printf("User with id %s not found", userId)
			errorResponse := response.NewErrorResponse(http.StatusNotFound, err.Error())
			c.JSON(http.StatusNotFound, errorResponse)
			return
		}

		log.Printf("Failed to update user: %s", err.Error())
		errorResponse := response.NewErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	successResponse := response.NewSuccessResponse(http.StatusOK, user)

	c.JSON(http.StatusOK, successResponse)
}

func (u *UserHandler) DeleteUser(c *gin.Context) {
	userId := c.Param("userId")

	user, err := u.useCase.DeletUser(userId)
	if err != nil {
		if err.Error() == "user not found" {
			log.Printf("User with id %s not found", userId)
			errorResponse := response.NewErrorResponse(http.StatusNotFound, err.Error())
			c.JSON(http.StatusNotFound, errorResponse)
			return
		}

		log.Printf("Failed to de;ete user: %s", err.Error())
		errorResponse := response.NewErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	successResponse := response.NewSuccessResponse(http.StatusOK, user)

	c.JSON(http.StatusOK, successResponse)
}
