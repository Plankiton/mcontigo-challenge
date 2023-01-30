package repository

import (
	"encoding/json"
	"fmt"
	"sync"

	newsletter "git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
	"github.com/google/uuid"
)

var (
	repo newsletter.Repository
	once sync.Once
)

type repository struct {
	data []*subscriptionDBModel
}

type Option func(*repository) error

func New(opts ...Option) (newsletter.Repository, error) {
	r := &repository{
		data: []*subscriptionDBModel{
			{
				UserID:    uuid.NewString(),
				BlogID:    uuid.NewString(),
				Interests: []string{string(newsletter.InterestPolitics)},
			},
		},
	}

	rJson, _ := json.MarshalIndent(r.data, "", "  ")
	fmt.Println(string(rJson))

	for _, opt := range opts {
		if err := opt(r); err != nil {
			return nil, err
		}
	}
	return r, nil
}

func Must() newsletter.Repository {
	once.Do(func() {
		r, err := New()
		if err != nil {
			panic(err)
		}

		repo = r
	})

	return repo
}
