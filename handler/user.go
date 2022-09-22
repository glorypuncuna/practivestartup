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

		response := helper.APIResponse("Cannot process your request. Bad request",
			422, "Unprocesable Entity", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.APIResponse("Failed to register the account",
			500, "Internasl Server Error", nil))
		return
	}

	formatter := user.FormatUser(newUser, "tokentokentoken")

	c.JSON(http.StatusOK, helper.APIResponse("User has been registered.",
		200, "success", formatter))

}

func (h *userHandler) LoginUser(c *gin.Context) {

	var input user.LoginUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := user.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}

		response := helper.APIResponse("Cannot process your request. Bad request",
			422, "Unprocesable Entity", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.LoginUser(input)
	if err != nil {
		errors := err.Error()
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Password does not match our records", 401, "Unauthorized", errorMessage)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	formatter := user.FormatUser(newUser, "tokentokentoken")
	response := helper.APIResponse("Login Success", 200, "Success", formatter)
	c.JSON(http.StatusOK, response)
}
