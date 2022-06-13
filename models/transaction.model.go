package models

type Transaction struct {
	TransactionId          uint    `gorm:"primaryKey"`
	AccountId              string  `json:"id"`
	TransactionInformation string  `json:"transaction_infor"`
	DebitAmount            float64 `json:"debit_amount"`
	CreditAmount           float64 `json:"credit_amount"`
}
