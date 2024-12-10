package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func RoleMiddleware(requiredRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		roles, ok := c.Locals("Roles").([]string)
		if !ok {
			log.Warn("Roles not found in context")
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Roles not found",
			})
		}

		log.Info("User roles:", roles)

		for _, role := range roles {
			for _, requiredRole := range requiredRoles {
				if role == requiredRole {
					log.Info("Access granted with role:", role)
					return c.Next()
				}
			}
		}

		log.Warn("Access denied. Required roles:", requiredRoles)
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":          "Insufficient role permissions",
			"required_roles": requiredRoles,
		})
	}
}
