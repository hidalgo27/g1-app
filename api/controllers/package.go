package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hidalgo27/app-g1/pkg/dto"
	"github.com/hidalgo27/app-g1/pkg/services"
	"strconv"
)

type PackageController struct{}

func NewPackageController() *PackageController {
	return &PackageController{}
}

func (p PackageController) ConfigPath(app *fiber.App) *fiber.App {
	app.Get("/prices", p.HandlerGetPackagesWithPrices)
	app.Get("/itinerary", p.HandlerGetPackagesWithItineraries)
	return app
}

func (p PackageController) HandlerGetPackagesWithPrices(c *fiber.Ctx) error {
	packages, err := services.GetAllPackagesWithPrices()
	if err != nil {
		//return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		//	"error": err.Error(),
		//})
		return c.Status(fiber.StatusBadRequest).JSON(
			dto.ErrorResponse{
				Error: dto.Error{
					Code:    strconv.Itoa(fiber.StatusBadRequest),
					Message: err.Error(),
				},
			})
	}

	return c.JSON(fiber.Map{
		"data": packages,
	})
}

func (p PackageController) HandlerGetPackagesWithItineraries(c *fiber.Ctx) error {
	packages, err := services.GetAllPackagesWithItineraries()
	if err != nil {
		//return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		//	"error": err.Error(),
		//})
		return c.Status(fiber.StatusBadRequest).JSON(
			dto.ErrorResponse{
				Error: dto.Error{
					Code:    strconv.Itoa(fiber.StatusBadRequest),
					Message: err.Error(),
				},
			})
	}

	return c.JSON(fiber.Map{
		"data": packages,
	})
}
