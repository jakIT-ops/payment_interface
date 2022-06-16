package entities

import (
	"errors"
	"interface/database"
	"interface/models"

	"math/rand"

	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Account struct {
	ID        int32  `gorm:"primaryKey"`
	AccountId string `json:"id"`
	Nickname  string `json:"nick_name"`
	Currency  string `json:"currency"`
}

// interface functions
func (Account Account) GetAccountBalance(id string) float64 {

	var Amount float64
	var creditAmount float64
	var debitAmount float64

	database.Database.Db.Raw("SELECT SUM(debit_amount) FROM transactions WHERE account_id = ?", id).Scan(&debitAmount)

	database.Database.Db.Raw("SELECT SUM(credit_amount) FROM transactions WHERE account_id = ?", id).Scan(&creditAmount)

	Amount = creditAmount - debitAmount
	return Amount
}

type ResultTra struct {
	DebitAmount  float64 `json:"debit_amount"`
	CreditAmount float64 `json:"credit_amount"`
}

func (Account Account) GetLastTransaction(id string) (float64, float64) {
	var res ResultTra

	database.Database.Db.Raw("SELECT debit_amount, credit_amount FROM transactions WHERE account_id = ? ORDER BY transaction_id DESC LIMIT 1", id).Scan(&res)
	return res.DebitAmount, res.CreditAmount
}

func (acc Account) GetCurrency(id string) string {
	var res string

	database.Database.Db.Raw("SELECT currency FROM accounts WHERE account_id = ?", id).Scan(&res)
	return res
}

// CRUD
func (acc Account) GetAccounts(c *fiber.Ctx) error {
	accounts := []models.Account{}
	database.Database.Db.Find(&accounts)
	return c.Status(200).JSON(accounts)
}

// find account
func findAccount(id string, account *Account) error {
	database.Database.Db.Find(&account, "account_id = ?", id)

	if account.AccountId == "" {
		return errors.New("account does not exist")
	}
	// if err := database.Database.Db.Find(&account, "account_id = ?", id).Error; err != nil {
	// 	return errors.New("account does not exist")
	// }
	return nil
}

func (acc Account) GetAccount(c *fiber.Ctx) error {
	id := c.Params("id")
	var account Account
	if err := findAccount(id, &account); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.Status(200).JSON(account)
}

func (acc Account) CreateAccount(c *fiber.Ctx) error {
	var account Account

	if err := c.BodyParser(&account); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&account)
	return c.Status(200).JSON(account)
}

func (acc Account) UpdateAccount(c *fiber.Ctx) error {
	id := c.Params("id")

	var account Account
	err := findAccount(id, &account)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateAccount struct {
		Nickname string `json:"nick_name"`
		Currency string `json:"currency"`
	}

	var updateData UpdateAccount

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	account.Nickname = updateData.Nickname
	account.Currency = updateData.Currency

	database.Database.Db.Save(&account)

	return c.Status(200).JSON(account)
}

func (acc Account) DeleteAccount(c *fiber.Ctx) error {
	id := c.Params("id")

	var account Account

	err := findAccount(id, &account)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&account).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).JSON("Successfully deleted Account")
}

// Genereate Account Id
func (a *Account) BeforeCreate(tx *gorm.DB) error {
	rand.Seed(time.Now().UnixNano())
	a.AccountId = generateAccountID(10)
	if a.AccountId == "" {
		return errors.New("cant't save invalid data")
	}
	return nil
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func generateAccountID(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
