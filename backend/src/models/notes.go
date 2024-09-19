package models

import (
	"fmt"

	"github.com/GreenCodeBook/src/utility"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Id       int    `json:"id" gorm:"primary_key;autoIncrement:true" validate:"required"`
	UserId   int    ` json:"userid" validate:"required"`
	MarkUser User   `gorm:"foreignKey:UserId"`
	Text     string `json:"text" validate:"required"`
}

type NoteDto struct {
	Token string `json:"token" validate:"required"`
	Text  string `json:"text" validate:"required"`
}

func (dto *NoteDto) MapToNote() (Note, error) {

	claims, er := utility.ParseToken(dto.Token)
	if er != nil {
		fmt.Println("error in parse ", er)
		return Note{}, er
	}
	id := claims.ID
	ans := Note{
		Model:  gorm.Model{},
		Id:     0,
		UserId: id,
		Text:   dto.Text,
	}
	return ans, nil

}
