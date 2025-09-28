package post

import (
	"blog_project/dbs"
	"blog_project/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostController struct{}

// 创建文章接口
func (p PostController) CreatePost(c *gin.Context) {
	var postParam dbs.Post
	c.ShouldBindJSON(&postParam)
	// 校验文章标题、内容是否为空
	if postParam.Title == "" || postParam.Content == "" {
		c.Error(middleware.NewAppError(http.StatusInternalServerError, "POST_ERROR", "文章标题、内容不能为空", nil))
		return
	}
	// 校验文章标题是否重复
	var post dbs.Post
	dbs.DB.Where("title = ?", postParam.Title).First(&post)
	if post.Title != "" {
		c.Error(middleware.NewAppError(http.StatusInternalServerError, "POST_ERROR", "文章标题已存在,请更换其他标题再发布。", nil))
		return
	}
	postParam.UserId = c.GetUint("userId")
	dbs.DB.Create(&postParam)
	c.JSON(200, "文章创建成功")
}

// 查询文章接口
func (p PostController) SelectPost(c *gin.Context) {
	var postParam dbs.Post
	c.ShouldBindJSON(&postParam)
	var postList []dbs.Post
	db := dbs.DB.Model(&postList)

	if postParam.Title != "" {
		db = db.Where("title like ?", "%"+postParam.Title+"%")
	}
	if postParam.Content != "" {
		db = db.Where("content like ?", "%"+postParam.Content+"%")
	}
	db.Order("updated_at desc").Find(&postList)
	c.JSON(200, gin.H{"data": postList})
}

// 更新文章接口
func (p PostController) UpdatePost(c *gin.Context) {
	var postParam dbs.Post
	c.ShouldBindJSON(&postParam)
	// 校验文章标题、内容是否为空
	if postParam.Id == 0 {
		c.Error(middleware.NewAppError(http.StatusInternalServerError, "POST_ERROR", "文章Id不能为空", nil))
		return
	}
	if postParam.Title == "" && postParam.Content == "" {
		c.Error(middleware.NewAppError(http.StatusInternalServerError, "POST_ERROR", "文章标题、内容 至少修改一项", nil))
		return
	}
	var post dbs.Post
	// 校验文章标题是否重复
	if postParam.Title != "" {
		dbs.DB.Where("title = ? AND Id != ?", postParam.Title, postParam.Id).First(&post)
		if post.Title != "" {
			c.Error(middleware.NewAppError(http.StatusInternalServerError, "POST_ERROR", "文章标题已存在,请重新填写。", nil))
			return
		}
	}
	dbs.DB.Where("id = ?", postParam.Id).First(&post)
	if post.Id == 0 {
		c.Error(middleware.NewAppError(http.StatusInternalServerError, "POST_ERROR", "文章不存在", nil))
		return
	}

	if post.UserId != c.GetUint("userId") {
		c.Error(middleware.NewAppError(http.StatusInternalServerError, "POST_ERROR", "你没有权限修改该文章,只有文章创建者才能更改该文章", nil))
		return
	}
	dbs.DB.Debug().Model(&dbs.Post{}).Where("id = ?", postParam.Id).Updates(map[string]interface{}{
		"title":   postParam.Title,
		"content": postParam.Content,
	})
	c.JSON(200, "文章更新成功")
}

// 删除文章接口
func (p PostController) DeletePost(c *gin.Context) {
	var postParam dbs.Post
	c.ShouldBindJSON(&postParam)
	var post dbs.Post
	dbs.DB.Debug().Where("id = ?", postParam.Id).First(&post)
	if post.Id == 0 {
		c.Error(middleware.NewAppError(http.StatusInternalServerError, "POST_ERROR", "文章不存在", nil))
		return
	}
	if post.UserId != c.GetUint("userId") {
		c.Error(middleware.NewAppError(http.StatusInternalServerError, "POST_ERROR", "你没有权限修改该文章,只有文章创建者才能更改该文章", nil))
		return
	}
	dbs.DB.Debug().Where("id = ?", postParam.Id).Delete(&post)
	c.JSON(200, "文章删除成功")
}

// 测试接口
func (uc PostController) TestMiddleware(c *gin.Context) {
	c.JSON(200, "验证成功")
}
