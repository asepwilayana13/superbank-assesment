package main

import (
	"myapp-go/src/container"
	"myapp-go/src/database"
	"myapp-go/src/routes"
	_ "myapp-go/docs"
	"os"


	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html 
// @BasePath /v1/
func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Bisa diubah ke domain tertentu
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Koneksi ke Database
	database.ConnectDB()

	// Inisialisasi container
	cont := container.InitContainer()

	// Setup semua routes dari satu tempat
	routes.SetupRoutes(app, cont)

	// Route Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Middleware untuk menangani route yang tidak ditemukan (404)
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"code":    404,
			"message": "Route not found",
		})
	})

	// Cek argumen CLI (untuk migrasi dan seeding)
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "migrate":
			database.MigrateDB()
			return
		case "seed":
		database.SeedDB()
		return
		}
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})

	app.Listen(":5000")
}
