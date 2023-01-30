package repository

import (
	"context"

	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
)

func (r *repository) Insert(
	ctx context.Context,
	sub newsletter.Subscription,
) error {

	r.data = append(r.data, &subscriptionDBModel{
		UserID:    sub.UserID.String(),
		BlogID:    sub.BlogID.String(),
		Interests: parseFromInterests(sub.Interests),
	})
	return nil
}
