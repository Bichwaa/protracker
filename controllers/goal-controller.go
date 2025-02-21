package controllers

import (
	"protracker/models"

	"gorm.io/gorm"
)

func CreateGoalController(db *gorm.DB, g models.Goal) error {
	if err := models.CreateGoal(db, &g); err != nil {
		return err
	}
	return nil
}

func GetGoalsController(db *gorm.DB) ([]models.Goal, error) {
	goals, err := models.GetGoals(db)
	if err != nil {
		return nil, err
	}
	return goals, nil
}

func GetGoalController(db *gorm.DB, id uint) (models.Goal, error) {
	goal, err := models.GetGoal(db, id)
	if err != nil {
		return models.Goal{}, err
	}
	return goal, nil
}

func UpdateGoalController(db *gorm.DB, gd models.GoalUpdateDTO) error {
	goal := new(models.Goal)
	if er := db.First(&goal, gd.ID).Error; er != nil {
		return er
	}
	goal.Description = gd.Description
	goal.Name = gd.Name
	goal.EstimatedEndDate = gd.EstimatedEndDate

	err := models.UpdateGoal(db, goal)
	if err != nil {
		return err
	}
	return nil
}

func DeleteGoalController(db *gorm.DB, id uint) error {
	err := models.DeleteGoal(db, id)
	if err != nil {
		return err
	}
	return nil
}
