package controllers

import (
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u UserController) GetUserList(c *gin.Context) {
	SuccessResp(c, "user list", CodeMapping.Success, 0)
}

func (u UserController) AddUser(c *gin.Context) {
	SuccessResp(c, "add user", CodeMapping.Success, 0)
}

func (u UserController) DeleteUser(c *gin.Context) {
	SuccessResp(c, "delete user", CodeMapping.Success, 0)
}

func (u UserController) Info(c *gin.Context) {
	id := c.Param("id")
	SuccessResp(c, map[string]string{
		"id": id,
	}, CodeMapping.Success, 0)
}

type QueryInfo struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func (u UserController) QueryInfo(c *gin.Context) {
	param := make(map[string]interface{})
	err := c.BindJSON(&param)

	if err != nil {
		ErrorResp(c, "参数错误1", CodeMapping.Error)
		return
	}

	structQuery := QueryInfo{}
	// sErr := c.BindJSON(&structQuery)

	// if sErr != nil {
	// 	panic(sErr.Error())
	// 	ErrorResp(c, "参数错误2", CodeMapping.Error)
	// 	return
	// }

	SuccessResp(c, map[string]interface{}{
		"param":  param,
		"struct": structQuery,
	}, CodeMapping.Success, 0)
}
