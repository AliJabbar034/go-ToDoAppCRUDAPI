package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/alijabbar034/pkg/model"
	"github.com/gorilla/mux"
)

var Task *model.Task

func WelcomeRoute(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome to the Google")

}

func CreateTasks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	//EXpected Json Format
	CreatedTask := &model.Task{}

	err := json.NewDecoder(r.Body).Decode(&CreatedTask)
	fmt.Println("Task created successfully")

	if err != nil {

		http.Error(w, "Failed to decode Data", http.StatusBadRequest)
		return
	}

	if CreatedTask.Title == "" || CreatedTask.Description == "" || CreatedTask.StartedAt == "" || CreatedTask.EndAt == "" {
		errorMessage := map[string]string{"message": "Please provide all required fields"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	d := CreatedTask.CreateTask()

	eror := json.NewEncoder(w).Encode(d)

	if eror != nil {

		http.Error(w, "Failed to encode", http.StatusInternalServerError)
		return
	}

	r.Body.Close()

}

func GetTasks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	tasks := model.GetAllTasks()
	eror := json.NewEncoder(w).Encode(tasks)

	if eror != nil {
		http.Error(w, "Error decoding tasks", http.StatusInternalServerError)
		return
	}
	r.Body.Close()

}

func DeleteById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	taskId := vars["id"]

	id, err := strconv.ParseInt(taskId, 10, 64)

	if err != nil {
		http.Error(w, "Error parsing taskId", http.StatusInternalServerError)
		return
	}
	model.DeleteTask(id)

	json.NewEncoder(w).Encode("DeleteTask SuccessFUlly")
	r.Body.Close()

}

func GetTaskById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	taskId := vars["id"]
	id, err := strconv.ParseInt(taskId, 10, 64)

	if err != nil {
		http.Error(w, "Error parsing taskId", http.StatusInternalServerError)
		return
	}

	task := model.GetATask(id)

	eror := json.NewEncoder(w).Encode(task)
	if eror != nil {
		http.Error(w, "Error encoding task ", http.StatusInternalServerError)
		return
	}
	r.Body.Close()

}

func UpdatedTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var task = &model.Task{}

	json.NewDecoder(r.Body).Decode(&task)
	vars := mux.Vars(r)

	taskId := vars["id"]
	id, err := strconv.ParseInt(taskId, 10, 64)

	if err != nil {
		http.Error(w, "Error parsing taskId", http.StatusInternalServerError)
		return
	}

	taskDetail, db := model.UpdatedTask(id)

	if task.Title != "" {
		taskDetail.Title = task.Title
	}
	if task.Description != "" {
		taskDetail.Description = task.Description
	}
	if task.StartedAt != "" {
		taskDetail.StartedAt = task.StartedAt
	}
	db.Save(&taskDetail)

	json.NewEncoder(w).Encode(taskDetail)
	r.Body.Close()
}
