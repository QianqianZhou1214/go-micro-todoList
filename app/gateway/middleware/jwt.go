package middleware

import (
	"go-micro-todoList/pkg/ctl"
	"go-micro-todoList/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code uint32

		code = 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
			c.JSON(500, gin.H{
				"code": code,
				"msg":  "failed to authenticate token",
			})
		}
		claims, err := jwt.ParseToken(token)
		if err != nil {
			code = 401
			c.JSON(500, gin.H{
				"code": code,
				"msg":  "failed to authenticate token",
			})
			c.Abort()
		}
		c.Request = c.Request.WithContext(ctl.NewContext(c.Request.Context(), &ctl.UserInfo{Id: claims.Id}))

		c.Next()
	}

}
