package controllers

import (
	"myapp-go/src/services"
	"myapp-go/src/utils"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	Service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{Service: service}
}

// Get all users With pagination
func (h *UserController) GetAllUsers(c *fiber.Ctx) error {
	// Ambil query param limit & page (default: 10 & 1)
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	page, _ := strconv.Atoi(c.Query("page", "1"))

	// Get All User
	usersList, err := h.Service.GetAllUsers(limit, page)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch users")
	}

	// response
	return utils.SuccessResponse(c, http.StatusOK, "Users retrieved successfully", usersList, nil)
}
