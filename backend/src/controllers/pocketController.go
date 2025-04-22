package controllers

import (
	// "fmt"
	"myapp-go/src/dto"
	"myapp-go/src/services"
	"myapp-go/src/utils"

	"strconv"

	"myapp-go/src/utils/httpError"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type PocketController struct {
	Service *services.PocketService
}

func NewPocketController(service *services.PocketService) *PocketController {
	return &PocketController{Service: service}
}

// Create Pocket
func (p *PocketController) Create(c *fiber.Ctx) error {

	userId := c.Locals("user_id")
	if userId == nil {
		return utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
	}

	userIDStr, ok := userId.(string)
	if !ok {
		return utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid user ID", nil)
	}

	var req dto.CreatePocketRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, httpError.ErrBadRequest, nil)
	}

	// Create Pocket
	pocket, err := p.Service.Create(req, userIDStr)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error(), nil)
	}
	
	return utils.SuccessResponse(c, http.StatusCreated, "Pocket created successfully",pocket, nil)
}

// Get all pocket by bank account dengan pagination
func (p *PocketController) GetAll(c *fiber.Ctx) error {
	// Get query param limit & page (default: 10 & 1)
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	page, _ := strconv.Atoi(c.Query("page", "1"))
	bank_account_id := c.Params("bankaccountid")


	userId := c.Locals("user_id")
	if userId == nil {
		return utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
	}

	// Get All Bank Account
	PocketList, err := p.Service.GetAll(bank_account_id, limit, page)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch users")
	}

	// response
	return utils.SuccessResponse(c, http.StatusOK, "Get Data successfully", PocketList, nil)
}

// Update Bank Account
func (p *PocketController) Update(c *fiber.Ctx) error {
	// Get param
	pocketId := c.Params("pocketId")
	
	userId := c.Locals("user_id")
	if userId == nil {
		return utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
	}

	var req dto.UpdatePocketRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, httpError.ErrBadRequest, nil)
	}

	// update pocket
	pocket, err := p.Service.Update(req, pocketId)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error(), nil)
	}
	
	return utils.SuccessResponse(c, http.StatusOK, "Pocket Updated successfully",pocket, nil)
}

// Get Detail Pocket
func (p *PocketController) FindOne(c *fiber.Ctx) error {
	// Get param
	pocketId := c.Params("pocketId")
	
	userId := c.Locals("user_id")
	if userId == nil {
		return utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
	}

	// GET One Pocket By Id
	pocket, err := p.Service.GetOneById(pocketId)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error(), nil)
	}
	
	return utils.SuccessResponse(c, http.StatusOK, "Get Detail Pocket successfully",pocket, nil)
}

func (p *PocketController) Delete(c *fiber.Ctx) error {
	// Get param
	pocketId := c.Params("pocketId")
	
	userId := c.Locals("user_id")
	if userId == nil {
		return utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
	}

	// Register user
	err := p.Service.Delete(pocketId)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error(), nil)
	}
	
	return utils.SuccessResponse(c, http.StatusOK, "Remove Pocket successfully",nil, nil)
}

