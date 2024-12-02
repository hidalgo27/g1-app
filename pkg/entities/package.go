package entities

type PackageDTO struct {
	ID          int               `json:"id"`
	Codigo      string            `json:"codigo"`
	Titulo      string            `json:"titulo"`
	PrecioTours int               `json:"precio_tours"`
	Url         string            `json:"url"`
	Prices      []PackagePriceDTO `json:"prices"`
}
type PackageWithItinerariesDTO struct {
	ID          int            `json:"id"`
	Codigo      string         `json:"codigo"`
	Titulo      string         `json:"titulo"`
	PrecioTours int            `json:"precio_tours"`
	Url         string         `json:"url"`
	Itineraries []ItineraryDTO `json:"itineraries"`
}
