package entities

type ItineraryDTO struct {
	ID          int    `json:"id"`
	Codigo      string `json:"codigo"`
	Dia         int    `json:"dia"`
	Titulo      string `json:"titulo"`
	Resumen     string `json:"resumen"`
	Descripcion string `json:"descripcion"`
}
