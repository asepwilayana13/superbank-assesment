package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"myapp-go/src/controllers"
	"myapp-go/src/dto"
	"myapp-go/src/mocks/customer_mock"
	"myapp-go/src/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func setupAppAndController(mockService *customer_mock.CustomerServiceMock) (*fiber.App, *controllers.CustomerController) {
	app := fiber.New()
	controller := controllers.NewCustomerController(mockService)
	app.Post("/customers", controller.Create)
	return app, controller
}

func TestCustomerController_Create(t *testing.T) {
	
	t.Run("testCreateCustomer_Success", func(t *testing.T) {
	mockService := new(customer_mock.CustomerServiceMock)
	app, _ := setupAppAndController(mockService)
	controller := &controllers.CustomerController{Service: mockService}

	app.Post("/customer", func(c *fiber.Ctx) error {
		c.Locals("user_id", "some-user-id") // Mock JWT middleware
		return controller.Create(c)
	})

	// Setup request
	reqBody := dto.CreateCustomerRequest{
		FullName:    "John Doe",
		UserName:    "jhondoe",
		PhoneNumber: "08123456789",
		Address:     "Jl. Test",
		DateOfBirth: "2000-01-01",
		Email: "jhondoe@email.com",
		AccountType: "saving",
		Balance: 100000,
		Currency: "IDR",
	}
	reqBytes, _ := json.Marshal(reqBody)

	// Mock Service response
	mockCustomer := &models.Customer{FullName: "John Doe"}
	mockUser := &models.User{Username: "johndoe"}
	mockService.On("Create", reqBody).Return(mockCustomer, mockUser, nil)

	// Buat HTTP request
	req := httptest.NewRequest(http.MethodPost, "/customer", bytes.NewReader(reqBytes))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	// Assertion
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	mockService.AssertExpectations(t)
	})

	t.Run("TestCreateCustomer_Unauthorized", func(t *testing.T) {
		mockService := new(customer_mock.CustomerServiceMock)
		app, _ := setupAppAndController(mockService)
	
		req := httptest.NewRequest(http.MethodPost, "/customers", nil)
		req.Header.Set("Content-Type", "application/json")
	
		resp, _ := app.Test(req)
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})
	t.Run("TestCreateCustomer_BodyParserError", func(t *testing.T) {
		mockService := new(customer_mock.CustomerServiceMock)
		app, _ := setupAppAndController(mockService)
		controller := &controllers.CustomerController{Service: mockService}

		app.Post("/customer", func(c *fiber.Ctx) error {
			c.Locals("user_id", "some-user-id") // Mock JWT middleware
			return controller.Create(c)
		})

		req := httptest.NewRequest(http.MethodPost, "/customer", bytes.NewReader([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})
	t.Run("TestCreateCustomer_ValidationError", func(t *testing.T) {
		mockService := new(customer_mock.CustomerServiceMock)
		app, _ := setupAppAndController(mockService)
		controller := &controllers.CustomerController{Service: mockService}

		app.Post("/customer", func(c *fiber.Ctx) error {
			c.Locals("user_id", "some-user-id") // Mock JWT middleware
			return controller.Create(c)
		})

		// Kirim JSON kosong biar gagal validasi
		body, _ := json.Marshal(dto.CreateCustomerRequest{})
		req := httptest.NewRequest(http.MethodPost, "/customer", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("user_id", "mock-user-id")

		resp, _ := app.Test(req)

		assert.Equal(t,http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("TestCreateCustomer_ServiceError", func(t *testing.T) {
		mockService := new(customer_mock.CustomerServiceMock)
    app := fiber.New()

    // Inject controller dengan mock service
    controller := &controllers.CustomerController{Service: mockService}

    // Register route dan simulasikan JWT middleware
    app.Post("/customer", func(c *fiber.Ctx) error {
        c.Locals("user_id", "mock-user-id") // Simulasi user_id dari JWT
        return controller.Create(c)
    })

    // Setup request yang valid
    reqBody := dto.CreateCustomerRequest{
        FullName:    "John Doe",
        UserName:    "johndoe",
        PhoneNumber: "08123456789",
        Address:     "Jl. Test",
        DateOfBirth: "2000-01-01",
        Email:       "jhondoe@email.com",
        AccountType: "saving",
        Balance:     100000,
        Currency:    "IDR",
    }

    // Mock service: return error
		mockCustomer := &models.Customer{FullName: "John Doe"}
		mockUser := &models.User{Username: "johndoe"}
    mockService.On("Create", reqBody).Return(mockCustomer, mockUser, errors.New("something went wrong"))

    // Buat HTTP request
    body, _ := json.Marshal(reqBody)
    req := httptest.NewRequest(http.MethodPost, "/customer", bytes.NewReader(body))
    req.Header.Set("Content-Type", "application/json")

    // Jalankan test
    resp, _ := app.Test(req)
    assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
    mockService.AssertExpectations(t)
	})
}

func TestCustomerController_Update(t *testing.T) {
	t.Run("TestUpdateCustomer_Success", func(t *testing.T) {
		mockService := new(customer_mock.CustomerServiceMock)
		app := fiber.New()

		controller := &controllers.CustomerController{Service: mockService}

		app.Put("/customer/:id", func(c *fiber.Ctx) error {
			c.Locals("user_id", "mock-user-id") // Simulasi user_id dari JWT
			return controller.Update(c)
		})

		// Request & parameter
		userId := "123"
		reqBody := dto.UpdateCustomerRequest{
			FullName:    "John Doe",
			UserName:    "johndoe",
			PhoneNumber: "08123456789",
			Address:     "Jl. Test",
			DateOfBirth: "2000-01-01",
			Email:       "jhondoe@email.com",
			AccountType: "saving",
			Currency:    "IDR",
		}

		expectedCustomer := &models.Customer{
			FullName: "Updated Name",
			PhoneNumber: "08123456789",
			Address:     "Jl. Test",
		}

		mockService.On("Update", reqBody, userId).Return(expectedCustomer, nil)

		body, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPut, "/customer/"+userId, bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("TestUpdateCustomer_BodyParserError", func(t *testing.T) {
		mockService := new(customer_mock.CustomerServiceMock)
		app := fiber.New()
		controller := &controllers.CustomerController{Service: mockService}

		app.Put("/customer/:id", func(c *fiber.Ctx) error {
			return controller.Update(c)
		})

		req := httptest.NewRequest(http.MethodPut, "/customer/123", bytes.NewReader([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("TestUpdateCustomer_ServiceError", func(t *testing.T) {
		mockService := new(customer_mock.CustomerServiceMock)
		app := fiber.New()
		controller := &controllers.CustomerController{Service: mockService}

		app.Put("/customer/:id", func(c *fiber.Ctx) error {
			c.Locals("user_id", "mock-user-id") // Simulasi user_id dari JWT
			return controller.Update(c)
		})

		userId := "123"
		reqBody := dto.UpdateCustomerRequest{
			FullName:    "Error Test",
			PhoneNumber: "08123456789",
			Address:     "Test",
			DateOfBirth: "2000-01-01",
		}

		expectedCustomer := &models.Customer{
		
		}

		mockService.On("Update", reqBody, userId).Return(expectedCustomer, errors.New("update failed"))

		body, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPut, "/customer/"+userId, bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestCustomerController_GetCustomerInfo(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockService := new(customer_mock.CustomerServiceMock)
		app := fiber.New()
		controller := &controllers.CustomerController{Service: mockService}

		app.Get("/customer/me", func(c *fiber.Ctx) error {
			c.Locals("user_id", "mock-user-id") // Simulasi JWT Middleware
			return controller.GetCustomerInfo(c)
		})

		expectedCustomer := &models.Customer{
			ID: uuid.New(),
			FullName:    "Error Test",
			PhoneNumber: "08123456789",
			Address:     "Test",
			// DateOfBirth: "2000-01-01",
		}

		mockService.On("GetByUserID", "mock-user-id").Return(expectedCustomer, nil)

		req := httptest.NewRequest(http.MethodGet, "/customer/me", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("Unauthorized - Missing user_id", func(t *testing.T) {
		mockService := new(customer_mock.CustomerServiceMock)
		app := fiber.New()
		controller := &controllers.CustomerController{Service: mockService}

		app.Get("/customer/info", controller.GetCustomerInfo)

		req := httptest.NewRequest(http.MethodGet, "/customer/info", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("Unauthorized - Invalid user_id type", func(t *testing.T) {
		mockService := new(customer_mock.CustomerServiceMock)
		app := fiber.New()
		controller := &controllers.CustomerController{Service: mockService}

		app.Get("/customer/info", func(c *fiber.Ctx) error {
			c.Locals("user_id", 12345) // Bukan string
			return controller.GetCustomerInfo(c)
		})

		req := httptest.NewRequest(http.MethodGet, "/customer/info", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("Service Error", func(t *testing.T) {
		mockService := new(customer_mock.CustomerServiceMock)
		app := fiber.New()
		controller := &controllers.CustomerController{Service: mockService}

		app.Get("/customer/info", func(c *fiber.Ctx) error {
			c.Locals("user_id", "mock-user-id")
			return controller.GetCustomerInfo(c)
		})

		expectedCustomer := &models.Customer{
		
		}

		mockService.On("GetByUserID", "mock-user-id").Return(expectedCustomer, errors.New("customer not found"))

		req := httptest.NewRequest(http.MethodGet, "/customer/info", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestCustomerController_GetAllCustomer(t *testing.T) {
	t.Run("Success Get All Customer", func(t *testing.T) {
		mockService := new(customer_mock.CustomerServiceMock)
		app := fiber.New()
		controller := &controllers.CustomerController{Service: mockService}

		app.Get("/customers", controller.GetAllCustomer)

		// Dummy data
		customers := []*models.Customer{
			{ID:  uuid.New(), FullName: "John Doe"},
			{ID:  uuid.New(), FullName: "Jane Smith"},
		}

		customersList := &models.CustomersList{
			Customer: customers,
		}

		// Expectation
		mockService.On("GetAllCustomer", "", 10, 1).Return(customersList, nil)

		req := httptest.NewRequest(http.MethodGet, "/customers", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("Success With Query Params", func(t *testing.T) {
		mockService := new(customer_mock.CustomerServiceMock)
		app := fiber.New()
		controller := &controllers.CustomerController{Service: mockService}

		app.Get("/customers", controller.GetAllCustomer)

		customers := []*models.Customer{
		
		}

		customersList := &models.CustomersList{
			Customer: customers,
		}

		mockService.On("GetAllCustomer", "john", 5, 2).Return(customersList, nil)

		req := httptest.NewRequest(http.MethodGet, "/customers?search=john&limit=5&page=2", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("Internal Server Error from Service", func(t *testing.T) {
		mockService := new(customer_mock.CustomerServiceMock)
		app := fiber.New()
		controller := &controllers.CustomerController{Service: mockService}

		app.Get("/customers", controller.GetAllCustomer)

		customers := []*models.Customer{
		
		}

		customersList := &models.CustomersList{
			Customer: customers,
		}

		mockService.On("GetAllCustomer", "", 10, 1).Return(customersList, errors.New("DB error"))

		req := httptest.NewRequest(http.MethodGet, "/customers", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}


