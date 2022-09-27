package transaction

import (
	"bwastartup/campaign"
	"errors"
	"fmt"
)

type Service interface {
	GetByCampaignId(input TransactionCampaignInput) ([]Transaction, error)
}

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository, campaignRepository}
}

func (s *service) GetByCampaignId(input TransactionCampaignInput) ([]Transaction, error) {

	campaign, err := s.campaignRepository.FindById(input.ID)
	fmt.Println(campaign.UserId)
	fmt.Println(input.User.ID)
	if campaign.UserId != input.User.ID || err != nil {
		return nil, errors.New("Unauthorized")
	}

	transactions, err := s.repository.FindByCampaignId(input.ID)
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}
