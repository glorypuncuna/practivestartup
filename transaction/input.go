package transaction

import "bwastartup/user"

type TransactionCampaignInput struct {
	ID   int `uri:"campaignId" binding:"required"`
	User user.User
}
