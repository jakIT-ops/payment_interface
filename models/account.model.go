package models

type Account struct {
	ID        int32  `gorm:"primaryKey"`
	AccountId string `json:"id"`
	Nickname  string `json:"nick_name"`
	Currency  string `json:"currency"`
}
