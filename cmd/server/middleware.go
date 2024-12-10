package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/hidalgo27/app-g1/pkg/auth"
)

var (
	transactionIDHeader = "Transaction-ID"
	applicationIDHeader = "Application-ID"
	ContextKeyUserID    = "UserID" // Contexto para almacenar UserID
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

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Obtener el token del encabezado Authorization
		tokenString := c.Get("Authorization")
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing token"})
		}

		// Validar el token
		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		// Guardar UserID, Roles y Permisos en el contexto
		c.Locals("UserID", claims.UserID)
		c.Locals("Roles", claims.Roles)
		c.Locals("Permissions", claims.Permissions)

		return c.Next()
	}
}

//func JWTMiddleware() fiber.Handler {
//	return func(c *fiber.Ctx) error {
//		// Obtener el token JWT de la cabecera Authorization
//		tokenString, err := getTokenFromHeader(c)
//		if tokenString == "" {
//			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing token"})
//		}
//
//		// Eliminar prefijo "Bearer " si existe
//		if strings.HasPrefix(tokenString, "Bearer ") {
//			tokenString = strings.TrimPrefix(tokenString, "Bearer ")
//		}
//
//		// Validar el token JWT
//		claims, err := auth.ValidateToken(tokenString)
//		if err != nil {
//			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
//		}
//
//		// Almacenar UserID en el contexto
//		c.Locals(ContextKeyUserID, claims.UserID)
//
//		return c.Next()
//	}
//}
//
//func getTokenFromHeader(c *fiber.Ctx) (string, error) {
//	token := c.Get("Authorization")
//	if token == "" {
//		return "", fiber.NewError(fiber.StatusUnauthorized, "Missing token")
//	}
//
//	// Eliminar prefijo "Bearer " si existe
//	if strings.HasPrefix(token, "Bearer ") {
//		token = strings.TrimPrefix(token, "Bearer ")
//	}
//
//	if token == "" {
//		return "", fiber.NewError(fiber.StatusUnauthorized, "Invalid token format")
//	}
//
//	return token, nil
//}

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
