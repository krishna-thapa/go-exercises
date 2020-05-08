package handler

import (
	"newsfeeder/platform/newsfeed"
	"net/http"
	"github.com/gin-gonic/gin"
)

func NewsfeedGet(feed newsfeed.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := feed.GetAll()
		c.JSON(http.StatusOK, results)
	}
}