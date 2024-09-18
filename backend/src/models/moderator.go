package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	id      uint   `gorm:"primary_key;autoIncrement:true" `
	Login   string `gorm:"unique"`
	HashPas string
}
