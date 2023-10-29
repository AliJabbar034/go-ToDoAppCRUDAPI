package model

import (
	"github.com/alijabbar034/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Task struct {
	gorm.Model
	UserId      int64  `json:"userid"`
	Title       string `json:"title"`
	Description string `json:"description"`
	StartedAt   string `json:"startedat"`
	EndAt       string `json:"endat"`
}

func Init() {
	config.ConnectDB()
	db = config.GETDB()
	db.AutoMigrate(&Task{})
}

func (t *Task) CreateTask() *Task {
	db.Create(&t)
	return t
}

func GetAllTasks() []*Task {
	var tasks []*Task
	db.Find(&tasks)
	return tasks

}

func DeleteTask(id int64) {
	var t Task

	db.Delete(&t, id)

}

func GetATask(id int64) *Task {

	var t *Task

	db.Find(&t, id)
	return t

}
func UpdatedTask(id int64) (*Task, *gorm.DB) {

	var t *Task

	db.Find(&t, id)
	return t, db

}
