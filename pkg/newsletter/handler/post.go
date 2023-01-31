package handler

import (
	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
	"github.com/gin-gonic/gin"
)

// Get defines GET /newsletter/subscriptions endpoint
// @Summary      Read subscriptions
// @Tags         subscriptions
// @Router       /newsletter/subscriptions [get]
// @Param        page	        query  int		 true   "Result page"                                   example(1)
// @Param        maxPageSize	query  int		 true   "Max page size"                                  example(10)
// @Param        userId	        query  string	 false  "User ID"                                        example(c3d96267-9f4e-4980-95a1-b7080348d456)
// @Param        blogId	        query  string	 false  "Blog ID"                                        example(6d52cf12-84c0-4d6b-a3a1-cf6350362590)
// @Param        interests	    query  []string  false  "Interests"                                      example(["tech","sports"])
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
