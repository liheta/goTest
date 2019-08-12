package handlers

import (
	"demo/dto"
	"demo/model"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var userDto dto.UserDto
	c.ShouldBind(&userDto)
	users, _ := model.GetUsers(userDto)
	dto.PageReturn(c, users, model.GetUsersCount(userDto))
}

func UpdateUser(c *gin.Context) {
	var userDto dto.UpdateUserDto
	c.ShouldBind(&userDto)
	res := model.UpdateUserById(userDto)
	if res {
		dto.BaseSuccess(c, "修改成功")
	} else {
		dto.BaseFail(c, "修改失败")
	}
}

func InsertUser(c *gin.Context) {
	var userDto dto.InsertUserDto
	c.ShouldBind(&userDto)
	res := model.InsertUser(userDto)
	if res {
		dto.BaseSuccess(c, "添加成功")
	} else {
		dto.BaseFail(c, "添加失败")
	}
}

func DeleteUser(c *gin.Context) {
	var userDto dto.DeleteUserDto
	c.ShouldBind(&userDto)
	for _, value := range userDto.Ids {
		//fmt.Println(index, "\t",value)
		if value == 1 {
			dto.BaseFail(c, "超级管理员不能被删除")
			return
		}
	}

	res := model.DeleteUsersByIds(userDto.Ids)
	if res {
		dto.BaseSuccess(c, "删除成功")
	} else {
		dto.BaseFail(c, "添加失败")
	}
}
