package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/hidalgo27/app-g1/pkg/interfaces"
)

type ProductMiddleware struct {
	repo interfaces.ProductRepository
}

func NewProductMiddleware(repo interfaces.ProductRepository) *ProductMiddleware {
	return &ProductMiddleware{repo: repo}
}

func (pm *ProductMiddleware) Handle(requiredProduct string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Verificar si UserID está en el contexto
		userID, ok := c.Locals("UserID").(uint)
		if !ok {
			log.Warn("UserID not found in context")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		log.Info("Checking access for UserID:", userID, "to product:", requiredProduct)

		// Verificar si el usuario tiene acceso al producto
		//hasAccess, err := pm.repo.UserHasAccessToProduct(userID, requiredProduct)
		hasAccess, err := pm.repo.UserHasAccessToProductAndRole(userID, requiredProduct)
		if err != nil {
			log.Error("Failed to validate product access:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to validate product access",
			})
		}

		if !hasAccess {
			log.Warn("Access denied for UserID:", userID, "to product:", requiredProduct)
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error":            "Access denied to product",
				"required_product": requiredProduct,
			})
		}

		log.Info("Access granted for UserID:", userID, "to product:", requiredProduct)
		return c.Next()
	}
}
