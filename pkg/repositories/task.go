package repositories

import (
	"github.com/IvanGelo1/go-todo-crud/pkg/config"
	"github.com/IvanGelo1/go-todo-crud/pkg/models"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Task struct {
	*models.Task
}

func init() {
	db = config.GetDB()
	db.AutoMigrate(&models.Task{})
}

func GetAllTasks() []models.Task {
	var allTasks []models.Task
	db.Find(&allTasks)
	return allTasks
}

func GetTaskById(ID int64) (*models.Task, *gorm.DB) {
	var task models.Task
	db := db.Where("ID=?", ID).Find(&task)
	return &task, db
}

func CreateTask(task *Task) *Task {
	db.NewRecord(task)
	db.Create(&task)
	return task
}
func UpdateTask(ID int64, t *Task) *models.Task {
	taskDetails, db := GetTaskById(ID)
	if t.Name != "" {
		taskDetails.Name = t.Name
	}
	if t.Description != "" {
		taskDetails.Description = t.Description
	}
	if t.Done {
		taskDetails.Done = t.Done
	}
	db.Save(&taskDetails)
	return taskDetails
}

func DeleteTask(ID int64) {
	var task models.Task
	db.Where("ID=?", ID).Delete(task)
}
