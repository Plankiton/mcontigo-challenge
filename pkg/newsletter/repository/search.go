package repository

import (
	"context"
	"fmt"

	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
	uuid "github.com/google/uuid"
)

func (r *repository) Search(
	ctx context.Context,
	userID uuid.UUID,
	blogID uuid.UUID,
	interests []newsletter.Interest,
	offset int,
	limit int,
) ([]*newsletter.Subscription, error) {
	subs := make([]*newsletter.Subscription, 0)

	fmt.Println(len(r.data), offset, limit)
	if len(r.data) < offset {
		return subs, nil
	}

	for i, d := range r.data[offset:] {
		if i == offset+limit {
			break
		}

		repoInterests := parseInterests(d.Interests)

		if d.BlogID == blogID.String() || d.UserID == userID.String() || isSourceHavingSomeOfInterest(interests, repoInterests) {
			subs = append(subs, &newsletter.Subscription{
				UserID:    uuid.MustParse(d.UserID),
				BlogID:    uuid.MustParse(d.BlogID),
				Interests: repoInterests,
			})
		}
	}

	return subs, nil
}

func (r *repository) Count(
	ctx context.Context,
) int {
	return len(r.data)
}
