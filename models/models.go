package models

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name             string    `gorm:"not null"`
	EstimatedEndDate time.Time `gorm:"not null"`
	Progress         int       `gorm:"not null"`
	Description      string
	Overseer         string
	Tags             string      `gorm:"default:'all'"`
	Objectives       []Objective `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE;"`
	Notes            []Note      `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE;"`
}

type ProjectUpdateDTO struct {
	ID               uint `gorm:"primarykey"`
	UpdatedAt        time.Time
	Name             string `gorm:"not null"`
	Overseer         string
	EstimatedEndDate time.Time `gorm:"not null"`
	Progress         int       `gorm:"not null"`
	Description      string
	Tags             string `gorm:"default:'all'"`
}

func (p *Project) UpdateProjectProgress(tx *gorm.DB) (err error) {
	var totalProgress int
	var objectiveCount int

	var objectives []Objective
	if err := tx.Where("project_id = ? AND deleted_at IS NULL", p.ID).Find(&objectives).Error; err != nil {
		return err // Handle error if fetching tasks fails
	}

	for _, obj := range objectives {
		totalProgress += int(obj.Progress)
		objectiveCount++
	}

	if objectiveCount > 0 {
		p.Progress = totalProgress / objectiveCount
	} else {
		p.Progress = 0 // Set Progress to 0 if there are no goals
	}
	if err := tx.Save(p).Error; err != nil {
		return err // Handle error if saving fails
	}
	return
}

type Objective struct {
	gorm.Model
	Name             string    `gorm:"not null"`
	EstimatedEndDate time.Time `gorm:"not null"`
	Progress         int       `gorm:"not null"`
	Description      string
	Overseer         string
	ProjectID        uint   `gorm:"not null"`
	Goals            []Goal `gorm:"foreignKey:ObjectiveID;constraint:OnDelete:CASCADE;"`
	Notes            []Note `gorm:"foreignKey:ObjectiveID;constraint:OnDelete:CASCADE;"`
}

type ObjectiveUpdateDTO struct {
	ID               uint `gorm:"primarykey"`
	UpdatedAt        time.Time
	Name             string `gorm:"not null"`
	Overseer         string
	EstimatedEndDate time.Time `gorm:"not null"`
	Progress         int       `gorm:"not null"`
	Description      string
}

func (o *Objective) UpdateObjectiveProgress(tx *gorm.DB) (err error) {
	var totalProgress int
	var goalCount int

	var goals []Goal
	if err := tx.Where("objective_id = ? AND deleted_at IS NULL", o.ID).Find(&goals).Error; err != nil {
		return err // Handle error if fetching tasks fails
	}

	for _, goal := range goals {
		totalProgress += int(goal.Progress)
		goalCount++
	}

	if goalCount > 0 {
		o.Progress = totalProgress / goalCount
	} else {
		o.Progress = 0 // Set Progress to 0 if there are no goals
	}
	if err := tx.Save(o).Error; err != nil {
		return err // Handle error if saving fails
	}
	var proj Project
	if err := tx.First(&proj, o.ProjectID).Error; err != nil {
		return err
	}
	proj.UpdateProjectProgress(tx)
	return
}

func (o *Objective) AfterUpdate(tx *gorm.DB) (err error) {
	var proj Project
	if err := tx.First(&proj, o.ProjectID).Error; err != nil {
		return err // Handle error if Goal is not found
	}
	proj.UpdateProjectProgress(tx)
	return
}

func (o *Objective) AfterCreate(tx *gorm.DB) (err error) {
	var proj Project
	if err := tx.First(&proj, o.ProjectID).Error; err != nil {
		return err // Handle error if Goal is not found
	}
	proj.UpdateProjectProgress(tx)
	return
}

func (o *Objective) AfterDelete(tx *gorm.DB) (err error) {
	var proj Project
	if err := tx.First(&proj, o.ProjectID).Error; err != nil {
		return err // Handle error if Goal is not found
	}
	proj.UpdateProjectProgress(tx)
	return
}

