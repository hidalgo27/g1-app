package controllers

import "github.com/gofiber/fiber/v2"

type ProductController struct{}

func NewProductController() *ProductController {
	return &ProductController{}
}

// GotoPeru maneja la ruta para GotoPeru
func (pc *ProductController) GotoPeru(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Welcome to GotoPeru",
	})
}

// GotoTalam maneja la ruta para GotoTalam
func (pc *ProductController) GotoTalam(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Welcome to GotoTalam",
	})
}

// ListProducts lista todos los productos
func (pc *ProductController) ListProducts(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"products": []string{"GotoPeru", "GotoTalam", "MachuPicchu"},
	})
}
