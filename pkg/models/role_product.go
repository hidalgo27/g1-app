package models

type RoleProduct struct {
	RoleID    uint `gorm:"not null"`
	ProductID uint `gorm:"not null"`
}

func (RoleProduct) TableName() string {
	return "role_products"
}
