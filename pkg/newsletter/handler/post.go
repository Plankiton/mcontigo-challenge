package handler

import (
	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
	"github.com/gin-gonic/gin"
)

// Get defines GET /newsletter/subscription endpoint
// @Summary      Create subscriptions
// @Tags         subscriptions
// @Router       /newsletter/subscriptions [post]
// @Param        sub body newsletter.Subscription true "Subscription"
// @Produce      json
// @Success      200
// nolint:gocyclo //error checking branches
func (h *handler) Post(c *gin.Context) {
	if c.Request.Header.Get("Content-Type") == "" {
		c.Request.Header.Add("Content-Type", "application/json")
	}

	ctx := c.Request.Context()

	var subscription newsletter.Subscription
	c.Bind(&subscription)

	h.svc.Post(ctx, subscription)
	c.String(200, "{\"success\":true}")

	h.svc.PrintData(ctx)
}
