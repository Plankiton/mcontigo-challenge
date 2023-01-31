package newsletter

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Search(
		ctx context.Context,
		userID uuid.UUID,
		blogID uuid.UUID,
		interests []Interest,
		limit int,
		offset int,
	) ([]*Subscription, error)
	Insert(
		ctx context.Context,
		sub Subscription,
	) error
	Count(
		ctx context.Context,
	) int
	PrintData(
		ctx context.Context,
	)
}
