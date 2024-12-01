package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var (
	transactionIDHeader = "Transaction-ID"
	applicationIDHeader = "Application-ID"
)

func ContextServerMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var (
			transactionID = c.Get(transactionIDHeader)
			applicationID = c.Get(applicationIDHeader)
		)

		if transactionID != "" {
			c.Locals(ContextKeyTransactionId, transactionID)
		} else {
			c.Locals(ContextKeyTransactionId, uuid.New().String())
		}

		if applicationID != "" {
			c.Locals(ContextKeyApplicationId, applicationID)
		} else {
			c.Locals(ContextKeyApplicationId, "")
		}

		return c.Next()
	}
}

func ResponseInterceptorMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()

		// Agrega el Transaction-ID al encabezado de la respuesta
		c.Set(transactionIDHeader, TransactionID(c))
		c.Set(applicationIDHeader, ApplicationID(c))

		// Devuelve el error, si lo hay
		return err
	}
}
