package models

import "gorm.io/gorm"

func CreateTask(db *gorm.DB, taskData *Task) error {
	cur := db.Create(taskData)
	if cur.Error != nil {
		return cur.Error
	}
	return nil
}

func GetTask(db *gorm.DB, id uint) (*Task, error) {
	var task Task
	if err := db.First(&task, id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func UpdateTask(db *gorm.DB, taskData *Task) error {
	if err := db.Save(taskData).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTask(db *gorm.DB, id uint) error {
	if err := db.Delete(&Task{}, id).Error; err != nil {
		return err
	}
	return nil
}
