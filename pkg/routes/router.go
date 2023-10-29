package routes

import (
	"log"
	"net/http"

	"github.com/alijabbar034/pkg/controllers"
	"github.com/gorilla/mux"
)

func InitializeRoute() {

	r := mux.NewRouter()

	r.HandleFunc("/", controllers.WelcomeRoute).Methods("GET")
	r.HandleFunc("/tasks", controllers.CreateTasks).Methods("POST")
	r.HandleFunc("/tasks", controllers.GetTasks).Methods("GET")
	r.HandleFunc("/task/{id}", controllers.DeleteById).Methods("DELETE")
	r.HandleFunc("/task/{id}", controllers.GetTaskById).Methods("GET")
	r.HandleFunc("/task/{id}", controllers.UpdatedTask).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", r))

}
