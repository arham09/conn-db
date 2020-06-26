package supplier

import (
	"context"

	"github.com/arham09/conn-db/models"
)

type Repository interface {
	FetchAll(ctx context.Context) (res []*models.Supplier, err error)
	FetchById(ctx context.Context, id int64) (res *models.Supplier, err error)
}
