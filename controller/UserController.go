package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"validatorDemo/model"
)

type UserController struct {
}

func (UserController) Index(c *gin.Context) {

	var user model.User

	fmt.Println(c.PostForm("name"))

	if err := c.ShouldBind(&user); err != nil {
		fmt.Println("dhhdhd", user)
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": user,
	})
}
