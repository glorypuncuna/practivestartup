package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	campaignService campaign.Service
}

func NewCampaignHandler(campaignService campaign.Service) *campaignHandler {
	return &campaignHandler{campaignService}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.campaignService.GetCampaigns(userId)
	if err != nil {
		response := helper.APIResponse("Could not retrieve campaign", 400, "Bad Request", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := campaign.FormatAllCampaign(campaigns)
	response := helper.APIResponse("Success to retrieve campaigns", 200, "Success", formatter)
	c.JSON(http.StatusOK, response)

}
