package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id          uint   `json:"id" gorm:"primary_key;autoIncrement:true" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" gorm:"unique" validate:"required"`
	HashPas     string `json:"pass" validate:"required"`
	PhoneNumber string `json:"phone" validate:"required"`
	Role        string `json:"role" validate:"required"`
}
type UserDto struct {
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" gorm:"unique" validate:"required"`
	Password    string `json:"pass" validate:"required"`
	PhoneNumber string `json:"phone" validate:"required"`
}

type LoginUserDto struct {
	Email    string `json:"email" gorm:"unique" validate:"required"`
	Password string `json:"pass" validate:"required"`
}

func (u UserDto) MapNewUserDtoToUser() User {
	var ans User
	ans.Model = gorm.Model{}
	ans.Email = u.Email
	ans.Name = u.Name
	ans.PhoneNumber = u.PhoneNumber
	ans.HashPas = u.Password
	ans.Role = "user"
	return ans
}

func (u LoginUserDto) MapToUser() User {
	var ans User
	ans.Model = gorm.Model{}
	ans.Email = u.Email
	ans.Name = ""
	ans.PhoneNumber = ""
	ans.HashPas = u.Password
	ans.Role = ""
	return ans
}
