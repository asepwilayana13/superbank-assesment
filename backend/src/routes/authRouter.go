package routes

import (
	"myapp-go/src/controllers"

	"github.com/gofiber/fiber/v2"
)

// Setup User Routes
func SetupAuthRoutes(app *fiber.App, authController *controllers.AuthController) {
	authGroup := app.Group("/v1/auth")

	authGroup.Post("/register", authController.Register)
	authGroup.Post("/login", authController.Login)
	authGroup.Post("/logout", authController.Logout)
}