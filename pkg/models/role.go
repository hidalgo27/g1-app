package models

import "time"

type Role struct {
	ID          uint         `gorm:"primaryKey"`
	Name        string       `gorm:"size:255;unique;not null"`
	Description string       `gorm:"type:text"`
	Permissions []Permission `gorm:"many2many:role_has_permissions"`
	Products    []Product    `gorm:"many2many:role_products"`
	Users       []User       `gorm:"many2many:user_roles"`
	CreatedAt   time.Time    `gorm:"autoCreateTime"`
	UpdatedAt   time.Time    `gorm:"autoUpdateTime"`
}

func (Role) TableName() string {
	return "roles"
}
