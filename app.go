package main

import (
	"context"
	"fmt"
	"protracker/controllers"
	"protracker/models"

	"gorm.io/gorm"
)

// App struct
type App struct {
	ctx context.Context
	db  gorm.DB
}

// NewApp creates a new App application struct
func NewApp(db gorm.DB) *App {
	var app = &App{}
	app.db = db
	return app
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (c *App) CreateProject(payload models.Project) error {
	c.ctx.Value("key")
	return controllers.CreateProjectController(&c.db, payload)
}

func (c *App) FetchProject(id uint) (models.Project, error) {
	return controllers.GetProjectController(&c.db, id)
}
func (c *App) FetchProjects() ([]models.Project, error) {
	return controllers.GetProjectsController(&c.db)
}

func (c *App) UpdateProject(pd models.ProjectUpdateDTO) error {
	return controllers.UpdateProjectController(&c.db, pd)
}

func (c *App) DeleteProject(id uint) error {
	fmt.Println("the ID at app.go level is =====>", id)
	return controllers.DeleteProjectController(&c.db, id)
}
