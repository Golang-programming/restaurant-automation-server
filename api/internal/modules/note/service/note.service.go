package service

import (
	"github.co/golang-programming/restaurant/api/internal/entity"
	"github.co/golang-programming/restaurant/api/internal/modules/note/dto"
	"github.co/golang-programming/restaurant/api/internal/modules/note/repository"
)

func CreateNote(input *dto.CreateNoteInput) (*entity.Note, error) {
	note := &entity.Note{
		Description: input.Description,
		OrderItemID: input.OrderItemID,
	}

	if err := repository.CreateNote(note); err != nil {
		return nil, err
	}

	return note, nil
}

func GetNoteByID(id uint) (*entity.Note, error) {
	return repository.GetNoteByID(id)
}

func UpdateNote(id uint, input *dto.UpdateNoteInput) (*entity.Note, error) {
	note, err := repository.GetNoteByID(id)
	if err != nil {
		return nil, err
	}

	note.Description = input.Description
	note.OrderItemID = input.OrderItemID

	if err := repository.UpdateNote(note); err != nil {
		return nil, err
	}

	return note, nil
}

func DeleteNote(id uint) error {
	note, err := repository.GetNoteByID(id)
	if err != nil {
		return err
	}
	return repository.DeleteNote(note)
}

func ListNotes() ([]*entity.Note, error) {
	return repository.ListNotes()
}
