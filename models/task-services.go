package models

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func CreateTask(db *gorm.DB, taskData *Task) error {
	cur := db.Create(taskData)
	if cur.Error != nil {
		return cur.Error
	}
	return nil
}

func GetTasks(db *gorm.DB) ([]Task, error) {
	var tasks []Task
	cur := db.Preload("Notes").Find(&tasks)
	if cur.Error != nil {
		return nil, cur.Error
	}
	return tasks, nil
}

func GetTask(db *gorm.DB, id uint) (Task, error) {
	var task Task
	task.ID = id
	cur := db.Preload("Notes").Find(&task, id)
	if cur.Error != nil {
		if errors.Is(cur.Error, gorm.ErrRecordNotFound) {
			return Task{}, fmt.Errorf("task with ID %d not found", id)
		}
		return Task{}, cur.Error
	}
	fmt.Println(task)
	return task, nil
}

func UpdateTask(db *gorm.DB, taskData *Task) error {
	cur := db.Save(taskData)
	if cur.Error != nil {
		return cur.Error
	}
	return nil
}

func DeleteTask(db *gorm.DB, id uint) error {
	var task Task
	task.ID = id
	db.Find(&task, id) //must be passed to context to be used by pre-delelete hook
	cur := db.Delete(&task, id)
	if cur.Error != nil {
		return cur.Error
	}
	return nil
}
