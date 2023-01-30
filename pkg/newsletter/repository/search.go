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

func parseInterests(interests []string) []newsletter.Interest {
	var parsed []newsletter.Interest
	for _, i := range interests {
		parsed = append(parsed, newsletter.Interest(i))
	}

	return parsed
}

// isSourceHavingSomeOfInterest searchs on sourceInterests slice elements from searchInterestQuery and if one of then
// are on sourceInterests it returns true else false
func isSourceHavingSomeOfInterest(searchInterestQuery []newsletter.Interest, sourceInterests []newsletter.Interest) bool {
	for _, q := range searchInterestQuery {
		for _, i := range sourceInterests {
			if q == i {
				return true
			}
		}
	}

	return false
}
