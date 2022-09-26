package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"bwastartup/user"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	campaignService campaign.Service
	userService     user.Service
}

func NewCampaignHandler(campaignService campaign.Service, userService user.Service) *campaignHandler {
	return &campaignHandler{campaignService, userService}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.campaignService.GetCampaigns(userId)
	if err != nil {
		errorMessage := gin.H{"error": err}
		response := helper.APIResponse("Could not retrieve campaign", 400, "Bad Request", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := campaign.FormatAllCampaign(campaigns)
	response := helper.APIResponse("Success to retrieve campaigns", 200, "Success", formatter)
	c.JSON(http.StatusOK, response)

}

func (h *campaignHandler) GetCampaignById(c *gin.Context) {
	var input campaign.CampaignDetailInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		errorMessage := gin.H{"error": err}
		response := helper.APIResponse("Could not process yout input", 404, "Bad Request", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignDetail, err := h.campaignService.GetCampaignById(input)
	if err != nil {
		errorMessage := gin.H{"error": err}
		response := helper.APIResponse("Could not find the campaign", 404, "Bad Request", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := h.userService.GetUserById(campaignDetail.UserId)
	if err != nil {
		errorMessage := gin.H{"error": err}
		response := helper.APIResponse("Could not find the campaign", 404, "Bad Request", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := campaign.FormatDetailCampaign(campaignDetail, user)
	response := helper.APIResponse("Campaign found", 200, "Success", formatter)
	c.JSON(http.StatusOK, response)

}

func (h *campaignHandler) CreateCampaign(c *gin.Context) {

	currentUser := c.MustGet("currentUser").(user.User)
	var input campaign.CreateCampaignInput
	err := c.ShouldBindJSON(&input)
	input.User = currentUser

	if err != nil {
		errors := user.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Could not process yout input", 402, "Unprocessable Entity", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newCampaign, err := h.campaignService.CreateCampaign(input)
	if err != nil {
		errorMessage := gin.H{"error": err}
		response := helper.APIResponse("Failed to create the campaign", 400, "Bad Request", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := campaign.FormatCampaign(newCampaign)
	response := helper.APIResponse("Campaign Created", 200, "Success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) UpdateCampaign(c *gin.Context) {

	currentUser := c.MustGet("currentUser").(user.User)

	var id campaign.CampaignDetailInput
	err := c.ShouldBindUri(&id)
	if err != nil {
		errors := user.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Could not process yout input", 402, "Unprocessable Entity", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var input campaign.CreateCampaignInput
	err = c.ShouldBindJSON(&input)
	input.User = currentUser

	if err != nil {
		errors := user.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Could not process yout input", 402, "Unprocessable Entity", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newCampaign, err := h.campaignService.UpdateCampaign(id.ID, input)
	if err != nil {
		errorMessage := gin.H{"error": err}
		response := helper.APIResponse("Failed to create the campaign", 400, "Bad Request", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := campaign.FormatCampaign(newCampaign)
	response := helper.APIResponse("Campaign Created", 200, "Success", formatter)
	c.JSON(http.StatusOK, response)

}

func (h *campaignHandler) CreateImage(c *gin.Context) {

	currentUser := c.MustGet("currentUser").(user.User)

	var input campaign.CampaignImageInput

	err := c.ShouldBind(&input)
	if err != nil {
		errors := user.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Could not process yout input", 402, "Unprocessable Entity", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	campaignId := campaign.CampaignDetailInput{
		ID: input.CampaignID,
	}

	campaign, err := h.campaignService.GetCampaignById(campaignId)
	if err != nil {
		errors := user.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Could not process yout input", 402, "Unprocessable Entity", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if campaign.User != currentUser {
		errorMessage := gin.H{"error": errors.New("Unauthorized")}
		response := helper.APIResponse("You are not an authorized user", 401, "Unauthorized", errorMessage)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		errorMessage := gin.H{"error": err}
		response := helper.APIResponse("Could not process yout input", 402, "Unprocessable Entity", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	path := fmt.Sprintf("images/%d-%s", currentUser.ID, file.Filename)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		errorMessage := gin.H{"error": err}
		response := helper.APIResponse("Could not process yout input", 402, "Unprocessable Entity", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	_, err = h.campaignService.SaveCampaignImage(input, path)
	if err != nil {
		errorMessage := gin.H{"error": err}
		response := helper.APIResponse("Could not process yout input", 402, "Unprocessable Entity", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	message := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Success", 200, "OK", message)
	c.JSON(http.StatusOK, response)

}
