package controllers

import "github.com/gin-gonic/gin"

type UserController struct{}

func (u UserController) GetUserList(c *gin.Context) {
	SuccessModel(c, 200, "success", "user list", 0)
}

func (u UserController) AddUser(c *gin.Context) {
	SuccessModel(c, 200, "success", "add user", 0)
}

func (u UserController) DeleteUser(c *gin.Context) {
	SuccessModel(c, 200, "success", "delete user", 0)
}
