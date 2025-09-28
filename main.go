package main

import (
	"blog_project/middleware"
	"blog_project/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.ErrorHandlerMiddleware())
	routes.InitRoutersInit(r)
	routes.UserRoutersInit(r)
	routes.PostRouteInit(r)
	routes.CommentRouteInit(r)
	r.Run()
}
