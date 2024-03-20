package controllers

import "github.com/gin-gonic/gin"

type BaseController struct{}

func (u BaseController) Login(c *gin.Context) {
	SuccessResp(c, "登录成功", CodeMapping.Success, 0)
}

func (u BaseController) Register(c *gin.Context) {
	SuccessResp(c, "注册成功", CodeMapping.Success, 0)
}

func (u BaseController) Reset(c *gin.Context) {
	SuccessResp(c, "重置成功", CodeMapping.Success, 0)
}
