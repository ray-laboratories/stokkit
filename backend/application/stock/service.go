package stock

import (
	"context"
	"stokkit/domain/stock"
	"strings"
)

type Service struct {
	garmentRepo GarmentReaderWriter
}

func NewService(garmentRepo GarmentReaderWriter) *Service {
	return &Service{garmentRepo: garmentRepo}
}

func (svc *Service) Save(ctx context.Context, g stock.Garment) (stock.Garment, error) {
	// Trim excess
	g.Description = strings.Trim(g.Description, " \n\t")
	return svc.garmentRepo.Save(ctx, g)
}

func (svc *Service) Get(ctx context.Context, id int) (stock.Garment, error) {
	return svc.garmentRepo.Get(ctx, id)
}

func (svc *Service) GetAll(ctx context.Context) ([]stock.Garment, error) {
	return svc.garmentRepo.GetAll(ctx)
}

func (svc *Service) Delete(ctx context.Context, id int) error {
	return svc.garmentRepo.Delete(ctx, id)
}
