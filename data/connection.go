package data

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"protracker/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func getDBPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	appDir := filepath.Join(homeDir, ".protracker")
	if err := os.MkdirAll(appDir, 0755); err != nil {
		return "", err
	}
	return filepath.Join(appDir, "projects.db"), nil
}

func InitializeDB() (gorm.DB, error) {
	dbPath, er := getDBPath()
	if er != nil {
		return gorm.DB{}, er
	}
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
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
