package supplier

import (
	"context"

	"github.com/arham09/conn-db/models"
)

type Repository interface {
	FetchAll(ctx context.Context) (res []*models.Supplier, err error)
}
