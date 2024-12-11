package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hidalgo27/app-g1/cmd/server"
	"github.com/hidalgo27/app-g1/pkg/repositories"
)

type ProductController struct {
	productMiddleware *server.ProductMiddleware
}

func NewProductController() *ProductController {
	productRepo := &repositories.ProductRepositoryImpl{}
	return &ProductController{
		productMiddleware: server.NewProductMiddleware(productRepo),
	}
}

func (p ProductController) ConfigPath(app *fiber.App) *fiber.App {
	app.Get("/gotoperu", p.productMiddleware.Handle("gotoperu.com"), p.GotoPeru)
	app.Get("/gototalam", p.productMiddleware.Handle("gototalam.com"), p.GotoTalam)
	app.Get("/list", p.ListProducts)
	return app
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
