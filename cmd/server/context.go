package server

import (
	"github.com/gofiber/fiber/v2"
)

var (
	ContextKeyTransactionId = "TransactionId"
	ContextKeyApplicationId = "ApplicationId"
)

// Transaction ID
func TransactionID(ctx *fiber.Ctx) string {
	transactionID := ctx.Locals(ContextKeyTransactionId)

	return transactionID.(string)
}

// Application ID
func ApplicationID(ctx *fiber.Ctx) string {
	applicationID := ctx.Locals(ContextKeyApplicationId)

	return applicationID.(string)
}
