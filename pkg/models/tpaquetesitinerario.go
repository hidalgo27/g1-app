package models

import "time"

type TPaqueteItinerario struct {
	ID           int `gorm:"primaryKey;autoIncrement"`
	IDPaquetes   int `gorm:"column:idpaquetes;index"`   // Foreign key a TPaquetes
	IDItinerario int `gorm:"column:iditinerario;index"` // Foreign key a TItinerario
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
	TPaquete     TPaquetes   `gorm:"foreignKey:IDPaquetes;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TItinerario  TItinerario `gorm:"foreignKey:IDItinerario;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (TPaqueteItinerario) TableName() string {
	return "tpaquetesitinerario"
}