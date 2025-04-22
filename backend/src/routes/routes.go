package routes

import (
	"myapp-go/src/container"

	"github.com/gofiber/fiber/v2"
)

// Setup All Routes
func SetupRoutes(app *fiber.App, cont *container.Container) {
	SetupUserRoutes(app, cont.UserController)
	SetupAuthRoutes(app, cont.AuthController)
	SetupCustomerRoutes(app, cont.CustomerController)
	SetupBankAccountRoutes(app, cont.BankAccountController)
	SetupPocketRoutes(app, cont.PocketController)
}
