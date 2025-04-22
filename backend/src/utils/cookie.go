package utils

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

// SET COOKIE
func SetCookie(c *fiber.Ctx, name, value string, expires time.Time) {
	c.Cookie(&fiber.Cookie{
		Name:     name,
		Value:    value,
		Expires:  expires,
		HTTPOnly: true,
	})
}