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

func UpdateProjectController(db *gorm.DB, pd models.ProjectUpdateDTO) error {
	project := new(models.Project)
	if er := db.First(&project, pd.ID).Error; er != nil {
		return er
	}
	project.Description = pd.Description
	project.Name = pd.Name
	project.Tags = pd.Tags
	project.EstimatedEndDate = pd.EstimatedEndDate

	err := models.UpdateProject(db, *project)
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
