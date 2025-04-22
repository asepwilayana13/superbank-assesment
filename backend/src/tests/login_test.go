package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"

	// "github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"myapp-go/src/container"
	"myapp-go/src/database"
	"myapp-go/src/dto"

	// "myapp-go/src/models"
	"myapp-go/src/routes"
)

var app *fiber.App

func TestMain(m *testing.M) {

	database.ConnectDB()

	app = fiber.New()

	
	// Inisialisasi container
	cont := container.InitContainer()

	// Setup semua routes dari satu tempat
	routes.SetupRoutes(app, cont)
	// routes.SetupRoutes(app) // Load semua route aplikasi

	code := m.Run()

	// Cleanup setelah test selesai
	os.Exit(code)
}



// Test POST auth/login
func TestLogin(t *testing.T) {
	user := dto.CreateLoginRequest{Email: "test1@example.com", Password: "securepassword"}
	body, _ := json.Marshal(user)
	fmt.Println(body)

	req := httptest.NewRequest(http.MethodPost, "http://localhost:3000/v1/auth/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	fmt.Println(req)
	resp, err := app.Test(req)
	fmt.Println(resp)
	if err != nil {
		log.Fatal(err)
	}

		// 🔥 Tambahkan log untuk melihat response status dan body
		respBody, _ := io.ReadAll(resp.Body)
		log.Println("Response Status:", resp.StatusCode)
		log.Println("Response Body:", string(respBody))
	
		assert.Equal(t, http.StatusOK, resp.StatusCode)
}