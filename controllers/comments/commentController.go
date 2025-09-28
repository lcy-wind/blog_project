package comments

import (
	"blog_project/dbs"
	"blog_project/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentController struct{}
type CommentParam struct {
	Content   string `json:"content"`
	PostTitle string `json:"postTitle"`
}

func (comm CommentController) CreateComment(c *gin.Context) {
	var commentParam CommentParam
	c.ShouldBindJSON(&commentParam)
	if commentParam.Content == "" {
		c.Error(middleware.NewAppError(http.StatusInternalServerError, "COMMENT_ERROR", "评论内容不能为空", nil))
		return
	}
	if commentParam.PostTitle == "" {
		c.Error(middleware.NewAppError(http.StatusInternalServerError, "COMMENT_ERROR", "文章标题不能为空", nil))
		return
	}
	var post dbs.Post
	dbs.DB.Where("title = ?", commentParam.PostTitle).First(&post)
	if post.Id == 0 {
		c.Error(middleware.NewAppError(http.StatusInternalServerError, "COMMENT_ERROR", "文章不存在,请重新选择文章评论", nil))
		return
	}

	var comment dbs.Comment
	comment.Content = commentParam.Content
	comment.UserId = c.GetUint("userId")
	comment.PostId = post.Id
	dbs.DB.Create(&comment)
	c.JSON(200, gin.H{"message": "评论成功.."})
}
func (comm CommentController) SelectComment(c *gin.Context) {
	var commentParam CommentParam
	c.ShouldBindJSON(&commentParam)
	var comments []dbs.Comment
	if commentParam.PostTitle != "" {
		var post dbs.Post
		dbs.DB.Where("title = ?", commentParam.PostTitle).First(&post)
		if post.Id == 0 {
			c.Error(middleware.NewAppError(http.StatusInternalServerError, "COMMENT_ERROR", "该文章不存在，请选择存在的文章查询", nil))
			return
		}
		dbs.DB.Where("post_id = ?", post.Id).Find(&comments)
		c.JSON(200, gin.H{"data": comments})
		return
	}
	dbs.DB.Find(&comments)
	c.JSON(200, gin.H{"data": comments})
}
