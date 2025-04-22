package controllers

import (
	"fmt"
	"myapp-go/src/dto"
	"myapp-go/src/services"
	"myapp-go/src/utils"
	"myapp-go/src/utils/httpError"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	Service *services.AuthService
}

func NewAuthController(service *services.AuthService) *AuthController {
	return &AuthController{Service: service}
}

func (h *AuthController) Register(c *fiber.Ctx) error {
	var req dto.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, httpError.ErrBadRequest, nil)
	}

	if err := req.Validate(); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, httpError.ErrValidation, utils.ValidationErrors(err))
	}

	// Register user
	user, session, err := h.Service.RegisterUser(req.Username, req.Email, req.Password)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	// untuk set cookie
	utils.SetCookie(c, "sessionid", session.Sessionid.String(), session.Expires)

	// Response
	return utils.SuccessResponse(c, http.StatusCreated, httpError.CreatedUserSuccess, fiber.Map{
		"user":    user,
		"session": session,
	}, nil)
}

// =============== LOGIN ============== //
// Login godoc
// @Summary Login user
// @Description Melakukan login dengan Email dan Password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param data body dto.CreateLoginRequest true "Login Request Body"
// @Success 200 {object} utils.BaseResponse
// @Failure 400 {object} utils.BaseResponse
// @Failure 401 {object} utils.BaseResponse
// @Failure 500 {object} utils.BaseResponse
// @Router /auth/login [post]
func (h *AuthController) Login(c *fiber.Ctx) error {
	var req dto.CreateLoginRequest  
	fmt.Println(req);

	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, httpError.ErrBadRequest, nil)
	}

	if err := req.Validate(); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, httpError.ErrValidation, utils.ValidationErrors(err))
	}

	// Register user
	user, session, err := h.Service.LoginUser(req.Email, req.Password)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	// untuk set cookie
	utils.SetCookie(c, "sessionid", session.Sessionid.String(), session.Expires)

	// Response
	return utils.SuccessResponse(c, http.StatusOK, httpError.LoginSuccess, fiber.Map{
		"user":    user,
		"session": session,
	}, nil)
}

// Logout user
func (h *AuthController) Logout(c *fiber.Ctx) error {
	type LogoutRequest struct {
		SessionID string `json:"sessionid"`
	}

	var req LogoutRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid JSON")
	}

	// Proses logout
	err := h.Service.Logout(req.SessionID)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, err.Error())
	}

	// Hapus cookie session
	c.ClearCookie("sessionid")

	return utils.SuccessResponse(c, http.StatusOK, "Success logout", nil, nil)
}