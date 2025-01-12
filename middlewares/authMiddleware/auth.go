package authMiddleware

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Auth(c *fiber.Ctx) error {
	cookie := c.Cookies("auth-token")
	if cookie == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	jwtSecret := os.Getenv("JWT_SECRET")

	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid token")
	}

	userId, ok := claims["userId"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid token claims")
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid token claims")
	}

	if time.Now().Unix() > int64(exp) {
		return c.Status(fiber.StatusUnauthorized).SendString("Token has expired")
	}

	log.Println("User ID:", userId)

	c.Locals("userId", userId)

	return c.Next()
}
