package models

import "time"

type TItinerario struct {
	ID          int    `gorm:"primaryKey;autoIncrement"`
	Codigo      string `gorm:"size:50"`
	Dia         int
	Titulo      string `gorm:"size:150"`
	Resumen     string `gorm:"type:longtext"`
	Descripcion string `gorm:"type:longtext"`
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	TPaquetes   []TPaqueteItinerario `gorm:"foreignKey:IDItinerario;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (TItinerario) TableName() string {
	return "titinerario"
}
