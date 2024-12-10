package controllers

import "github.com/gofiber/fiber/v2"

type ProtectedController struct{}

func NewProtectedController() *ProtectedController {
	return &ProtectedController{}
}

func (pc ProtectedController) ConfigPath(app *fiber.App) *fiber.App {
	app.Post("/", pc.ProtectedEndpoint)
	return app
}

func (pc *ProtectedController) ProtectedEndpoint(c *fiber.Ctx) error {
	// Obtener el UserID del contexto (establecido por el middleware JWT)
	userID := c.Locals("UserID").(uint)

	return c.JSON(fiber.Map{
		"message": "Access granted",
		"user": map[string]interface{}{
			"id": userID,
		},
	})
}
