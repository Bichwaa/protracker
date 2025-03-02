package models

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func CreateNote(db *gorm.DB, noteData *Note) error {
	cur := db.Create(noteData)
	if cur.Error != nil {
		return cur.Error
	}
	return nil
}

func GetNotes(db *gorm.DB) ([]Note, error) {
	var notes []Note
	cur := db.Preload("Notes").Find(&notes)
	if cur.Error != nil {
		return nil, cur.Error
	}
	return notes, nil
}

func GetNote(db *gorm.DB, id uint) (Note, error) {
	var note Note
	note.ID = id
	cur := db.Find(&note, id)
	if cur.Error != nil {
		if errors.Is(cur.Error, gorm.ErrRecordNotFound) {
			return Note{}, fmt.Errorf("note with ID %d not found", id)
		}
		return Note{}, cur.Error
	}
	return note, nil
}

func UpdateNote(db *gorm.DB, noteData *Note) error {
	cur := db.Save(noteData)
	if cur.Error != nil {
		return cur.Error
	}
	return nil
}

func DeleteNote(db *gorm.DB, id uint) error {
	var note Note
	note.ID = id
	cur := db.Delete(&note, id)
	if cur.Error != nil {
		return cur.Error
	}
	return nil
}
