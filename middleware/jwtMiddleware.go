package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type JwtUtil struct{}
type Claims struct {
	Username string `json:"username"`
	UserId   uint   `json:"user_id"`
	jwt.RegisteredClaims
}

var jwtSecret = []byte("blog-lcy-1234567890")

// 生成token
func (jwtUtil JwtUtil) CreateToken(c *gin.Context) (string, error) {
	name, ok := c.Get("username")
	userId, _ := c.Get("userId")
	if ok {
		claims := &Claims{
			Username: name.(string),
			UserId:   userId.(uint),
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(8 * time.Hour)), // 过期时间
				Issuer:    "blog-lcy",                                        // 签发人
			},
		}
		// 创建新的token对象，该方法需要一个Claims的实现和签名方法
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		return token.SignedString(jwtSecret)
	}
	return "", nil
}

// 验证token
func validateToken(tokenString string) (*Claims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// 验证令牌
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}

// JWT认证中间件
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从Authorization头中获取令牌
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "认证令牌缺失，请登录后重试"})
			c.Abort()
			return
		}

		// 检查令牌格式（Bearer <token>）
		var tokenString string
		_, err := fmt.Sscanf(authHeader, "Bearer %s", &tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "令牌格式错误，应为Bearer <token>"})
			c.Abort()
			return
		}

		// 验证令牌
		claims, err := validateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "登录失效，请重新登录"})
			c.Abort()
			return
		}

		// 将用户信息存储在上下文
		c.Set("username", claims.Username)
		c.Set("userId", claims.UserId)
		c.Next()
	}
}
