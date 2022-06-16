package entities

type PaymentInfor interface {
	GetAccountBalance(id string) float64
	GetCurrency(id string) string
	GetLastTransaction(id string) (float64, float64)
}
