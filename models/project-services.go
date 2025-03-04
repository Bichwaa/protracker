package models

import (
	"errors"
	"fmt"

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
	cur := db.Preload("Notes").Preload("Objectives", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Notes").Preload("Goals", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Tasks").Preload("Notes")
		})
	}).Find(&projects)
	if cur.Error != nil {
		return nil, cur.Error
	}
	return projects, nil
}

func GetProject(db *gorm.DB, id uint) (Project, error) {
	var project Project
	project.ID = id
	cur := db.Preload("Notes").Preload("Objectives", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Notes").Preload("Goals", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Tasks").Preload("Notes")
		})
	}).Find(&project, id)
	if cur.Error != nil {
		if errors.Is(cur.Error, gorm.ErrRecordNotFound) {
			return Project{}, fmt.Errorf("project with ID %d not found", id)
		}
		return Project{}, cur.Error
	}
	fmt.Println(project)
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
