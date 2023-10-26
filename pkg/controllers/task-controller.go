package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/IvanGelo1/go-todo-crud/pkg/repositories"
	"github.com/IvanGelo1/go-todo-crud/pkg/utils"
	"github.com/gorilla/mux"
)

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	var allTasks = repositories.GetAllTasks()
	var parsedTasks, _ = json.Marshal(allTasks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(parsedTasks)
}

func GetTaskById(w http.ResponseWriter, r *http.Request) {
	var variables = mux.Vars(r)
	var taskId = variables["id"]
	var ID, err = strconv.ParseInt(taskId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
		panic(err)
	}
	var task, _ = repositories.GetTaskById(ID)
	var parsedTask, _ = json.Marshal(task)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(parsedTask)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var taskModel = &repositories.Task{}
	utils.ParseBody(r, taskModel)
	var task = repositories.CreateTask(taskModel)
	var parsedTask, _ = json.Marshal(task)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(parsedTask)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	var variables = mux.Vars(r)
	var taskId = variables["id"]
	var ID, err = strconv.ParseInt(taskId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
		panic(err)
	}
	repositories.DeleteTask(ID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	var updateTask = &repositories.Task{}
	utils.ParseBody(r, updateTask)
	var variables = mux.Vars(r)
	var taskId = variables["id"]
	ID, err := strconv.ParseInt(taskId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
		panic(err)
	}
	updatedTask := repositories.UpdateTask(ID, updateTask)
	parsedTask, _ := json.Marshal(updatedTask)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(parsedTask)
}
