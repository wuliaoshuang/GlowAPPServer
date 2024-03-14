package model

type User struct {
	Name  string `json:"name" form:"name" validate:"required,max=12,min=5"`
	Age   int    `json:"age" form:"age" validate:"required,gte=0,lte=130"`
	Email string `json:"email" form:"email" validate:"required,email"`
}
