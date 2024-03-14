package router

import "github.com/gin-gonic/gin"

func IndexRouter(r *gin.Engine) {
	DefaultRouter(r)
	UserRouter(r)
}
