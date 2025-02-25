package controllers

import (
	"protracker/models"

	"gorm.io/gorm"
)

func CreateTaskController(db *gorm.DB, t models.Task) error {
	if err := models.CreateTask(db, &t); err != nil {
		return err
	}
	return nil
}

func GetTasksController(db *gorm.DB) ([]models.Task, error) {
	tasks, err := models.GetTasks(db)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func GetTaskController(db *gorm.DB, id uint) (models.Task, error) {
	task, err := models.GetTask(db, id)
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func UpdateTaskController(db *gorm.DB, td models.TaskUpdateDTO) error {
	task := new(models.Task)
	if er := db.First(&task, td.ID).Error; er != nil {
		return er
	}
	task.Name = td.Name
	task.EstimatedEndDate = td.EstimatedEndDate

	err := models.UpdateTask(db, task)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTaskController(db *gorm.DB, id uint) error {
	err := models.DeleteTask(db, id)
	if err != nil {
		return err
	}
	return nil
}

func MarkTaskAsCompleted(db *gorm.DB, id uint) error {
	task := new(models.Task)
	if er := db.First(&task, id).Error; er != nil {
		return er
	}
	if task.Progress == 100 {
		task.Progress = 0 // Set progress to 0%
	} else {
		task.Progress = 100 // Set progress to 100%
	}

	err := models.UpdateTask(db, task)
	if err != nil {
		return err
	}
	return nil
}
