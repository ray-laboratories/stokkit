package stock

import (
	"context"
	"stokkit/domain/stock"
)

type GarmentWriter interface {
	Save(ctx context.Context, g stock.Garment) (stock.Garment, error)
	Delete(ctx context.Context, id int) error
}

type GarmentReader interface {
	Get(ctx context.Context, id int) (stock.Garment, error)
	GetAll(ctx context.Context) ([]stock.Garment, error)
}

type GarmentReaderWriter interface {
	GarmentReader
	GarmentWriter
}
