package repository

import (
	"context"
	"database/sql"

	"github.com/arham09/conn-db/helpers"
	"github.com/arham09/conn-db/models"
	"github.com/arham09/conn-db/supplier"
	"github.com/sirupsen/logrus"
)

type pgSupplierRepository struct {
	Conn *sql.DB
}

// NewPgSupplierRepository will create an object that represent the supplier.Repository interface
func NewPgSupplierRepository(Conn *sql.DB) supplier.Repository {
	return &pgSupplierRepository{Conn}
}

func (p *pgSupplierRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Supplier, error) {
	rows, err := p.Conn.QueryContext(ctx, query, args...)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	result := make([]*models.Supplier, 0)

	for rows.Next() {
		t := new(models.Supplier)
		err = rows.Scan(&t.ID, &t.Code, &t.Name, &t.Address, &t.Status)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, t)
	}

	return result, nil
}

func (p *pgSupplierRepository) FetchAll(ctx context.Context) (res []*models.Supplier, err error) {
	query := `SELECT id, code, name, address, status FROM suppliers`

	res, err = p.fetch(ctx, query)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (p *pgSupplierRepository) FetchById(ctx context.Context, id int64) (res *models.Supplier, err error) {
	query := `SELECT id, code, name, address, status FROM suppliers WHERE id = $1`

	list, err := p.fetch(ctx, query, id)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return nil, helpers.ErrNotFound
	}

	return res, nil
}
