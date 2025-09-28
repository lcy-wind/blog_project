package user

import (
	"blog_project/dbs"
	"blog_project/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// 加密成本，默认为10，范围是4-31
var bcryptCost = 12

// 调用类
type UserController struct{}

// 用户登录
func (uc UserController) Login(c *gin.Context) {
	var userParams dbs.User
	c.ShouldBindJSON(&userParams)
	if userParams.Username == "" || userParams.Password == "" {
		c.Error(middleware.NewAppError(http.StatusInternalServerError, "LOGIN_ERROR", "登录时 用户名、密码不能为空。", nil))
		return
	}
	var user dbs.User
	dbs.DB.Where("username = ?", userParams.Username).First(&user)
	if user.Username == "" {
		c.Error(middleware.NewAppError(http.StatusInternalServerError, "LOGIN_ERROR", "账户不存在，请注册后再登录。", nil))
		return
	}
	// 校验登录密码是否正确
	// err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userParams.Password))
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userParams.Password))
	if err != nil {
		c.Error(middleware.NewAppError(http.StatusInternalServerError, "LOGIN_ERROR", "密码验证错误", nil))
		return
	}
	c.Set("username", user.Username)
	c.Set("userId", user.Id)
	jwtToken, _ := middleware.JwtUtil{}.CreateToken(c)
	c.JSON(200, gin.H{"mesg": "登录成功", "Authorization": "Bearer " + jwtToken})
}

// 用户注册
func (uc UserController) Register(c *gin.Context) {
	var userParams dbs.User = dbs.User{}
	c.ShouldBindJSON(&userParams)
	if userParams.Username == "" || userParams.Password == "" || userParams.Email == "" {
		c.Error(middleware.NewAppError(http.StatusInternalServerError, "REGISTER_ERROR", "注册时 用户名称、密码 , 邮箱 不能为空", nil))
		return
	}
	var user dbs.User
	// 首先验证登录账号是否已经存在
	dbs.DB.Where("username = ?", userParams.Username).First(&user)
	if user.Username != "" {
		c.Error(middleware.NewAppError(http.StatusInternalServerError, "REGISTER_ERROR", "账号已存在 , 请勿重复注册", nil))
		return
	} else {
		// 生成哈希（bcrypt会自动生成随机盐并混入哈希值）
		hashBytes, err := bcrypt.GenerateFromPassword([]byte(userParams.Password), bcryptCost)
		if err != nil {
			c.JSON(500, gin.H{"err": err.Error()})
		}
		user.Username = userParams.Username
		user.Password = string(hashBytes)
		user.Email = userParams.Email
		dbs.DB.Create(&user)
		c.JSON(200, gin.H{"message": "账号注册成功"})
	}
}
