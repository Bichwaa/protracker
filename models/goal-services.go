package models

import "gorm.io/gorm"

func CreateGoal(db *gorm.DB, goalData *Goal) error {
	cur := db.Create(goalData)
	if cur.Error != nil {
		return cur.Error
	}
	return nil
}

func GetGoal(db *gorm.DB, id uint) (*Goal, error) {
	var goal Goal
	if err := db.First(&goal, id).Error; err != nil {
		return nil, err
	}
	return &goal, nil
}

func UpdateGoal(db *gorm.DB, goalData *Goal) error {
	cur := db.Save(goalData)
	if cur.Error != nil {
		return cur.Error
	}
	return nil
}

func DeleteGoal(db *gorm.DB, id uint) error {
	cur := db.Delete(&Goal{}, id)
	if cur.Error != nil {
		return cur.Error
	}
	return nil
}
