package handler

import (
	"encoding/json"
	"fmt"
	"strconv"

	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// nolint:lll // godoc
// Get godoc
// @Summary      Read subscriptions
// @Tags         subscriptions
// @Router       /newsletter/subscriptions [get]
// @Param        page	        query  int		 true   "Result page"                                   example(1)
// @Param        maxPageSize	query  int		 true   "Max page size"                                  example(10)
// @Param        userId	        query  string	 false  "User ID"                                        example(c3d96267-9f4e-4980-95a1-b7080348d456)
// @Param        blogId	        query  string	 false  "Blog ID"                                        example(6d52cf12-84c0-4d6b-a3a1-cf6350362590)
// @Param        interests	    query  []string  false  "Interests"                                      example(["tech","sports"])
// @Produce      json
// @Success      200  {array}  handler.ResponseDoc
// nolint:gocyclo //error checking branches
func (h *handler) Get(c *gin.Context) {
	ctx := c.Request.Context()

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	maxPageSize, err := strconv.Atoi(c.Query("maxPageSize"))
	if err != nil {
		maxPageSize = 10
	}

	userID, _ := uuid.Parse(c.Query("userId"))
	blogID, _ := uuid.Parse(c.Query("blogId"))

	var interests []newsletter.Interest
	rawInterests := c.Query("interests")
	err = json.Unmarshal([]byte(rawInterests), &interests)
	if rawInterests != "" && err != nil {
		c.String(400, "interests query need to be an json array like: [\"tech\",\"sports\"]")
		return
	}

	subscriptionsResult, err := h.svc.Get(ctx, userID, blogID, interests, page, maxPageSize)
	if err != nil {
		c.String(500, fmt.Sprintf("error getting subscriptions: %v", err))
		return
	}

	var resultsDoc []*ResultsDoc
	for _, element := range subscriptionsResult.Page.Elements {
		resultsDoc = append(resultsDoc, &ResultsDoc{
			UserID:    element.UserID.String(),
			BlogID:    element.BlogID.String(),
			Interests: element.Interests,
		})
	}

	userIDStr := userID.String()
	if (uuid.UUID{}) == userID {
		userIDStr = ""
	}

	blogIDStr := blogID.String()
	if (uuid.UUID{}) == blogID {
		blogIDStr = ""
	}

	response := &ResponseDoc{
		&FilterDoc{
			UserID:    userIDStr,
			BlogID:    blogIDStr,
			Interests: interests,
		},
		&PaginationDoc{
			Page:             page,
			NumberOfPages:    subscriptionsResult.Pages,
			PaginationString: fmt.Sprintf("%d/%d", page, subscriptionsResult.Pages),
			MaxPageSize:      maxPageSize,
			TotalElements:    subscriptionsResult.Total,
		},
		resultsDoc,
	}

	c.JSON(200, response)
}
