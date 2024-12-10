package models

import "time"

type User struct {
	ID               uint           `gorm:"primaryKey"`
	Name             string         `gorm:"size:255;not null"`
	Email            string         `gorm:"size:255;unique;not null"`
	Password         string         `gorm:"size:255;not null"`
	ProfilePhotoPath string         `gorm:"size:2048"`
	Roles            []Role         `gorm:"many2many:user_roles"`
	Products         []Product      `gorm:"many2many:user_products"`
	RefreshTokens    []RefreshToken `gorm:"foreignKey:UserID"`
	Activities       []UserActivity `gorm:"foreignKey:UserID"`
	CreatedAt        time.Time      `gorm:"autoCreateTime"`
	UpdatedAt        time.Time      `gorm:"autoUpdateTime"`
}

func (User) TableName() string {
	return "users"
}
