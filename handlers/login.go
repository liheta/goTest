package handlers

import (
	"demo/model"
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func Login(c *gin.Context) {
	var user User
	c.ShouldBind(&user)
	flag := model.Login(user.Username, user.Password)
	if flag {
		c.JSON(200, gin.H{
			"status":  0,
			"message": "登陆成功",
		})
	} else {
		c.JSON(200, gin.H{
			"status":  1,
			"message": "密码或账号错误",
		})
	}
}
