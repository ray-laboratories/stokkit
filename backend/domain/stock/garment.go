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

func (g *Garment) SetKey(id int) {
	g.ID = id
}

func (g *Garment) GetKey() int {
	return g.ID
}
