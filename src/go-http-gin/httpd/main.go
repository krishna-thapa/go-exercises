package main

import (
	"newsfeeder/httpd/handler"
	"newsfeeder/platform/newsfeed"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World")

	feed := newsfeed.New()
	r := gin.Default()
	 
	r.GET("/ping", handler.PingGet())
	r.GET("/newsfeed", handler.NewsfeedGet(feed))
	r.POST("/newsfeed", handler.NewsfeedPost(feed))
	
	r.Run() // listen and serve on 0.0.0.0:8080
}
