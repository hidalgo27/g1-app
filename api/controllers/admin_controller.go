package controllers

import "github.com/gofiber/fiber/v2"

type AdminController struct{}

func NewAdminController() *AdminController {
	return &AdminController{}
}

func (ac AdminController) ConfigPath(app *fiber.App) *fiber.App {
	app.Get("/", ac.Dashboard)
	return app
}
func (ac *AdminController) Dashboard(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Welcome to the admin dashboard",
	})
}
