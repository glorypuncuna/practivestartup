package transaction

import (
	"bwastartup/campaign"
	"bwastartup/user"
	"errors"
	"fmt"
)

type Service interface {
	GetByCampaignId(input TransactionCampaignInput) ([]Transaction, error)
	GetByUserId(input user.User) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
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

func (s *service) GetByUserId(input user.User) ([]Transaction, error) {
	var transactions []Transaction
	transactions, err := s.repository.FindByUserId(input.ID)
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	campaign, _ := s.campaignRepository.FindById(input.CampaignId)

	transaction := Transaction{
		Amount:     input.Amount,
		CampaignId: input.CampaignId,
		User:       input.User,
		UserId:     input.User.ID,
		Campaign:   campaign,
		Status:     "Pending",
	}
	newTransaction, err := s.repository.SaveTransaction(transaction)
	if err != nil {
		return newTransaction, err
	}
	return newTransaction, nil
}
