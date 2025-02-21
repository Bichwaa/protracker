package controllers

import (
	"protracker/models"

	// "github.com/gofiber/fiber"
	"gorm.io/gorm"
)

func CreateObjectiveController(db *gorm.DB, p models.Objective) error {
	if err := models.CreateObjective(db, &p); err != nil {
		return err
	}
	return nil
}

func GetObjectivesController(db *gorm.DB) ([]models.Objective, error) {
	projects, err := models.GetObjectives(db)
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func GetObjectiveController(db *gorm.DB, id uint) (models.Objective, error) {
	projects, err := models.GetObjective(db, id)
	if err != nil {
		return models.Objective{}, err
	}
	return projects, nil
}

func UpdateObjectiveController(db *gorm.DB, pd models.ObjectiveUpdateDTO) error {
	objective := new(models.Objective)
	if er := db.First(&objective, pd.ID).Error; er != nil {
		return er
	}
	objective.Description = pd.Description
	objective.Name = pd.Name
	objective.EstimatedEndDate = pd.EstimatedEndDate
	objective.Overseer = pd.Overseer

	err := models.UpdateObjective(db, objective)
	if err != nil {
		return err
	}
	return nil
}

func DeleteObjectiveController(db *gorm.DB, id uint) error {
	err := models.DeleteObjective(db, id)
	if err != nil {
		return err
	}
	return nil
}
