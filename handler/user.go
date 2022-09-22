package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {

		errors := user.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}

		response := helper.APIResponse("Cannot process your request. Bad request.",
			404, "not found", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.APIResponse("Cannot process your request. Bad request.",
			404, "not found", nil))
		return
	}

	formatter := user.FormatUser(newUser, "tokentokentoken")

	c.JSON(http.StatusOK, helper.APIResponse("User has been registered.",
		200, "success", formatter))

}