type Goal struct {
	gorm.Model
	Name             string    `gorm:"not null"`
	EstimatedEndDate time.Time `gorm:"not null"`
	Progress         int       `gorm:"not null"`
	Description      string
	ObjectiveID      uint   `gorm:"not null"`
	Notes            []Note `gorm:"foreignKey:GoalID;constraint:OnDelete:CASCADE;"`
	Tasks            []Task `gorm:"foreignKey:GoalID;constraint:OnDelete:CASCADE;"`
}

type GoalUpdateDTO struct {
	ID               uint `gorm:"primarykey"`
	UpdatedAt        time.Time
	Name             string    `gorm:"not null"`
	EstimatedEndDate time.Time `gorm:"not null"`
	Progress         int       `gorm:"not null"`
	Description      string
}

func (g *Goal) UpdateProgress(tx *gorm.DB) (err error) {
	var totalProgress int
	var taskCount int

	var tasks []Task
	if err := tx.Where("goal_id = ? AND deleted_at IS NULL", g.ID).Find(&tasks).Error; err != nil {
		return err // Handle error if fetching tasks fails
	}

	for _, task := range tasks {
		totalProgress += int(task.Progress)
		taskCount++
	}

	if taskCount > 0 {
		g.Progress = totalProgress / taskCount
	} else {
		g.Progress = 0 // Set Progress to 0 if there are no goals
	}
	if err := tx.Save(g).Error; err != nil {
		return err // Handle error if saving fails
	}
	var obj Objective
	if err := tx.First(&obj, g.ObjectiveID).Error; err != nil {
		return err
	}
	obj.UpdateObjectiveProgress(tx)
	return
}

func (g *Goal) AfterUpdate(tx *gorm.DB) (err error) {
	var obj Objective
	if err := tx.First(&obj, g.ObjectiveID).Error; err != nil {
		return err // Handle error if Goal is not found
	}
	obj.UpdateObjectiveProgress(tx)
	return
}

func (g *Goal) AfterCreate(tx *gorm.DB) (err error) {
	var obj Objective
	if err := tx.First(&obj, g.ObjectiveID).Error; err != nil {
		return err // Handle error if Goal is not found
	}
	obj.UpdateObjectiveProgress(tx)
	return
}

func (g *Goal) AfterDelete(tx *gorm.DB) (err error) {
	var obj Objective
	if err := tx.First(&obj, g.ObjectiveID).Error; err != nil {
		return err // Handle error if Goal is not found
	}
	obj.UpdateObjectiveProgress(tx)
	return
}

type Task struct {
	gorm.Model
	Name             string    `gorm:"not null"`
	EstimatedEndDate time.Time `gorm:"not null"`
	Progress         int       `gorm:"not null"`
	GoalID           uint
	Notes            []Note `gorm:"foreignKey:TaskID;constraint:OnDelete:CASCADE;"`
}

type TaskUpdateDTO struct {
	ID               uint
	Name             string
	EstimatedEndDate time.Time
	Progress         uint8
}

func (t *Task) AfterUpdate(tx *gorm.DB) (err error) {
	var goal Goal
	if err := tx.First(&goal, t.GoalID).Error; err != nil {
		return err // Handle error if Goal is not found
	}
	goal.UpdateProgress(tx)
	return
}

func (t *Task) AfterCreate(tx *gorm.DB) (err error) {
	var goal Goal
	if err := tx.First(&goal, t.GoalID).Error; err != nil {
		return err // Handle error if Goal is not found
	}
	goal.UpdateProgress(tx)
	return
}

func (t *Task) AfterDelete(tx *gorm.DB) (err error) {
	var goal Goal
	if err := tx.First(&goal, t.GoalID).Error; err != nil {
		return err // Handle error if Goal is not found
	}
	goal.UpdateProgress(tx)
	return
}

type Note struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Content     string `gorm:"type:text;not null"`
	ProjectID   uint
	ObjectiveID uint
	GoalID      uint
	TaskID      uint
}

type NoteUpdateDTO struct {
	Title   string
	Content string
}
