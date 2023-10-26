package routes

import (
	"github.com/IvanGelo1/go-todo-crud/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterTaskRoutes = func(router *mux.Router) {
	router.HandleFunc("/tasks", controllers.GetAllTasks).Methods("GET")
	router.HandleFunc("/task/{id}", controllers.GetTaskById).Methods("GET")
	router.HandleFunc("/task", controllers.CreateTask).Methods("POST")
	router.HandleFunc("/task/{id}", controllers.UpdateTask).Methods("PUT")
	router.HandleFunc("/task/{id}", controllers.DeleteTask).Methods("DELETE")
}
