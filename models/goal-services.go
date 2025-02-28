package models

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func CreateGoal(db *gorm.DB, goalData *Goal) error {
	cur := db.Create(goalData)
	if cur.Error != nil {
		return cur.Error
	}
	return nil
}

func GetGoals(db *gorm.DB) ([]Goal, error) {
	var goals []Goal
	cur := db.Preload("Notes").Preload("Tasks").Find(&goals)
	if cur.Error != nil {
		return nil, cur.Error
	}
	return goals, nil
}

func GetGoal(db *gorm.DB, id uint) (Goal, error) {
	var goal Goal
	goal.ID = id
	cur := db.Preload("Tasks").Find(&goal, id)
	if cur.Error != nil {
		if errors.Is(cur.Error, gorm.ErrRecordNotFound) {
			return Goal{}, fmt.Errorf("goal with ID %d not found", id)
		}
		return Goal{}, cur.Error
	}
	fmt.Println(goal)
	return goal, nil
}

func UpdateGoal(db *gorm.DB, goalData *Goal) error {
	cur := db.Save(goalData)
	if cur.Error != nil {
		return cur.Error
	}
	return nil
}

func DeleteGoal(db *gorm.DB, id uint) error {
	var goal Goal
	goal.ID = id
	if err := db.First(&goal, id).Error; err != nil {
		return err
	}
	cur := db.Delete(&goal, id)
	if cur.Error != nil {
		return cur.Error
	}
	return nil
}
