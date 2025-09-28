package routes

import (
	"blog_project/controllers/user"

	"github.com/gin-gonic/gin"
)

func UserRoutersInit(r *gin.Engine) {
	route := r.Group("/user")
	{
		// 注册登录接口 单独走路由 不走中间件验证token
		route.POST("/register", user.UserController{}.Register)
		route.POST("/login", user.UserController{}.Login)
	}
}
