package routes

import (
	"myapp-go/src/controllers"
	"myapp-go/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

// Setup Pockets Routes
func SetupPocketRoutes(app *fiber.App, pocketController *controllers.PocketController) {
	userGroup := app.Group("/v1/pocket", middlewares.JWTMiddleware)

	userGroup.Post("/create", pocketController.Create)
	userGroup.Get("/lists/:bankaccountid", pocketController.GetAll)
	userGroup.Put("/update/:pocketId", pocketController.Update)
	userGroup.Get("/detail/:pocketId", pocketController.FindOne)
	userGroup.Delete("/delete/:pocketId", pocketController.Delete)
}
