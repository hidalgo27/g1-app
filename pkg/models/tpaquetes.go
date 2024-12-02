package models

import "time"

type TPaquetes struct {
	ID                  int    `gorm:"primaryKey;autoIncrement"`
	Codigo              string `gorm:"size:50"`
	CodigoF             string `gorm:"size:50"`
	Titulo              string `gorm:"size:255"`
	PrecioTours         int    `gorm:"not null"`
	URL                 string `gorm:"size:255"`
	Duracion            int
	Altitud             string `gorm:"size:100"`
	GroupSize           string `gorm:"size:100"`
	Descripcion         string `gorm:"type:longtext"`
	Incluye             string `gorm:"type:text"`
	NoIncluye           string `gorm:"type:text"`
	Opcional            string `gorm:"type:longtext"`
	Imagen              string `gorm:"size:255"`
	Fecha               time.Time
	Estado              int
	OffersHome          int
	Descuento           int
	EstadoSlider        int
	IsPaquete           int
	IsTours             int
	IsPT                int
	Maps                string `gorm:"type:longtext"`
	CreatedAt           *time.Time
	UpdatedAt           *time.Time
	Mapa                string               `gorm:"size:255"`
	TPrecioPaquetes     []TPrecioPaquetes    `gorm:"foreignKey:IDPaquetes;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TPaqueteItinerarios []TPaqueteItinerario `gorm:"foreignKey:IDPaquetes;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (TPaquetes) TableName() string {
	return "tpaquetes"
}
