package repository

import (
	"context"

	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
	uuid "github.com/google/uuid"
)

func (r *repository) Search(
	ctx context.Context,
	userID uuid.UUID,
	blogID uuid.UUID,
	interests []newsletter.Interest,
	limit int,
	offset int,
) ([]*newsletter.Subscription, error) {
	var subs []*newsletter.Subscription

	for _, d := range r.data {
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
