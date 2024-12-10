package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

//func PermissionMiddleware(requiredPermission string) fiber.Handler {
//	return func(c *fiber.Ctx) error {
//		permissions, ok := c.Locals("Permissions").([]string)
//		fmt.Println(permissions)
//		if !ok {
//			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Permissions not found"})
//		}
//
//		for _, permission := range permissions {
//			if permission == requiredPermission {
//				return c.Next()
//			}
//		}
//
//		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Insufficient permissions"})
//	}
//}

func PermissionMiddleware(requiredPermissions ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		permissions, ok := c.Locals("Permissions").([]string)
		if !ok {
			log.Warn("Permissions not found in context")
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Permissions not found",
			})
		}

		log.Info("User permissions:", permissions)

		// Comparar permisos requeridos con los del usuario
		for _, permission := range permissions {
			for _, requiredPermission := range requiredPermissions {
				if permission == requiredPermission {
					log.Info("Access granted with permission:", permission)
					return c.Next()
				}
			}
		}

		log.Warn("Access denied. Required permissions:", requiredPermissions)
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":                "Insufficient permissions",
			"required_permissions": requiredPermissions,
		})
	}
}
