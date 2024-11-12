package middlewares

import (
	"fmt"

	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/golang-jwt/jwt/v5"
)

func CheckMiddleware(c *fiber.Ctx) error {
	start := time.Now()
	fmt.Printf("URL = %s, Method = %s, Time = %s\n", c.OriginalURL(), c.Method(), start)
	return c.Next()
}

func IsAdmin(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	fmt.Print(claims["isAdmin"])
	if claims["isAdmin"] != "admin" {
		return c.Status(fiber.StatusUnauthorized).SendString("You are not admin")
	}
	// isAdmin := claims.i
	// return c.JSON(fiber.Map{
	// 	"name": name,

	// })
	return c.Next()
}
