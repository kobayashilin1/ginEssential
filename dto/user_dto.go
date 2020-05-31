package dto

import "github.com/kobayashilin1/ginEssential/model"

type UserDto struct {
	Name string `json:"name"`
	Telephone string`json:"telephone"`
}

//UserDto为结构体UserDto 的method，只返回给前端Name 和string 两个字段信息
func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name: user.Name,
		Telephone: user.Telephone,
	}
}
