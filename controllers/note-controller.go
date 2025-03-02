package controllers

import (
	"protracker/models"

	"gorm.io/gorm"
)

func CreateNoteController(db *gorm.DB, n models.Note) error {
	if err := models.CreateNote(db, &n); err != nil {
		return err
	}
	return nil
}

func GetNotesController(db *gorm.DB) ([]models.Note, error) {
	notes, err := models.GetNotes(db)
	if err != nil {
		return nil, err
	}
	return notes, nil
}

func GetNoteController(db *gorm.DB, id uint) (models.Note, error) {
	note, err := models.GetNote(db, id)
	if err != nil {
		return models.Note{}, err
	}
	return note, nil
}

func UpdateNoteController(db *gorm.DB, nd models.NoteUpdateDTO) error {
	note := new(models.Note)
	if er := db.First(&note, nd.ID).Error; er != nil {
		return er
	}
	note.Title = nd.Title
	note.Content = nd.Content
	note.Tags = nd.Tags

	err := models.UpdateNote(db, note)
	if err != nil {
		return err
	}
	return nil
}

func DeleteNoteController(db *gorm.DB, id uint) error {
	err := models.DeleteNote(db, id)
	if err != nil {
		return err
	}
	return nil
}
