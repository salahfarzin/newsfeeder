package main

import (
	"github.com/gin-gonic/gin"
	"newsfeeder/httpd/handler"
	"newsfeeder/platform/newsfeed"
)

func main() {
	feed := newsfeed.New()

	r := gin.Default()

	r.GET("/ping", handler.PingGet)
	r.GET("/newsfeed", handler.NewFeedGet(feed))
	r.POST("/newsfeed", handler.NewFeedPost(feed))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
