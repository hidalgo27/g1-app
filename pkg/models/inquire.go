package models

import "time"

type Inquire struct {
	ID            int       `json:"id"`
	Producto      string    `json:"producto"`
	Package       string    `json:"package"`
	Hotel         string    `json:"hotel"`
	Destinos      string    `json:"destinos"`
	Pasajeros     string    `json:"pasajeros"`
	Duracion      string    `json:"duracion"`
	Nombre        string    `json:"nombre"`
	Email         string    `json:"email"`
	TravelDate    time.Time `json:"travel_date"`
	Telefono      string    `json:"telefono"`
	Comentario    string    `json:"comentario"`
	CodigoPais    string    `json:"codigo_pais"`
	Device        string    `json:"device"`
	Origen        string    `json:"origen"`
	Browser       string    `json:"browser"`
	PrecioInicial int       `json:"precio_inicial"`
	PrecioVenta   int       `json:"precio_venta"`
	SubProfit     int       `json:"sub_profit"`
	Profit        int       `json:"profit"`
	Vendedor      int       `json:"vendedor"`
	SaleDate      time.Time `json:"sale_date"`
	Sent          int       `json:"sent"`
	Estado        int       `json:"estado"`
	InquireDate   time.Time `json:"inquire_date"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (Inquire) TableName() string {
	return "tinquire"
}

type InquireResponse struct {
	ID     int    `json:"id"`
	Hotel  string `json:"hotel"`
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
}
