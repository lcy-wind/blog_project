package routes

import (
	"blog_project/controllers/comments"
	"blog_project/middleware"

	"github.com/gin-gonic/gin"
)

func CommentRouteInit(r *gin.Engine) {
	route := r.Group("/comment").Use(middleware.JWTAuthMiddleware())
	{
		// 发布评论接口
		route.POST("/createComment", comments.CommentController{}.CreateComment)
		// 查询评论接口
		route.POST("/selectComment", comments.CommentController{}.SelectComment)
	}
}
