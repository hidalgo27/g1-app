package models

import "time"

type Product struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"size:255;unique;not null"`
	Description string    `gorm:"type:text"`
	Users       []User    `gorm:"many2many:user_products"`
	Roles       []Role    `gorm:"many2many:role_products"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

func (Product) TableName() string {
	return "products"
}
