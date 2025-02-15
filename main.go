package main

import (
	"embed"
	"protracker/data"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend
var assets embed.FS

// var db gorm.DB // Declare db variable

func main() {
	// Initialize the database
	var er error
	db, er := data.InitializeDB() // Initialize the database here
	if er != nil {
		println("Error initializing database:", er.Error())
		return // Exit if there's an error
	}

	// Create an instance of the app structure
	app := NewApp(db)
	// Create application with options
	err := wails.Run(&options.App{
		Title:  "protracker",
		Width:  1448,
		Height: 916,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
