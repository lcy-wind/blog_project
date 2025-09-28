package routes

import (
	"blog_project/controllers/initDB"

	"github.com/gin-gonic/gin"
)

func InitRoutersInit(r *gin.Engine) {
	initRoute := r.Group("/init")
	{
		initRoute.GET("/createTable", initDB.InitController{}.CreateTable)
		initRoute.GET("/loggerInit", initDB.InitController{}.LoggerInit)
	}
}
