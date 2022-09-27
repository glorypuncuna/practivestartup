package transaction

import (
	"time"
)

type TransactionByCampaignFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
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
