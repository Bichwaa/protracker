package main

import (
	"context"
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
	return controllers.DeleteProjectController(&c.db, id)
}

func (c *App) CreateObjective(payload models.Objective) error {
	c.ctx.Value("key")
	return controllers.CreateObjectiveController(&c.db, payload)
}

func (c *App) FetchObjective(id uint) (models.Objective, error) {
	return controllers.GetObjectiveController(&c.db, id)
}
func (c *App) FetchObjectives() ([]models.Objective, error) {
	return controllers.GetObjectivesController(&c.db)
}

func (c *App) UpdateObjective(pd models.ObjectiveUpdateDTO) error {
	return controllers.UpdateObjectiveController(&c.db, pd)
}

func (c *App) DeleteObjective(id uint) error {
	return controllers.DeleteObjectiveController(&c.db, id)
}

func (c *App) CreateGoal(payload models.Goal) error {
	return controllers.CreateGoalController(&c.db, payload)
}

func (c *App) FetchGoals() ([]models.Goal, error) {
	return controllers.GetGoalsController(&c.db)
}

func (c *App) FetchGoal(id uint) (models.Goal, error) {
	return controllers.GetGoalController(&c.db, id)
}

func (c *App) UpdateGoal(gd models.GoalUpdateDTO) error {
	return controllers.UpdateGoalController(&c.db, gd)
}

func (c *App) DeleteGoal(id uint) error {
	return controllers.DeleteGoalController(&c.db, id)
}

func (c *App) CreateTask(payload models.Task) error {
	return controllers.CreateTaskController(&c.db, payload)
}

func (c *App) FetchTasks() ([]models.Task, error) {
	return controllers.GetTasksController(&c.db)
}

func (c *App) FetchTask(id uint) (models.Task, error) {
	return controllers.GetTaskController(&c.db, id)
}

func (c *App) UpdateTask(td models.TaskUpdateDTO) error {
	return controllers.UpdateTaskController(&c.db, td)
}

func (c *App) DeleteTask(id uint) error {
	return controllers.DeleteTaskController(&c.db, id)
}

func (c *App) MarkTaskAsCompleted(id uint) error {
	return controllers.MarkTaskAsCompleted(&c.db, id)
}
