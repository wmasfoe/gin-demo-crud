package router

import (
	"gin-demo/controllers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	userRoute := r.Group("/user")

	userRoute.GET("/list", controllers.UserController{}.GetUserList)

	userRoute.POST("/add", controllers.UserController{}.AddUser)

	userRoute.DELETE("/delete", controllers.UserController{}.DeleteUser)

	return r
}
