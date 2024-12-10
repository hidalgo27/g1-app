package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hidalgo27/app-g1/pkg/auth"
	"github.com/hidalgo27/app-g1/pkg/dto"
	"github.com/hidalgo27/app-g1/pkg/repositories"
	"github.com/hidalgo27/app-g1/pkg/services"
	"strconv"
)

type AuthController struct{}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (a AuthController) ConfigPath(app *fiber.App) *fiber.App {
	app.Post("/login", a.Login)
	app.Post("/logout", a.Logout)
	app.Post("/refresh", a.RefreshToken)
	return app
}

func (a AuthController) Login(c *fiber.Ctx) error {
	var data dto.LoginRequest
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: dto.Error{
				Code:    strconv.Itoa(fiber.StatusBadRequest),
				Message: err.Error(),
			},
		})
	}

	token, refreshToken, user, err := services.Authenticate(data.Email, data.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse{
			Error: dto.Error{
				Code:    strconv.Itoa(fiber.StatusUnauthorized),
				Message: err.Error(),
			},
		})
	}

	return c.JSON(fiber.Map{
		"token":         token,
		"refresh_token": refreshToken,
		"user":          user,
	})

}
func (a AuthController) Logout(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: dto.Error{
				Code:    strconv.Itoa(fiber.StatusBadRequest),
				Message: "Missing token",
			},
		})
	}

	if len(tokenString) > 7 && tokenString[0:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	claims, err := auth.ValidateToken(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse{
			Error: dto.Error{
				Code:    strconv.Itoa(fiber.StatusUnauthorized),
				Message: err.Error(),
			},
		})
	}

	err = repositories.DeleteAllRefreshTokensByUser(claims.UserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: dto.Error{
				Code:    strconv.Itoa(fiber.StatusInternalServerError),
				Message: err.Error(),
			},
		})
	}
	err = services.LogUserActivity(claims.UserID, "LOGOUT")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: dto.Error{
				Code:    strconv.Itoa(fiber.StatusInternalServerError),
				Message: err.Error(),
			},
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully logged out",
	})

}
func (a AuthController) RefreshToken(c *fiber.Ctx) error {
	var data dto.RefreshTokenRequest
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: dto.Error{
				Code:    strconv.Itoa(fiber.StatusBadRequest),
				Message: err.Error(),
			},
		})
	}
	newToken, newRefreshToken, err := services.RefreshAccessToken(data.Token)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse{
			Error: dto.Error{
				Code:    strconv.Itoa(fiber.StatusUnauthorized),
				Message: err.Error(),
			},
		})
	}
	return c.JSON(fiber.Map{
		"access_token":  newToken,
		"refresh_token": newRefreshToken,
	})
}
