package handler

import (
	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
)

type handler struct {
	svc newsletter.Service
}

// @contact.name                Grupo MContigo
// @title                       Newsletter API
// @version                     1.0
// @description                 Newsletter API
func New(opts ...Option) (newsletter.Handler, error) {
	h := &handler{}

	for _, opt := range opts {
		if err := opt(h); err != nil {
			return nil, err
		}
	}

	return h, nil
}

func Must(
	svc newsletter.Service,
) newsletter.Handler {
	h, err := New(WithService(svc))
	if err != nil {
		panic(err)
	}

	return h
}
