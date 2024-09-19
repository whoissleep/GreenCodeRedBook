package services

import (
	"fmt"

	"github.com/GreenCodeBook/src/models"
	"gorm.io/gorm"
)

type noteService struct {
	db *gorm.DB
}

type NoteService interface {
	AddNote(user models.Note) error
	GetAllNotes(id int) ([]models.Note, error)
}

func (r noteService) AddNote(note models.Note) error {
	erdb := r.db.Create(&note)
	if erdb.Error != nil {
		fmt.Println(erdb)
		return erdb.Error
	}
	return nil
}

func (r noteService) GetAllNotes(id int) (ans []models.Note, er error) {
	er = nil
	erdb := r.db.Where("user_id = ?", id).Find(&ans)
	er = erdb.Error
	return ans, er
}

func NewNoteSevice(db *gorm.DB) NoteService {
	return &noteService{db: db}
}
