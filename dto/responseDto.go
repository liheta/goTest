package dto

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BaseReturn(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
	})
}

func BaseSuccess(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": message,
	})
}

func BaseFail(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": message,
	})
}

func PageReturn(c *gin.Context, data interface{}, total int) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "查询失败",
		"data":    data,
		"total":   total,
	})
}
