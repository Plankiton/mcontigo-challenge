package repository

import newsletter "git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"

func parseInterests(interests []string) []newsletter.Interest {
	var parsed []newsletter.Interest
	for _, i := range interests {
		parsed = append(parsed, newsletter.Interest(i))
	}

	return parsed
}

func parseFromInterests(interests []newsletter.Interest) []string {
	var parsed []string
	for _, i := range interests {
		parsed = append(parsed, string(i))
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
