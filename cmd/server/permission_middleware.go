package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func PermissionMiddleware(requiredPermission string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		permissions, ok := c.Locals("Permissions").([]string)
		fmt.Println(permissions)
		if !ok {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Permissions not found"})
		}

		for _, permission := range permissions {
			if permission == requiredPermission {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Insufficient permissions"})
	}
}
