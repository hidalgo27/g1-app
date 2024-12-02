package entities

type PackagePriceDTO struct {
	ID        int    `json:"id"`
	Estrellas int    `json:"estrellas"`
	PrecioS   int    `json:"precio_s"`
	PrecioD   int    `json:"precio_d"`
	PrecioT   int    `json:"precio_t"`
	CodigoS   string `json:"codigo_s"`
	CodigoD   string `json:"codigo_d"`
	CodigoT   string `json:"codigo_t"`
}
