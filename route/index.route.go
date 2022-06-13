package route

import (
	"interface/entities"

	"github.com/gofiber/fiber/v2"
)

var a entities.Account
var t entities.Transaction
var b entities.Balance

func RouteInit(r *fiber.App) {
	r.Post("/account", a.CreateAccount)
	r.Get("/account", a.GetAccounts)
	r.Get("/account/:id", a.GetAccount)
	r.Delete("/account/:id", a.DeleteAccount)
	r.Put("/account/:id", a.UpdateAccount)
	// Transaction
	r.Post("/account/:id/debit", t.DebitAmount)
	r.Post("/account/:id/credit", t.CreditAmount)
	// Balance
	r.Get("/account/:id/balance", b.GetAccountBalance)
}
