package routes

import (
	"blog_project/controllers/post"
	"blog_project/middleware"

	"github.com/gin-gonic/gin"
)

func PostRouteInit(r *gin.Engine) {
	route := r.Group("/post").Use(middleware.JWTAuthMiddleware())
	{
		// 发布文章接口
		route.POST("/createPost", post.PostController{}.CreatePost)
		// 查询文章接口
		route.POST("/selectPost", post.PostController{}.SelectPost)
		// 更新文章接口
		route.POST("/updatePost", post.PostController{}.UpdatePost)
		// 删除文章接口
		route.POST("/deletePost", post.PostController{}.DeletePost)

		// 删除文章接口
		route.POST("/testMiddleware", post.PostController{}.TestMiddleware)
	}
}
