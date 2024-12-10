package models

import "time"

type UserActivity struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	Action    string    `gorm:"size:255;not null"`
	Timestamp time.Time `gorm:"autoCreateTime"`

	// Relación con Usuario
	User User `gorm:"foreignKey:UserID"`
}

func (UserActivity) TableName() string {
	return "user_activities"
}
