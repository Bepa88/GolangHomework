package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Class struct {
	Students  []Student
	Teacher   Teacher
	ClassName string
}

type Student struct {
	ID        int
	FirstName string
	LastName  string
	Type      string
	GPA       float64
}

type StudentWithoutGPA struct {
	ID        int
	FirstName string
	LastName  string
}

type Teacher struct {
	ID        int
	FirstName string
	LastName  string
	Type      string
}

type User struct {
	Username string
	Password string
	Type     string
}

var class = Class{
	Students: []Student{
		{ID: 1, FirstName: "Oleg", LastName: "Levko", Type: "Student", GPA: 3.5},
		{ID: 2, FirstName: "Ira", LastName: "Grab", Type: "Student", GPA: 5},
	},
	Teacher:   Teacher{ID: 1, FirstName: "Viktor", LastName: "Viktorov", Type: "Teacher"},
	ClassName: "5-A",
}

var users = []User{
	{
		Username: "Viktor",
		Password: "12345",
		Type:     "Teacher",
	},
	{
		Username: "Oleg",
		Password: "54321",
		Type:     "Student",
	},
	{
		Username: "Ira",
		Password: "112233",
		Type:     "Student",
	},
}

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello!")
	})

	mux.HandleFunc("/class", processClass)
	mux.HandleFunc("/student/{id}", checkAuthForGetStudent(getStudent))

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error happenned", err.Error())
		return
	}
}

func processClass(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		checkAuthForGetClass(getClassForTeacher)(w, r)
	case http.MethodPut:
		putClass(w, r)
	}
}

func getClassForStudents(w http.ResponseWriter, r *http.Request) {

	classWithoutGPA := struct {
		Students  []StudentWithoutGPA
		Teacher   Teacher
		ClassName string
	}{
		Teacher:   class.Teacher,
		ClassName: class.ClassName,
	}

	for _, student := range class.Students {
		classWithoutGPA.Students = append(classWithoutGPA.Students, StudentWithoutGPA{
			ID:        student.ID,
			FirstName: student.FirstName,
			LastName:  student.LastName,
		})
	}

	err := json.NewEncoder(w).Encode(classWithoutGPA)
	if err != nil {
		println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func getClassForTeacher(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(class)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func putClass(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&class)
	if err != nil {
		println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(class)
	if err != nil {
		println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func getStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}
	for _, student := range class.Students {
		if student.ID == id {
			err := json.NewEncoder(w).Encode(student)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
	}

	http.Error(w, "Student not found", http.StatusNotFound)
}

func checkAuthForGetStudent(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		for _, user := range users {
			if user.Username == username && user.Password == password {
				if user.Type != "Teacher" {
					w.WriteHeader(http.StatusForbidden)
					return
				}
				next.ServeHTTP(w, r)
				return
			}
		}

		w.WriteHeader(http.StatusUnauthorized)
	}
}

func checkAuthForGetClass(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		for _, user := range users {
			if user.Username == username && user.Password == password {
				if user.Type != "Teacher" {
					getClassForStudents(w, r)
					return
				}
				next.ServeHTTP(w, r)
				return
			}
		}

		w.WriteHeader(http.StatusUnauthorized)
	}
}
