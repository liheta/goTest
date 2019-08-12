package dto

type UserDto struct {
	Page       int    `json:"page" form:"page"`
	Limit      int    `json:"limit" form:"limit"`
	SelectWord string `json:"select_word" form:"select_word"`
}

type UpdateUserDto struct {
	Id       int    `json:"Id" form:"Id"`
	Password string `json:"Password" form:"Password"`
	Name     string `json:"Name" form:"Name"`
	Username string `json:"Username" form:"Username"`
}

type InsertUserDto struct {
	Password string `json:"Password" form:"Password"`
	Name     string `json:"Name" form:"Name"`
	Username string `json:"Username" form:"Username"`
}

type DeleteUserDto struct {
	Ids []int `json:"ids" form:"ids"`
}
