package supplier

import (
	"context"

	"github.com/arham09/conn-db/models"
)

type Usecase interface {
	FetchAll(ctx context.Context) ([]*models.Supplier, error)
	FetchById(c context.Context, id int64) (*models.Supplier, error)
}
