package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/heptiolabs/healthcheck"
	"go.uber.org/zap"
)

func RunServer(app *fiber.App) {
	// Crea un manejador de healthcheck
	health := healthcheck.NewHandler()

	health.AddLivenessCheck("check server", func() error {
		return nil
	})

	// Implementación del rediness según corresponda. Ejemplo:
	/*
		health.AddReadinessCheck("database", func() error {
			if !database.CheckConnection() {
				zap.S().Errorw("Error al verificar la conexión a la base de datos")
			}
			return nil
		})*/

	// Registra el path /liveness y /readiness para que se puedan hacer pruebas de salud
	app.Get("/liveness", func(c *fiber.Ctx) error {
		err := adaptor.HTTPHandlerFunc(health.LiveEndpoint)(c)
		if err != nil {
			c.Status(fiber.StatusInternalServerError).SendString("Error")
		} else {
			c.Status(fiber.StatusOK).SendString("OK")
		}
		return nil
	})

	app.Get("/readiness", func(c *fiber.Ctx) error {
		err := adaptor.HTTPHandlerFunc(health.ReadyEndpoint)(c)
		if err != nil {
			c.Status(fiber.StatusInternalServerError).SendString("Error")
		} else {
			c.Status(fiber.StatusOK).SendString("OK")
		}
		return nil
	})

	// Iniciar el servidor
	err := app.Listen(":8081")
	if err != nil {
		zap.S().Errorw("Error al iniciar el servidor", "error", err)
	}
}
