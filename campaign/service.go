package campaign

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userId int) ([]Campaign, error)
	GetCampaignById(input CampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	UpdateCampaign(id int, input CreateCampaignInput) (Campaign, error)
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

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {

	stringSlug := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	campaign := Campaign{
		UserId:           input.User.ID,
		Name:             input.Name,
		ShortDescription: input.ShortDescription,
		Description:      input.Description,
		Perks:            input.Perks,
		GoalAmount:       input.GoalAmount,
		Slug:             slug.Make(stringSlug),
		User:             input.User,
	}

	newCampaign, err := s.repository.SaveCampaign(campaign)

	if err != nil {
		return newCampaign, err
	}
	return newCampaign, nil
}

func (s *service) UpdateCampaign(id int, input CreateCampaignInput) (Campaign, error) {
	campaign, err := s.repository.FindById(id)
	if err != nil {
		return campaign, err
	}

	// campaign := Campaign{
	// 	ID:               id,
	// 	UserId:           input.User.ID,
	// 	Name:             input.Name,
	// 	ShortDescription: input.ShortDescription,
	// 	Description:      input.Description,
	// 	Perks:            input.Perks,
	// 	GoalAmount:       input.GoalAmount,
	// 	User:             input.User,
	// }

	campaign.UserId = input.User.ID
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.Perks = input.Perks
	campaign.GoalAmount = input.GoalAmount
	if campaign.User != input.User {
		return campaign, errors.New("Unathorized")
	}

	newCampaign, err := s.repository.UpdateCampaign(campaign)
	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil
}
