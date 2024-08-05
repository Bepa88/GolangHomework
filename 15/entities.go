package main

type Task struct {
	ID    int
	Title string
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
