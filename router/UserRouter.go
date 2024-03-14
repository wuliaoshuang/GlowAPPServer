package router

import "github.com/gin-gonic/gin"

func UserRouter(r *gin.Engine) {
	userGroup := r.Group("/user")

	{
		userGroup.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "请求成功！！",
			})
		})
	}
}
