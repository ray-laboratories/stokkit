package stock

import "stokkit/domain/base"

type Garment struct {
	base.Entity
	GarmentValues
}

type GarmentValues struct {
	Year        int    `json:"year"`
	Style       string `json:"style"`
	Description string `json:"description"`
}
