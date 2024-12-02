package models

import "time"

type TPrecioPaquetes struct {
	ID         int `gorm:"primaryKey;autoIncrement"`
	Estrellas  int
	PrecioS    int
	PrecioD    int
	PrecioT    int
	CodigoS    string `gorm:"size:255"`
	CodigoD    string `gorm:"size:255"`
	CodigoT    string `gorm:"size:255"`
	Estado     int
	IDPaquetes int `gorm:"column:idpaquetes;index"` // Asegúrate de que GORM use el nombre correcto de la columna
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
	TPaquete   TPaquetes `gorm:"foreignKey:IDPaquetes;references:ID"`
}

func (TPrecioPaquetes) TableName() string {
	return "tpreciopaquetes"
}
