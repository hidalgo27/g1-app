package models

import "time"

type Permission struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"size:255;unique;not null"`
	Description string    `gorm:"type:text"`
	Roles       []Role    `gorm:"many2many:role_has_permissions"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

func (Permission) TableName() string {
	return "permissions"
}
