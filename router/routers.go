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
	userRoute.GET("/info/:id", controllers.UserController{}.Info)
	userRoute.GET("/info", controllers.UserController{}.QueryInfo)

	baseRoute := r.Group("/")
	baseRoute.POST("/login", controllers.BaseController{}.Login)
	baseRoute.POST("/register", controllers.BaseController{}.Register)
	baseRoute.POST("/reset", controllers.BaseController{}.Reset)

	return r
}
