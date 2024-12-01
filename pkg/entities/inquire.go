package entities

import "time"

type InquireDTO struct {
	ID        int       `json:"id"`
	Hotel     string    `json:"hotel"`
	Nombre    string    `json:"nombre"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
