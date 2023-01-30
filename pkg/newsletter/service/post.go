package service

import (
	"context"

	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
)

func (s *service) Post(
	ctx context.Context,
	sub newsletter.Subscription,
) error {
	return s.repo.Insert(ctx, sub)
}
