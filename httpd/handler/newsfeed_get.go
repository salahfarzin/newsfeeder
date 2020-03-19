package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"newsfeeder/platform/newsfeed"
)

func NewFeedGet(feed newsfeed.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := feed.GetAll()

		c.JSON(http.StatusOK, results)
	}
}
