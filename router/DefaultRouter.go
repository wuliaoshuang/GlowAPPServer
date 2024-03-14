package router

import (
	"github.com/gin-gonic/gin"
	"validatorDemo/controller"
	"validatorDemo/validate"
)

func DefaultRouter(r *gin.Engine) {
	defaultGroup := r.Group("/")
	{
		defaultGroup.GET("/", validate.UserValidate{}.Index, controller.UserController{}.Index)
		defaultGroup.POST("/", validate.UserValidate{}.Index, controller.UserController{}.Index)

	}
}
