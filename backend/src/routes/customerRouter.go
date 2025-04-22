package routes

import (
	"myapp-go/src/controllers"
	"myapp-go/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

// Setup Customer Routes
func SetupCustomerRoutes(app *fiber.App, customerController *controllers.CustomerController) {
	app.Get("/v1/customer", customerController.GetAllCustomer)
	
	userGroup := app.Group("/v1/customer", middlewares.JWTMiddleware)
	userGroup.Post("/create", customerController.Create)
	userGroup.Put("/update/:id", customerController.Update)
	userGroup.Get("/me", customerController.GetCustomerInfo)
}
