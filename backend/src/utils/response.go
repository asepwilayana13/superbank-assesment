package utils

import (
	"github.com/gofiber/fiber/v2"
)

// BaseResponse is a standard response format
type BaseResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
	Errors     interface{} `json:"errors,omitempty"`
}

// Pagination structure for paginated responses
type Pagination struct {
	TotalItems   int `json:"total_items"`
	TotalPages   int `json:"total_pages"`
	CurrentPage  int `json:"current_page"`
	ItemsPerPage int `json:"items_per_page"`
}

// SuccessResponse returns a successful response
func SuccessResponse(c *fiber.Ctx, statusCode int, message string, data interface{}, pagination *Pagination) error {
	response := BaseResponse{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
		Pagination: pagination,
	}
	return c.Status(statusCode).JSON(response)
}

// ErrorResponse returns an error response
func ErrorResponse(c *fiber.Ctx, statusCode int, message string, errors ...map[string]string) error {
	response := fiber.Map{
		"status_code": statusCode,
		"message":     message,
	}

	if len(errors) > 0 {
		response["errors"] = errors[0]
	}

	return c.Status(statusCode).JSON(response)
}
