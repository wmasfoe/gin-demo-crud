package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonStruct struct {
	Code  int         `json:"code"`
	Msg   interface{} `json:"msg" default:"success"`
	Data  interface{} `json:"data"`
	Count uint64      `json:"count"`
}

type CodeMappingSt struct {
	Success  int
	Error    int
	NotFound int
}

var CodeMapping = CodeMappingSt{
	Success:  200,
	Error:    0,
	NotFound: 404,
}

func SuccessResp(c *gin.Context, data interface{}, code int, count uint64, msg ...interface{}) {
	json := &JsonStruct{code, msg, data, count}
	c.JSON(http.StatusOK, json)
}

func ErrorResp(c *gin.Context, msg interface{}, code int) {
	json := &JsonStruct{Code: code, Msg: msg}
	c.JSON(0, json)
}
