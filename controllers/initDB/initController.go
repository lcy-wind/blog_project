package initDB

import (
	"blog_project/dbs"

	"github.com/gin-gonic/gin"
)

type InitController struct {
}

func (in InitController) CreateTable(c *gin.Context) {
	dbs.InitDB()
}

func (in InitController) LoggerInit(c *gin.Context) {

}
