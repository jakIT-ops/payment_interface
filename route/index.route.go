package route

import (
	"errors"
	"fmt"
	"interface/database"
	"interface/entities"
	"interface/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var a entities.Account
var t entities.Transaction
var b entities.Balance

// type balanceRequest struct {
// 	Id string `json:"id"`
// }

// type balanceResponse struct {
// 	A float64 `json:"amount"`
// }

func handleBalance(c *fiber.Ctx) error {
	id := c.Params("id")
	var transaction models.Transaction
	database.Database.Db.Find(&transaction, "account_id = ?", id)
	if transaction.AccountId == "" {
		return c.Status(400).JSON("account does not transaction")
	}
	amount := b.GetAccountBalance(id)
	return c.Status(200).JSON(fiber.Map{"amount": amount})
}

// func dosomething(p entities.Payment) {
// 	p.GetAccountBalance()
// }
func HandleDebitAmount(c *fiber.Ctx) error {
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
	err := t.DebitAmount(id, &transaction)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.Status(200).JSON(transaction)
}

func HandleCreditAmount(c *fiber.Ctx) error {
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
	err := t.CreditAmount(id, &transaction)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.Status(200).JSON(transaction)
}

func someFunc1(id string) {
	fmt.Println("Account id: ", id)
	fmt.Println("Currency :", a.GetCurrency(id))
	debit, credit := t.GetLastTransaction(id)
	fmt.Printf("Last Trasnaction Debit: (%0.2f, Credit: %0.2f)\n", debit, credit)
	fmt.Println("Balance : ", b.GetAccountBalance(id))
}

func someFunc(p entities.PaymentInfor, id string) {
	fmt.Println("Account id: ", id)
	fmt.Println("Currency :", p.GetCurrency(id))
	debit, credit := p.GetLastTransaction(id)
	fmt.Printf("Last Trasnaction Debit: (%0.2f, Credit: %0.2f)\n", debit, credit)
	fmt.Println("Balance : ", p.GetAccountBalance(id))
}

func handlePaymentInfor(c *fiber.Ctx) error {
	id := c.Params("id")
	//someFunc1(id)
	// someFunc(acc, id)
	var account entities.Account
	result := database.Database.Db.Find(&account, "account_id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Print()
		}
	}
	if account.AccountId == "" {
		return errors.New("account does not exist")
	}

	currency := a.GetCurrency(id)
	debit, credit := t.GetLastTransaction(id)
	amount := b.GetAccountBalance(id)

	return c.Status(200).JSON(fiber.Map{
		"Account id": id,
		"Currency":   currency,
		"Debit":      debit,
		"Credit":     credit,
		"Amount":     amount,
	})
}

func RouteInit(r *fiber.App) {
	r.Post("/account", a.CreateAccount)
	r.Get("/account", a.GetAccounts)
	r.Get("/account/:id", a.GetAccount)
	r.Delete("/account/:id", a.DeleteAccount)
	r.Put("/account/:id", a.UpdateAccount)
	// Transaction
	r.Post("/account/:id/debit", HandleDebitAmount)
	r.Post("/account/:id/credit", HandleCreditAmount)
	// Balance
	r.Get("/account/:id/balance", handleBalance)
	// Payment
	r.Get("/account/:id/paymentinfor", handlePaymentInfor)
}
