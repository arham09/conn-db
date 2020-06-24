package repository

import (
	"context"
	"database/sql"

	"github.com/arham09/conn-db/faktur"
	"github.com/arham09/conn-db/models"
	"github.com/sirupsen/logrus"
)

type pgFakturRepository struct {
	Conn *sql.DB
}

func NewPgFakturRepository(Conn *sql.DB) faktur.Repository {
	return &pgFakturRepository{Conn}
}

func (p *pgFakturRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]models.Faktur, error) {
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

	result := make([]models.Faktur, 0)

	for rows.Next() {
		t := new(models.Faktur)
		err = rows.Scan(&t.ID, &t.SupplierID, &t.Code, &t.ExternalID, &t.Name, &t.Status)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, *t)
	}

	return result, nil
}

func (p *pgFakturRepository) FetchAllFaktur(ctx context.Context, supplierID int64) (res []models.Faktur, err error) {
	query := `SELECT id, supplier_id, COALESCE(code, ''), COALESCE(external_id, ''), name, status FROM invoice_groups where supplier_id=$1`

	res, err = p.fetch(ctx, query, supplierID)

	if err != nil {
		return nil, err
	}

	return res, nil
}
