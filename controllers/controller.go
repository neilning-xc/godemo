package controllers

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendJSONResponse(c *gin.Context, code int, data interface{}) {
	response := Response{Code: code, Message: "success", Data: data}
	c.JSON(code, response)
}
