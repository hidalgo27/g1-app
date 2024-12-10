package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func RoleMiddleware(requiredRole string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		roles, ok := c.Locals("Roles").([]string)
		fmt.Println(roles)
		if !ok {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Roles not found"})
		}

		for _, role := range roles {
			if role == requiredRole {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Insufficient role permissions"})
	}
}
