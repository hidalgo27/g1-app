package config

import "github.com/gofiber/fiber/v2"

type MicroServicio interface {
	ConfigPath(app *fiber.App) *fiber.App
}

// Use
func Use(prefix string, r fiber.Router, micro MicroServicio) {
	r.Mount(prefix, micro.ConfigPath(fiber.New()))
}
