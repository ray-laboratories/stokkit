package memcache

import (
	"context"
	"stokkit/domain/stock"
)

type GarmentMemCache struct {
	memCache *MemCache[int, stock.Garment]
}

func NewGarmentMemCache() *GarmentMemCache {
	return &GarmentMemCache{
		memCache: NewMemCache[int, stock.Garment](),
	}
}

func (g GarmentMemCache) Get(ctx context.Context, id int) (stock.Garment, error) {
	return g.memCache.Get(ctx, id)
}

func (g GarmentMemCache) GetAll(ctx context.Context) ([]stock.Garment, error) {
	return g.memCache.GetAll(ctx)
}

func (g GarmentMemCache) Save(ctx context.Context, garment stock.Garment) (stock.Garment, error) {
	if garment.ID == 0 {
		newID, err := g.memCache.Create(ctx, garment)
		if err != nil {
			return stock.Garment{}, err
		}
		garment.ID = newID
	}
	err := g.memCache.Update(ctx, garment.ID, garment)
	if err != nil {
		return stock.Garment{}, err
	}
	return garment, nil
}

func (g GarmentMemCache) Delete(ctx context.Context, id int) error {
	return g.memCache.Delete(ctx, id)
}
