package usecase

import (
	"context"
	"time"

	"github.com/arham09/conn-db/faktur"
	"github.com/arham09/conn-db/models"
	"github.com/arham09/conn-db/supplier"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
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
	// mapFaktur := make(map[int64]*models.Supplier)

	// for _, sup := range data {
	// 	mapFaktur[sup.ID]
	// }
	// for _, item := range data {
	// 	res, err := s.fakturRepo.FetchAllFaktur(c, item.ID)

	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	item.Faktur = res
	// }

	// return data, nil
	g, ctx := errgroup.WithContext(c)

	mapFaktur := map[int64][]*models.Faktur{}

	for _, supplier := range data {
		mapFaktur[supplier.ID] = []*models.Faktur{}
	}

	fakturChan := make(chan []*models.Faktur)

	for supplierID := range mapFaktur {
		supplierID := supplierID

		g.Go(func() error {
			res, err := s.fakturRepo.FetchAllFaktur(ctx, supplierID)

			if err != nil {
				return err
			}

			fakturChan <- res
			return nil
		})
	}

	go func() {
		err := g.Wait()

		if err != nil {
			logrus.Error(err)
			return
		}

		close(fakturChan)
	}()

	for faktur := range fakturChan {
		if faktur != nil && len(faktur) != 0 {
			mapFaktur[faktur[0].SupplierID] = faktur
		}
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	for index, item := range data {
		if f, ok := mapFaktur[item.ID]; ok {
			data[index].Faktur = f
		}
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
