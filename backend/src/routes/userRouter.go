package routes

import (
	"myapp-go/src/controllers"

	"github.com/gofiber/fiber/v2"
)

// Setup User Routes
func SetupUserRoutes(app *fiber.App, userController *controllers.UserController) {
	userGroup := app.Group("/v1/users")

	userGroup.Get("/", userController.GetAllUsers)
}
