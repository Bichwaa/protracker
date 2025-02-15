package controllers

import (
	"protracker/models"

	// "github.com/gofiber/fiber"
	"gorm.io/gorm"
)

func CreateProjectController(db *gorm.DB, p models.Project) error {
	if err := models.CreateProject(db, &p); err != nil {
		return err
	}
	return nil
}

func GetProjectsController(db *gorm.DB) ([]models.Project, error) {
	projects, err := models.GetProjects(db)
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func GetProjectController(db *gorm.DB, id uint) (models.Project, error) {
	projects, err := models.GetProject(db, id)
	if err != nil {
		return models.Project{}, err
	}
	return projects, nil
}

func UpdateProjectController(db *gorm.DB, pd models.Project) error {
	err := models.UpdateProject(db, pd)
	if err != nil {
		return err
	}
	return nil
}

func DeleteProjectController(db *gorm.DB, id uint) error {
	err := models.DeleteProject(db, id)
	if err != nil {
		return err
	}
	return nil
}
