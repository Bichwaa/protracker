package data

import (
	"fmt"
	"log"

	"protracker/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitializeDB() (gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("projects.db"), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatal("fucking coulldn't connect to db")
		return *db, err
	} else {
		fmt.Println("Connection to database successful")
		err := db.AutoMigrate(
			&models.Project{},
			&models.Objective{},
			&models.Goal{},
			&models.Task{},
			&models.Note{},
		)
		if err != nil {
			log.Fatal("fucking coulldn't connect to db", err)
			panic("migration error")
		}
		fmt.Println("Migrated")
		return *db, nil
	}
}
