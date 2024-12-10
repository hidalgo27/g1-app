package models

type RolePermission struct {
	RoleID       uint `gorm:"not null"`
	PermissionID uint `gorm:"not null"`
}

func (RolePermission) TableName() string {
	return "role_has_permissions"
}
