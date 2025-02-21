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

type Goal struct {
	gorm.Model
	Name             string    `gorm:"not null"`
	EstimatedEndDate time.Time `gorm:"not null"`
	Progress         int       `gorm:"not null"`
	Description      string
	ObjectiveID      uint   `gorm:"not null"`
	Notes            []Note `gorm:"foreignKey:GoalID;constraint:OnDelete:CASCADE;"`
}

type GoalUpdateDTO struct {
	ID               uint `gorm:"primarykey"`
	UpdatedAt        time.Time
	Name             string    `gorm:"not null"`
	EstimatedEndDate time.Time `gorm:"not null"`
	Progress         int       `gorm:"not null"`
	Description      string
}

type Task struct {
	gorm.Model
	Name             string    `gorm:"not null"`
	EstimatedEndDate time.Time `gorm:"not null"`
	Progress         int       `gorm:"not null"`
	Description      string
	Notes            []Note `gorm:"foreignKey:TaskID;constraint:OnDelete:CASCADE;"`
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
