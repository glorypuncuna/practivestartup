package transaction

import "bwastartup/user"

type TransactionCampaignInput struct {
	ID   int `uri:"campaignId" binding:"required"`
	User user.User
}

type CreateTransactionInput struct {
	Amount     int `json:"amount" binding:"required"`
	CampaignId int `json:"campaign_id" binding:"required"`
	User       user.User
}
