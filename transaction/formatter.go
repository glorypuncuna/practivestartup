package transaction

import (
	"bwastartup/campaign"
	"time"
)

type TransactionByCampaignFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type CampaignTransactionFormatter struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

type TransactionByUserFormatter struct {
	ID        int                          `json:"id"`
	Amount    int                          `json:"amount"`
	Status    string                       `json:"status"`
	CreatedAt time.Time                    `json:"created_at"`
	Campaign  CampaignTransactionFormatter `json:"campaign"`
}

type TransactionFormatter struct {
	ID         int    `json:"id"`
	CampaignId int    `json:"campaign_id"`
	UserId     int    `json:"user_id"`
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
	Code       string `json:"code"`
	PaymentUrl string `json:"payment_url"`
}

func FormatByCampaign(transaction Transaction) TransactionByCampaignFormatter {
	formatter := TransactionByCampaignFormatter{
		ID:        transaction.ID,
		Name:      transaction.User.Name,
		Amount:    transaction.Amount,
		CreatedAt: transaction.CreatedAt,
	}
	return formatter
}

func FormatAllByCampaign(transactions []Transaction) []TransactionByCampaignFormatter {
	var formatAll []TransactionByCampaignFormatter

	if len(transactions) == 0 {
		return formatAll
	}

	for _, transaction := range transactions {
		formatter := FormatByCampaign(transaction)
		formatAll = append(formatAll, formatter)
	}
	return formatAll
}

func FormatByUser(transaction Transaction) TransactionByUserFormatter {
	var images campaign.CampaignImage
	if len(transaction.Campaign.CampaignImages) != 0 {
		for _, i := range transaction.Campaign.CampaignImages {
			if i.IsPrimary == 1 {
				images = i
			}
		}
	}

	campaignFormatter := CampaignTransactionFormatter{
		Name:     transaction.Campaign.Name,
		ImageUrl: images.FileName,
	}

	formatter := TransactionByUserFormatter{
		ID:        transaction.ID,
		Amount:    transaction.Amount,
		Status:    transaction.Status,
		CreatedAt: transaction.CreatedAt,
		Campaign:  campaignFormatter,
	}
	return formatter
}

func FormatAllByUser(transactions []Transaction) []TransactionByUserFormatter {
	var formatterAll []TransactionByUserFormatter
	for _, transaction := range transactions {
		formatter := FormatByUser(transaction)
		formatterAll = append(formatterAll, formatter)
	}
	return formatterAll
}

func FormatTransaction(transaction Transaction) TransactionFormatter {
	formatter := TransactionFormatter{
		ID:         transaction.ID,
		CampaignId: transaction.CampaignId,
		UserId:     transaction.UserId,
		Amount:     transaction.Amount,
		Status:     transaction.Status,
		Code:       transaction.Code,
		PaymentUrl: transaction.PaymentUrl,
	}
	return formatter
}
