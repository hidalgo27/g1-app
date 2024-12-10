package entities

import "time"

type InquireDTO struct {
	ID        uint      `json:"id"`
	Hotel     string    `json:"hotel"`
	Package   string    `json:"package"`
	Name      string    `json:"nombre"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
