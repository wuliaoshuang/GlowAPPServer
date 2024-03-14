package main

import (
	"github.com/gin-gonic/gin"
	"validatorDemo/router"
)

func main() {
	r := gin.Default()

	router.IndexRouter(r)

	err := r.Run()
	if err != nil {
		return
	}

}
