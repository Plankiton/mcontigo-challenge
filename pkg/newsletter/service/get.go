package service

import (
	"context"

	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
	"github.com/google/uuid"
)

func (s *service) Get(
	ctx context.Context,
	UserID uuid.UUID,
	BlogID uuid.UUID,
	Interests []newsletter.Interest,
	Page int,
	MaxPageSize int,
) (*newsletter.Result[*newsletter.Subscription], error) {
	offset := MaxPageSize * (Page - 1)
	elements, _ := s.repo.Search(ctx, UserID, BlogID, Interests, offset, MaxPageSize)

	return &newsletter.Result[*newsletter.Subscription]{
		Total: 1,
		Pages: 1,
		Page: newsletter.Page[*newsletter.Subscription]{
			Number:   1,
			Elements: elements,
		},
	}, nil
}
