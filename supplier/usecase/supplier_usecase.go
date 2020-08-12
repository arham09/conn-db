package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/arham09/conn-db/faktur"
	"github.com/arham09/conn-db/helpers/caching"
	"github.com/arham09/conn-db/models"
	"github.com/arham09/conn-db/supplier"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

type supplierUsecase struct {
	supplierRepo   supplier.Repository
	fakturRepo     faktur.Repository
	redis          caching.Caching
	contextTimeout time.Duration
}

// NewSupplierUsecase will create new an supplierUsecase object representation of supplier.Usecase interface
func NewSupplierUsecase(s supplier.Repository, f faktur.Repository, r caching.Caching, timeout time.Duration) supplier.Usecase {
	return &supplierUsecase{
		supplierRepo:   s,
		fakturRepo:     f,
		redis:          r,
		contextTimeout: timeout,
	}
}

func (s *supplierUsecase) fillFakturDetails(c context.Context, data []*models.Supplier) ([]*models.Supplier, error) {
	g, ctx := errgroup.WithContext(c)

	mapFaktur := map[int64][]models.Faktur{}

	for _, supplier := range data {
		mapFaktur[supplier.ID] = []models.Faktur{}
	}

	fakturChan := make(chan []models.Faktur)

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

	value, ok := s.redis.GetItem(c, "get:supplier:all")

	if ok {
		res := make([]*models.Supplier, 0)

		err := json.Unmarshal([]byte(value), &res)

		if err != nil {
			return nil, err
		}

		return res, nil
	}

	res, err := s.supplierRepo.FetchAll(ctx)

	if err != nil {
		return nil, err
	}

	res, err = s.fillFakturDetails(c, res)

	if err != nil {
		return nil, err
	}

	err = s.redis.SetItem(c, "get:supplier:all", res, 1*time.Hour)

	if err != nil {
		return nil, err
	}

	fmt.Println("from pg")

	return res, nil
}

func (s *supplierUsecase) FetchById(c context.Context, id int64) (*models.Supplier, error) {
	key := fmt.Sprintf("get:supplier:%d", id)
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)

	defer cancel()

	value, ok := s.redis.GetItem(c, key)

	if ok {
		var res *models.Supplier

		err := json.Unmarshal([]byte(value), &res)

		if err != nil {
			return nil, err
		}

		return res, nil
	}

	res, err := s.supplierRepo.FetchById(ctx, id)

	if err != nil {
		return nil, err
	}

	faktur, err := s.fakturRepo.FetchAllFaktur(ctx, id)

	if err != nil {
		return nil, err
	}

	res.Faktur = faktur

	err = s.redis.SetItem(c, key, res, 1*time.Hour)

	if err != nil {
		return nil, err
	}

	return res, nil
}
