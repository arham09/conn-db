package supplier

import (
	"context"

	"github.com/arham09/conn-db/supplier/models"
)

type Usecase interface {
	FetchAll(ctx context.Context) ([]*models.Supplier, error)
}
