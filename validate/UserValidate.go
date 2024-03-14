package validate

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"validatorDemo/model"
)

type UserValidate struct {
}

var validate = validator.New()

func (UserValidate) Index(c *gin.Context) {

	user := model.User{}

	err := c.ShouldBind(&user)
	if err != nil {
		return
	}

	err2 := validate.Struct(user)
	fmt.Println(user)

	if err2 != nil {
		fmt.Println(err2)

		_, ok := err2.(*validator.ValidationErrors)

		if !ok {
			c.JSON(422, gin.H{
				"error": err2.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 1000,
				"msg":  "param is error",
			})
		}

		return
	}
	c.Next()
}
