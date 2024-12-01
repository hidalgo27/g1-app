package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hidalgo27/app-g1/pkg/dto"
	"github.com/hidalgo27/app-g1/pkg/services"
	"strconv"
)

type InquireController struct{}

func NewInquireController() *InquireController {
	return &InquireController{}
}

func (i InquireController) ConfigPath(app *fiber.App) *fiber.App {
	app.Get("/", i.HandlerGetInquire)
	return app
}
func (i InquireController) HandlerGetInquire(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		page = 1
	}
	size, err := strconv.Atoi(c.Query("size", "10"))
	if err != nil {
		page = 10
	}

	inquires, total, err := services.GetInquire(page, size)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			dto.ErrorResponse{
				Error: dto.Error{
					Code:    strconv.Itoa(fiber.StatusBadRequest),
					Message: err.Error(),
				},
			})
	}
	response := fiber.Map{
		"data":      inquires,
		"total":     total,
		"page":      page,
		"last_page": (total + int64(size) - 1) / int64(size),
	}

	return c.JSON(response)
}
