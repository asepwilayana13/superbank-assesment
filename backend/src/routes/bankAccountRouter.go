package routes

import (
	"myapp-go/src/controllers"
	// "myapp-go/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

// Setup Bank Account Routes
func SetupBankAccountRoutes(app *fiber.App, BAController *controllers.BankAccountController) {
	userGroup := app.Group("/v1/bank_account")

	userGroup.Post("/create", BAController.Create)
	// userGroup.Get("/lists", BAController.GetAllBankAccount)
	userGroup.Put("/update/:id", BAController.Update)
	userGroup.Get("/detail/:id", BAController.FindOne)
	userGroup.Delete("/delete/:id", BAController.Delete)
}
