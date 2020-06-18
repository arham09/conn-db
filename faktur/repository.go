package faktur

import (
	"context"

	"github.com/arham09/conn-db/faktur/models"
)

type Repository interface {
	FetchAllFaktur(ctx context.Context, supplierID int64) ([]*models.Faktur, error)
}
