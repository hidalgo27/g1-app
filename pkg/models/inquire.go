package models

import "time"

type Inquire struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Product      string    `gorm:"size:100" json:"product"`
	Package      string    `gorm:"size:100" json:"package"`
	Hotel        string    `gorm:"size:250" json:"hotel"`
	Destinations string    `gorm:"size:250" json:"destinations"`
	Passengers   string    `gorm:"size:250" json:"passengers"`
	Duration     string    `gorm:"size:250" json:"duration"`
	Name         string    `gorm:"size:250" json:"name"`
	Email        string    `gorm:"size:250" json:"email"`
	TravelDate   time.Time `json:"travel_date"`
	TravelDate2  time.Time `json:"travel_date2"`
	Phone        string    `gorm:"size:250" json:"phone"`
	Comment      string    `gorm:"size:250" json:"comment"`
	CountryCode  string    `gorm:"size:100" json:"country_code"`
	Device       string    `gorm:"size:100" json:"device"`
	Origin       string    `gorm:"size:100" json:"origin"`
	Browser      string    `gorm:"size:100" json:"browser"`
	InitialPrice int       `json:"initial_price"`
	SalePrice    int       `json:"sale_price"`
	SubProfit    int       `json:"sub_profit"`
	Profit       int       `json:"profit"`
	Seller       int       `json:"seller"`
	SaleDate     time.Time `json:"sale_date"`
	Sent         int       `json:"sent"`
	Status       int       `json:"status"`
	InquiryDate  time.Time `json:"inquiry_date"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (Inquire) TableName() string {
	return "inquire"
}
