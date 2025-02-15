package models

import "gorm.io/gorm"

func CreateObjective(db *gorm.DB, objectiveData *Objective) error {
	cur := db.Create(objectiveData)
	if cur.Error != nil {
		return cur.Error
	}
	return nil
}

func GetObjective(db *gorm.DB, id uint) (*Objective, error) {
	var objective Objective
	cur := db.First(&objective, id)
	if cur.Error != nil {
		return nil, cur.Error
	}
	return &objective, nil
}

func UpdateObjective(db *gorm.DB, objectiveData *Objective) error {
	cur := db.Save(objectiveData)
	if cur.Error != nil {
		return cur.Error
	}
	return nil
}

func DeleteObjective(db *gorm.DB, id uint) error {
	cur := db.Delete(&Objective{}, id)
	if cur.Error != nil {
		return cur.Error
	}
	return nil
}
