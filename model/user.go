package model

import (
	database "demo/dataBase"
	"demo/dto"
	"fmt"
	"strings"
)

type User struct {
	Username string
	Password string
	Name     string
	Id       int
}

func Login(username string, password string) bool {
	var userCount int
	db := database.DBConn()
	res := db.QueryRow("SELECT count(*) FROM users WHERE username=? and password=?", username, password)
	res.Scan(&userCount)
	defer db.Close()
	if userCount == 0 {
		return false
	}
	return true
}

func GetUsers(dto dto.UserDto) (users []User, err error) {
	db := database.DBConn()
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	defer db.Close()

	sql := "select id,username,password,`name` from users"

	if dto.SelectWord != "" {
		sql = sql + " where `name` like '%" + dto.SelectWord + "%'"
	}
	rows, err := db.Query(sql+" order by id desc limit ?,?", dto.Limit*(dto.Page-1), dto.Limit)
	//if err != nil {
	//	log.Fatal(err.Error())
	//}

	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Name)
		users = append(users, user)
	}

	defer rows.Close()
	return
}

func GetUsersCount(dto dto.UserDto) int {
	var userCount int
	db := database.DBConn()
	res := db.QueryRow("SELECT count(*) FROM users WHERE `name` like ?", "%"+dto.SelectWord+"%")
	res.Scan(&userCount)
	defer db.Close()
	return userCount
}

func UpdateUserById(userDto dto.UpdateUserDto) bool {
	db := database.DBConn()
	_, err := db.Exec("UPDATE users set username=?,password=?,`name`=? where id=?", userDto.Username, userDto.Password, userDto.Name, userDto.Id)
	if err != nil {
		return false
	} else {
		return true
	}
}

func InsertUser(userDto dto.InsertUserDto) bool {
	db := database.DBConn()
	_, err := db.Exec("insert into users(username,password,`name`) values (?,?,?)", userDto.Username, userDto.Password, userDto.Name)
	if err != nil {
		return false
	} else {
		return true
	}
}

func DeleteUsersByIds(ids []int) bool {
	db := database.DBConn()
	_, err := db.Exec("DELETE from users where id in (?)", strings.Replace(strings.Trim(fmt.Sprint(ids), "[]"), " ", ",", -1))
	if err != nil {
		return false
	} else {
		return true
	}
}
