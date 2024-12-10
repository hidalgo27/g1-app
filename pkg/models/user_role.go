package models

type UserRole struct {
	UserID uint `gorm:"not null"`
	RoleID uint `gorm:"not null"`
}

func (UserRole) TableName() string {
	return "user_roles"
}
