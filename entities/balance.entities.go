package entities

import "interface/database"

type Balance struct {
	Amount float64 `json:"amount"`
}

func (balance Balance) GetAccountBalance(id string) float64 {

	var Amount float64
	var creditAmount float64
	var debitAmount float64

	database.Database.Db.Raw("SELECT SUM(debit_amount) FROM transactions WHERE account_id = ?", id).Scan(&debitAmount)

	database.Database.Db.Raw("SELECT SUM(credit_amount) FROM transactions WHERE account_id = ?", id).Scan(&creditAmount)

	Amount = creditAmount - debitAmount
	return Amount
}
