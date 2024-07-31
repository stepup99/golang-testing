package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stepup99/golang-testing/user_gin/controller"
)

func main() {
	r := gin.Default()
	r.GET("/user/:user_id", controller.GetUser)
	r.Run(":9001")

}
