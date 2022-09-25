package campaign

import (
	"bwastartup/user"
	"strings"
)

type CampaignFormatter struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageUrl         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
	UserID           int    `json:"userId"`
}

type CampaignDetailFormatter struct {
	ID               int                            `json:"id"`
	Name             string                         `json:"name"`
	ShortDescription string                         `json:"short_description"`
	ImageUrl         string                         `json:"image_url"`
	GoalAmount       int                            `json:"goal_amount"`
	CurrentAmount    int                            `json:"current_amount"`
	UserID           int                            `json:"userId"`
	Description      string                         `json:"description"`
	Perks            []string                       `json:"perks"`
	User             CampaignDetailUserFormatter    `json:"user"`
	Images           []CampaignDetailImageFormatter `json:"images"`
}

type CampaignDetailUserFormatter struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

type CampaignDetailImageFormatter struct {
	ImageUrl  string `json:"image_url"`
	IsPrimary int    `json:"is_primary"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	url := ""
	if len(campaign.CampaignImages) > 0 {
		url = campaign.CampaignImages[0].FileName
	}

	formatter := CampaignFormatter{
		ID:               campaign.ID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		ImageUrl:         url,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		Slug:             campaign.Slug,
		UserID:           campaign.UserId,
	}
	return formatter
}

func FormatDetailCampaign(campaign Campaign, user user.User) CampaignDetailFormatter {
	url := ""

	if len(campaign.CampaignImages) > 0 {
		url = campaign.CampaignImages[0].FileName
	}

	perksArray := strings.Split(campaign.Perks, ",")

	userFormatter := CampaignDetailUserFormatter{
		Name:     user.Name,
		ImageUrl: user.AvatarFileName,
	}

	var imageArrayFormatter []CampaignDetailImageFormatter
	for _, image := range campaign.CampaignImages {
		imageFormatter := CampaignDetailImageFormatter{
			ImageUrl:  image.FileName,
			IsPrimary: image.IsPrimary,
		}

		imageArrayFormatter = append(imageArrayFormatter, imageFormatter)
	}

	formatter := CampaignDetailFormatter{
		ID:               campaign.ID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		ImageUrl:         url,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		UserID:           campaign.UserId,
		Description:      campaign.Description,
		User:             userFormatter,
		Perks:            perksArray,
		Images:           imageArrayFormatter,
	}
	return formatter
}

func FormatAllCampaign(campaign []Campaign) []CampaignFormatter {
	if len(campaign) == 0 {
		return []CampaignFormatter{}
	}

	var campaignFormatter []CampaignFormatter
	for _, c := range campaign {
		formatter := FormatCampaign(c)
		campaignFormatter = append(campaignFormatter, formatter)
	}

	return campaignFormatter
}
