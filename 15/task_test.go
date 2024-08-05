package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateTask(t *testing.T) {
	storage := NewStorage()
	taskResource := TaskResource{s: storage}

	task := Task{Title: "New Task"}
	body, _ := json.Marshal(task)

	req := httptest.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	taskResource.CreateOne(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status OK; got %v", resp.Status)
	}
}

func TestGetAllTasks(t *testing.T) {
	storage := NewStorage()
	taskResource := TaskResource{s: storage}

	req := httptest.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()

	taskResource.GetAll(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status OK; got %v", resp.Status)
	}
}
