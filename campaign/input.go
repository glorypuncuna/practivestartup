package campaign

type CampaignDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateCampaignInput struct {
	Name             string          `json:"name" binding:"required"`
	ShortDescription string          `json:"short_description" binding:"required"`
	Description      string          `json:"description" binding:"required"`
	Perks            string          `json:"perks" binding:"required"`
	Slug             string          `json:"slug" binding:"required"`
	CampaignImages   []CampaignImage `json:"campaign_images"`
}
