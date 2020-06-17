package usecase

import (
	"context"
	"time"

	"github.com/arham09/conn-db/supplier"
	"github.com/arham09/conn-db/supplier/models"
)

type supplierUsecase struct {
	supplierRepo   supplier.Repository
	contextTimeout time.Duration
}

// NewSupplierUsecase will create new an supplierUsecase object representation of supplier.Usecase interface
func NewSupplierUsecase(s supplier.Repository, timeout time.Duration) supplier.Usecase {
	return &supplierUsecase{
		supplierRepo:   s,
		contextTimeout: timeout,
	}
}

func (s *supplierUsecase) FetchAll(c context.Context) ([]*models.Supplier, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)

	defer cancel()

	res, err := s.supplierRepo.FetchAll(ctx)

	if err != nil {
		return nil, err
	}

	return res, nil
}
