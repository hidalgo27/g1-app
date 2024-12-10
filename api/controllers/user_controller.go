package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hidalgo27/app-g1/pkg/dto"
	"github.com/hidalgo27/app-g1/pkg/services"
	"strconv"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (u UserController) ConfigPath(app *fiber.App) *fiber.App {
	app.Post("/register", u.CreateUser)
	return app
}
func (u UserController) CreateUser(c *fiber.Ctx) error {
	var userDTO dto.UserCreateDTO
	if err := c.BodyParser(&userDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: dto.Error{
				Code:    strconv.Itoa(fiber.StatusBadRequest),
				Message: err.Error(),
			},
		})
	}

	user, err := services.CreateUser(userDTO)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: dto.Error{
				Code:    strconv.Itoa(fiber.StatusBadRequest),
				Message: err.Error(),
			},
		})
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}
