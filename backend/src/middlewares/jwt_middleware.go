package middlewares

import (
	"myapp-go/src/utils/httpError"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func JWTMiddleware(c *fiber.Ctx) error {

	bearerHeader := c.Get("Authorization")

	if bearerHeader != "" {
		headerParts := strings.Split(bearerHeader, " ")
		if len(headerParts) != 2 {
			return c.Status(fiber.StatusUnauthorized).JSON(httpError.NewUnauthorizedError(httpError.Unauthorized))
		}

		tokenString := headerParts[1]

		// Parse token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		// ketika token tidak valid
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token"})
		}

		// Ambil user_id dari token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token claims"})
		}

		// Simpan user_id di context
		c.Locals("user_id", claims["user_id"])
		return c.Next()
	}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Missing token"})
}
