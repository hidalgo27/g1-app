package models

type UserProduct struct {
	UserID    uint `gorm:"not null"`
	ProductID uint `gorm:"not null"`
}

func (UserProduct) TableName() string {
	return "user_products"
}
