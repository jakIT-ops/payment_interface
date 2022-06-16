package entities

import (
	"interface/database"
	"interface/models"
)

type Transaction struct {
	models.Transaction
}

type Result struct {
	DebitAmount  float64 `json:"debit_amount"`
	CreditAmount float64 `json:"credit_amount"`
}

func (Transaction Transaction) GetLastTransaction(id string) (float64, float64) {
	var res Result

	database.Database.Db.Raw("SELECT debit_amount, credit_amount FROM transactions WHERE account_id = ? ORDER BY transaction_id DESC LIMIT 1", id).Scan(&res)
	return res.DebitAmount, res.CreditAmount
}

func (Transaction Transaction) DebitAmount(id string, transaction *models.Transaction) error {

	transaction.AccountId = id
	database.Database.Db.Create(&transaction)
	return nil
}

func (Transaction Transaction) CreditAmount(id string, transaction *models.Transaction) error {

	transaction.AccountId = id
	database.Database.Db.Create(&transaction)
	return nil
}
