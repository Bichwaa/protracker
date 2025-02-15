package models

import (
	"gorm.io/gorm"
)

func CreateProject(db *gorm.DB, projectData *Project) error {
	cur := db.Create(projectData)
	if cur.Error != nil {
		return cur.Error
	}
	return nil
}

func GetProjects(db *gorm.DB) ([]Project, error) {
	var projects []Project
	cur := db.Preload("Notes").Preload("Objectives").Find(&projects)
	if cur.Error != nil {
		return nil, cur.Error
	}
	return projects, nil
}

func GetProject(db *gorm.DB, id uint) (Project, error) {
	var project Project
	cur := db.First(&project, id)
	if cur.Error != nil {
		return Project{}, cur.Error
	}
	return project, nil
}

func UpdateProject(db *gorm.DB, projectData Project) error {
	cur := db.Save(&projectData)
	if cur.Error != nil {
		return cur.Error
	}
	return nil
}

func DeleteProject(db *gorm.DB, id uint) error {
	cur := db.Delete(&Project{}, id)
	if cur.Error != nil {
		return cur.Error
	}
	return nil
}
