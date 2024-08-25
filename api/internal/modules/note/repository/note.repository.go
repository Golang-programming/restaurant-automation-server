package repository

import (
	"github.co/golang-programming/restaurant/api/internal/entity"
	"github.co/golang-programming/restaurant/api/pkg/database"
)

func CreateNote(note *entity.Note) error {
	return database.ActiveDB.Create(note).Error
}

func GetNoteByID(id uint) (*entity.Note, error) {
	var note entity.Note
	if err := database.ActiveDB.First(&note, id).Error; err != nil {
		return nil, err
	}
	return &note, nil
}

func UpdateNote(note *entity.Note) error {
	return database.ActiveDB.Save(note).Error
}

func DeleteNote(note *entity.Note) error {
	return database.ActiveDB.Delete(note).Error
}

func ListNotes() ([]*entity.Note, error) {
	var notes []*entity.Note
	if err := database.ActiveDB.Find(&notes).Error; err != nil {
		return nil, err
	}
	return notes, nil
}
