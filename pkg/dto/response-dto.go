package dto

import "github.com/gofiber/fiber/v2"

type SuccessResponse struct {
	Paging Paging `json:"paging,omitempty"`
}

type ErrorResponse struct {
	Error Error `json:"error,omitempty"`
}

type Error struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Details string `json:"details,omitempty"`
}

type Paging struct {
	Start    int         `json:"start,omitempty"`
	Limit    int         `json:"limit,omitempty"`
	Total    int         `json:"total,omitempty"`
	Links    fiber.Map   `json:"links,omitempty"`
	LastPage int         `json:"last_page,omitempty"`
	Data     interface{} `json:"data,omitempty"`
}
