package controllers

import (
	"myapp-go/src/dto"
	"myapp-go/src/services"
	"myapp-go/src/utils"
	"strconv"

	"myapp-go/src/utils/httpError"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type BankAccountController struct {
	Service *services.BankAccountService
}

func NewBankAccountController(service *services.BankAccountService) *BankAccountController {
	return &BankAccountController{Service: service}
}

// Create Bank Account
func (ba *BankAccountController) Create(c *fiber.Ctx) error {
	// Get user_id
	userId := c.Locals("user_id")
	if userId == nil {
		return utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
	}

	userIDStr, ok := userId.(string)
	if !ok {
		return utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid user ID", nil)
	}

	var req dto.CreateBARequest

	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, httpError.ErrBadRequest, nil)
	}

	// if err := req.Validate(); err != nil {
	// 	return utils.ErrorResponse(c, http.StatusBadRequest, httpError.ErrValidation, utils.ValidationErrors(err))
	// }

	// Register user
	bankAccount, err := ba.Service.Create(req, userIDStr)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error(), nil)
	}
	
	return utils.SuccessResponse(c, http.StatusOK, "Users Update successfully",bankAccount, nil)
}

// Get all bank account dengan pagination
func (h *BankAccountController) GetAllBankAccount(c *fiber.Ctx) error {
	// Ambil query param limit & page (default: 10 & 1)
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	page, _ := strconv.Atoi(c.Query("page", "1"))

	userId := c.Locals("user_id")
	if userId == nil {
		return utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
	}
	
	userIDStr, ok := userId.(string)
	if !ok {
		return utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid user ID", nil)
	}	

	// Get All Bank Account
	bankAccountList, err := h.Service.GetAllBankAccount(userIDStr, limit, page)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch users")
	}

	// response
	return utils.SuccessResponse(c, http.StatusOK, "Get Data successfully", bankAccountList, nil)
}

// Update Bank Account
func (ba *BankAccountController) Update(c *fiber.Ctx) error {
	// Get param
	bankAccountId := c.Params("id")
	
	userId := c.Locals("user_id")
	if userId == nil {
		return utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
	}

	userIDStr, ok := userId.(string)
	if !ok {
		return utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid user ID", nil)
	}

	var req dto.UpdateBARequest

	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, httpError.ErrBadRequest, nil)
	}

	// Register user
	bankAccount, err := ba.Service.Update(req, userIDStr, bankAccountId)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error(), nil)
	}
	
	return utils.SuccessResponse(c, http.StatusOK, "Users Update successfully",bankAccount, nil)
}

// Get Detail Bank Account
func (ba *BankAccountController) FindOne(c *fiber.Ctx) error {
	// Get params
	bankAccountId := c.Params("id")
	
	// userId := c.Locals("user_id")
	// if userId == nil {
	// 	return utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
	// }

	// Get One bank account
	bankAccount, err := ba.Service.GetOne(bankAccountId)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error(), nil)
	}
	
	return utils.SuccessResponse(c, http.StatusOK, "Get Bank Account Detail successfully",bankAccount, nil)
}

func (ba *BankAccountController) Delete(c *fiber.Ctx) error {
	// Get params
	bankAccountId := c.Params("id")
	
	userId := c.Locals("user_id")
	if userId == nil {
		return utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
	}

	// Delete Bank Account
	err := ba.Service.Delete(bankAccountId)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error(), nil)
	}
	
	return utils.SuccessResponse(c, http.StatusOK, "Remove Bank Account successfully",nil, nil)
}

