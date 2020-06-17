package supplier

import (
	"context"

	"github.com/arham09/conn-db/supplier/models"
)

type Repository interface {
	FetchAll(ctx context.Context) (res []*models.Supplier, err error)
}
