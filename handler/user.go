package handler

import (
	"bwastartup/auth"
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
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
			500, "Internal Server Error", nil))
		return
	}

	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.APIResponse("Failed to authenticate token",
			500, "Internal Server Error", nil))
		return
	}
	formatter := user.FormatUser(newUser, token)

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
		response := helper.APIResponse("Email or password does not match our records", 401, "Unauthorized", errorMessage)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.APIResponse("Failed to authenticate token",
			500, "Internal Server Error", nil))
		return
	}
	formatter := user.FormatUser(newUser, token)
	response := helper.APIResponse("Login Success", 200, "Success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmail(c *gin.Context) {

	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := user.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Enter valid email", 400, "Bad Request", errorMessage)
		c.JSON(http.StatusBadRequest, response)
	}

	check, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		response := helper.APIResponse("An error has occured", 500, "Internal Server Error", err)
		c.JSON(http.StatusInternalServerError, response)
	}

	formatter := user.FormatCheck(check)
	response := helper.APIResponse("Email is available", 200, "Success", formatter)
	if !check {
		response = helper.APIResponse("Email is not available", 200, "Success", formatter)
	}
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {

	file, err := c.FormFile("avatar")
	if err != nil {
		e := gin.H{"errors": err}
		response := helper.APIResponse("Something wrong with the file", 400, "Bad Request", e)
		c.JSON(http.StatusBadRequest, response)
	}

	path := "images/" + file.Filename
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		e := gin.H{"errors": err}
		response := helper.APIResponse("Could not save file", 500, "Inbternal Server Error", e)
		c.JSON(http.StatusInternalServerError, response)
	}

	isUploaded, err := h.userService.UploadAvatar(6, path)
	if err != nil {
		e := gin.H{"errors": err}
		response := helper.APIResponse("Could not save file", 500, "Inbternal Server Error", e)
		c.JSON(http.StatusInternalServerError, response)
	}

	res := gin.H{"is_uploaded": isUploaded}
	response := helper.APIResponse("Avatar is uploaded", 200, "Success", res)
	if !isUploaded {
		response = helper.APIResponse("Avatar is not uploaded", 500, "Internal Server Error", res)
		c.JSON(http.StatusInternalServerError, response)
	}
	c.JSON(http.StatusOK, response)
}
