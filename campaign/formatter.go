package campaign

type CampaignFormatter struct {
	ID               int    `json:"id"`
	Title            string `json:"title"`
	ShortDescription string `json:"short_description"`
	ImageUrl         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	UserID           int    `json:"userId"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	url := ""
	if len(campaign.CampaignImages) > 0 {
		url = campaign.CampaignImages[0].FileName
	}

	formatter := CampaignFormatter{
		ID:               campaign.ID,
		Title:            campaign.Name,
		ShortDescription: campaign.ShortDescription,
		ImageUrl:         url,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		UserID:           campaign.UserId,
	}
	return formatter
}

func FormatAllCampaign(campaign []Campaign) []CampaignFormatter {
	var campaignFormatter []CampaignFormatter
	for _, c := range campaign {
		formatter := FormatCampaign(c)
		campaignFormatter = append(campaignFormatter, formatter)
	}

	return campaignFormatter
}
