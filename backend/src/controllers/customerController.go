package controllers

import (
	"myapp-go/src/dto"
	services_interface "myapp-go/src/services/interface"
	"myapp-go/src/utils"
	"myapp-go/src/utils/httpError"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	// "github.com/google/uuid"
)

type CustomerController struct {
	Service services_interface.ICustomerService
}

func NewCustomerController(service services_interface.ICustomerService) *CustomerController {
	return &CustomerController{Service: service}
}

// Create Customer
func (ctrl *CustomerController) Create(c *fiber.Ctx) error {
	// Ambil user_id dari JWT middleware
	userId := c.Locals("user_id")
	if userId == nil {
		return utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
	}

	var req dto.CreateCustomerRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, httpError.ErrBadRequest, nil)
	}

	if err := req.Validate(); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, httpError.ErrValidation, utils.ValidationErrors(err))
	}

	// Register user
	customer, user, err := ctrl.Service.Create(req)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error(), nil)
	}
	
	// return utils.SuccessResponse(c, http.StatusOK, "Users Update successfully",customer, nil)
	return utils.SuccessResponse(c, http.StatusCreated, httpError.CreatedUserSuccess, fiber.Map{
		"user":    user,
		"customer": customer,
	}, nil)
}

// Update Customer
func (cc *CustomerController) Update(c *fiber.Ctx) error {
	// ambi param
	userId := c.Params("id")
	// Ambil user_id dari JWT middleware
	// userId := c.Locals("user_id")
	// if userId == nil {
	// 	return utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
	// }

	var req dto.UpdateCustomerRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, httpError.ErrBadRequest, nil)
	}

	// Register user
	customer, err := cc.Service.Update(req, userId)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error(), nil)
	}
	
	return utils.SuccessResponse(c, http.StatusOK, "Users Update successfully",customer, nil)
}

func (cc *CustomerController) GetCustomerInfo(c *fiber.Ctx) error {
	userId := c.Locals("user_id")
	if userId == nil {
		return utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
	}

	// Konversi userId string to UUID
	userIDStr, ok := userId.(string)
	if !ok {
		return utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid user ID", nil)
	}

	// Register user
	customer, err := cc.Service.GetByUserID(userIDStr)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	// response
	return utils.SuccessResponse(c, http.StatusOK, "Users Get Data successfully",customer, nil)
}

func (cc *CustomerController) GetAllCustomer(c *fiber.Ctx) error {
	// Ambil query param limit & page (default: 10 & 1)
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	page, _ := strconv.Atoi(c.Query("page", "1"))
	search := c.Query("search", "")

	// Register user
	customer, err := cc.Service.GetAllCustomer(search, limit, page)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch users")
	}

	// response
	return utils.SuccessResponse(c, http.StatusOK, "Get Data successfully", customer, nil)
}
