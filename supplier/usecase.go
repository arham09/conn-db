package supplier

import (
	"context"

	"github.com/arham09/conn-db/models"
)

type Usecase interface {
	FetchAll(ctx context.Context) ([]*models.Supplier, error)
}
