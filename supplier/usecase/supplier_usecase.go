package usecase

import (
	"context"
	"time"

	"github.com/arham09/conn-db/faktur"
	"github.com/arham09/conn-db/supplier"
	"github.com/arham09/conn-db/supplier/models"
)

type supplierUsecase struct {
	supplierRepo   supplier.Repository
	fakturRepo     faktur.Repository
	contextTimeout time.Duration
}

// NewSupplierUsecase will create new an supplierUsecase object representation of supplier.Usecase interface
func NewSupplierUsecase(s supplier.Repository, f faktur.Repository, timeout time.Duration) supplier.Usecase {
	return &supplierUsecase{
		supplierRepo:   s,
		fakturRepo:     f,
		contextTimeout: timeout,
	}
}

func (s *supplierUsecase) fillFakturDetails(c context.Context, data []*models.Supplier) ([]*models.Supplier, error) {
	for _, item := range data {
		res, err := s.fakturRepo.FetchAllFaktur(c, item.ID)

		if err != nil {
			return nil, err
		}

		item.Faktur = res
	}

	return data, nil
}

func (s *supplierUsecase) FetchAll(c context.Context) ([]*models.Supplier, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)

	defer cancel()

	res, err := s.supplierRepo.FetchAll(ctx)

	if err != nil {
		return nil, err
	}

	res, err = s.fillFakturDetails(c, res)

	if err != nil {
		return nil, err
	}

	return res, nil
}
