package controllers

import "github.com/gin-gonic/gin"

type JsonStruct struct {
	Code  int         `json:"code"`
	Msg   interface{} `json:"msg"`
	Data  interface{} `json:"data"`
	Count uint64      `json:"count"`
}

func SuccessModel(c *gin.Context, code int, meg interface{}, data interface{}, count uint64) {
	json := &JsonStruct{code, meg, data, count}
	c.JSON(200, json)
}

func ErrorModel(c *gin.Context, code int, meg interface{}) {
	json := &JsonStruct{Code: code, Msg: meg}
	c.JSON(200, json)
}
