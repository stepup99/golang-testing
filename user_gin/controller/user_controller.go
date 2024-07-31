package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/stepup99/golang-testing/user_gin/service"
)

var UserService service.UserService

func init() {
	// Initialize the service and controller
	UserService = service.NewUserService()
}

func GetUser(c *gin.Context) {
	userIdStr := c.Param("user_id")
	user_id, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	user, err := UserService.GetUser(user_id)
	fmt.Println("inside controller")
	if err != nil {
		fmt.Println("inside ctrl 29", err)
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("error - %v", err)})
		return
	}
	c.JSON(http.StatusOK, user)
}
