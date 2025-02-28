package models

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func CreateObjective(db *gorm.DB, objectiveData *Objective) error {
	cur := db.Create(objectiveData)
	if cur.Error != nil {
		return cur.Error
	}
	return nil
}

func GetObjectives(db *gorm.DB) ([]Objective, error) {
	var objectives []Objective
	cur := db.Preload("Notes").Preload("Goals").Find(&objectives)
	if cur.Error != nil {
		return nil, cur.Error
	}
	return objectives, nil
}

func GetObjective(db *gorm.DB, id uint) (Objective, error) {
	var objective Objective
	objective.ID = id
	cur := db.Preload("Notes").Preload("Goals").Find(&objective, id)
	if cur.Error != nil {
		if errors.Is(cur.Error, gorm.ErrRecordNotFound) {
			return Objective{}, fmt.Errorf("project with ID %d not found", id)
		}
		return Objective{}, cur.Error
	}
	fmt.Println(objective)
	return objective, nil
}

func UpdateObjective(db *gorm.DB, objectiveData *Objective) error {
	cur := db.Save(objectiveData)
	if cur.Error != nil {
		return cur.Error
	}
	return nil
}

func DeleteObjective(db *gorm.DB, id uint) error {
	var obj Objective
	obj.ID = id
	if err := db.First(&obj, id).Error; err != nil {
		return err
	}
	cur := db.Delete(&obj, id)
	if cur.Error != nil {
		return cur.Error
	}
	return nil
}
