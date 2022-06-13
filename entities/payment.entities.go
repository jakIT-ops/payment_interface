package entities

import (
	"errors"
	"interface/database"
	"interface/models"

	"github.com/gofiber/fiber/v2"
)

type Payment interface {
	DebitAmount(c *fiber.Ctx) error
	CreditAmount(c *fiber.Ctx) error
	GetAccountBalance(c *fiber.Ctx) error
}

type Balance struct {
	Amount    float64
	AccountId int32
}

type Transaction struct {
	models.Transaction
}

func (Transaction Transaction) DebitAmount(c *fiber.Ctx) error {
	id := c.Params("id")

	var account models.Account
	database.Database.Db.Find(&account, "account_id = ?", id)
	if account.AccountId == "" {
		return errors.New("account does not exist")
	}

	var transaction models.Transaction
	if err := c.BodyParser(&transaction); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	transaction.AccountId = id
	database.Database.Db.Create(&transaction)

	return c.Status(200).JSON(transaction)
}

func (Transaction Transaction) CreditAmount(c *fiber.Ctx) error {

	id := c.Params("id")

	var account models.Account
	database.Database.Db.Find(&account, "account_id = ?", id)
	if account.AccountId == "" {
		return errors.New("account does not exist")
	}

	var transaction models.Transaction
	if err := c.BodyParser(&transaction); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	transaction.AccountId = id

	database.Database.Db.Create(&transaction)
	return c.Status(200).JSON(transaction)
}

func (Balance Balance) GetAccountBalance(c *fiber.Ctx) error {
	id := c.Params("id")

	var transaction Transaction
	database.Database.Db.Find(&transaction, "account_id = ?", id)
	if transaction.AccountId == "" {
		return c.Status(400).JSON("account does not transaction")
	}

	var Amount float64
	var creditAmount float64
	var debitAmount float64

	database.Database.Db.Raw("SELECT SUM(debit_amount) FROM transactions WHERE account_id = ?", id).Scan(&debitAmount)

	database.Database.Db.Raw("SELECT SUM(credit_amount) FROM transactions WHERE account_id = ?", id).Scan(&creditAmount)

	Amount = creditAmount - debitAmount
	return c.Status(200).JSON(fiber.Map{"amount": Amount})
}
