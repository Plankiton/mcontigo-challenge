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
	println(len(elements))

	return &newsletter.Result[*newsletter.Subscription]{
		Total: s.repo.Count(ctx),
		Pages: s.repo.Count(ctx) / MaxPageSize,
		Page: newsletter.Page[*newsletter.Subscription]{
			Number:   Page,
			Elements: elements,
		},
	}, nil
}

func (s *service) PrintData(ctx context.Context) {
	s.repo.PrintData(ctx)
}
