package campaign

import (
	"bwastartup/user"
)

type Service interface {
	GetCampaigns(userId int) ([]Campaign, error)
	GetCampaignById(input CampaignDetailInput) (Campaign, error)
	CreateCampaign(user user.User, input CreateCampaignInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(userId int) ([]Campaign, error) {

	if userId != 0 {
		campaigns, err := s.repository.FindByUserId(userId)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}

	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil

}

func (s *service) GetCampaignById(input CampaignDetailInput) (Campaign, error) {
	campaign, err := s.repository.FindById(input.ID)
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (s *service) CreateCampaign(user user.User, input CreateCampaignInput) (Campaign, error) {

	campaign := Campaign{
		UserId:           user.ID,
		Name:             input.Name,
		ShortDescription: input.ShortDescription,
		Description:      input.Description,
		Perks:            input.Perks,
		Slug:             input.Slug,
		CampaignImages:   input.CampaignImages,
		User:             user,
	}

	newCampaign, err := s.repository.SaveCampaign(campaign)

	if err != nil {
		return newCampaign, err
	}
	return newCampaign, nil
}
