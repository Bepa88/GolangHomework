package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	//mux := http.NewServeMux()
	mux := mux.NewRouter()

	s := NewStorage()

	tasks := TaskResource{
		s: s,
	}

	users := UserResource{
		s: s,
	}

	auth := Auth{
		s: s,
	}

	// mux.HandleFunc("POST /users", users.CreateOne)
	// mux.HandleFunc("GET /tasks", auth.checkAuth(tasks.GetAll))
	// mux.HandleFunc("POST /tasks", auth.checkAuth(tasks.CreateOne))
	// mux.HandleFunc("DELETE /tasks/{id}", auth.checkAuth(tasks.DeleteOne))

	mux.HandleFunc("/users", users.CreateOne).Methods("POST")
	mux.HandleFunc("/tasks", auth.checkAuth(tasks.GetAll)).Methods("GET")
	mux.HandleFunc("/tasks", auth.checkAuth(tasks.CreateOne)).Methods("POST")
	mux.HandleFunc("/tasks/{id}", auth.checkAuth(tasks.DeleteOne)).Methods("DELETE")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Printf("Failed to listen and serve: %v\n", err)
	}

}

type TaskResource struct {
	s *Storage
}

func (t *TaskResource) GetAll(w http.ResponseWriter, r *http.Request) {
	tasks := t.s.GetAllTasks()

	err := json.NewEncoder(w).Encode(tasks)
	if err != nil {
		fmt.Printf("Failed to encode: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (t *TaskResource) CreateOne(w http.ResponseWriter, r *http.Request) {
	var task Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		fmt.Printf("Failed to decode: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	task.ID = t.s.CreateOneTask(task)

	err = json.NewEncoder(w).Encode(task)
	if err != nil {
		fmt.Printf("Failed to encode: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (t *TaskResource) DeleteOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Printf("Invalid id param: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ok := t.s.DeleteTaskByID(taskID)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

type UserResource struct {
	s *Storage
}

func (ur *UserResource) CreateOne(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Printf("Failed to decode: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ok := ur.s.CreateOneUser(user)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
