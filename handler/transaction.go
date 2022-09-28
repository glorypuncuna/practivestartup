package handler

import (
	"bwastartup/helper"
	"bwastartup/transaction"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	transactionService transaction.Service
}

func NewTransactionHandler(transactionService transaction.Service) *transactionHandler {
	return &transactionHandler{transactionService}
}

func (h *transactionHandler) GetTransactionsByCampaignId(c *gin.Context) {

	currentUser := c.MustGet("currentUser").(user.User)
	var input transaction.TransactionCampaignInput
	c.ShouldBindUri(&input)
	input.User = currentUser
	transactions, err := h.transactionService.GetByCampaignId(input)
	if err != nil {
		formatter := helper.APIResponse("Failed to fetch", 400, "Bad Request", err)
		c.JSON(http.StatusBadRequest, formatter)
		return
	}

	formatAll := transaction.FormatAllByCampaign(transactions)
	formatter := helper.APIResponse("Success to fetch transactions", 200, "Success", formatAll)
	c.JSON(http.StatusOK, formatter)
}

func (h *transactionHandler) GetTransactionsByUserId(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)

	transactions, err := h.transactionService.GetByUserId(currentUser)
	if err != nil {
		response := helper.APIResponse("Failed to fetch transactions", 500, "Internal Server Error", err)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	formatter := transaction.FormatAllByUser(transactions)
	response := helper.APIResponse("Succeed fetch transactions", 200, "Success", formatter)
	c.JSON(http.StatusOK, response)
}
