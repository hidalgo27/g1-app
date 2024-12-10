package controllers

import "github.com/gofiber/fiber/v2"

type EditorController struct{}

func NewEditorController() *EditorController {
	return &EditorController{}
}

func (ec EditorController) ConfigPath(app *fiber.App) *fiber.App {
	app.Post("/", ec.EditContent)
	return app
}
func (ec *EditorController) EditContent(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Content updated successfully",
	})
}
