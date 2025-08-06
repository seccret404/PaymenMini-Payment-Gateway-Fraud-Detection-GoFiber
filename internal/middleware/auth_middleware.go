package middleware

import (
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "not authorized",
			})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid token",
			})
		}

		tokenStr := parts[1]

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "unauthorized",
			})
		}

		claims := token.Claims.(jwt.MapClaims)
		c.Locals("userID", claims["id"])
		c.Locals("role", claims["role"])

		return c.Next()
	}
}

func RoleCheck(role string) fiber.Handler{
	return func(c *fiber.Ctx)error{
		userRole := c.Locals("role")
		if userRole != role{
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error" : "forbidden: insufficient role",
			})
		}
		return c.Next()
	}

}
