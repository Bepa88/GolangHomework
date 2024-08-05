package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUser(t *testing.T) {
	storage := NewStorage()
	userResource := UserResource{s: storage}

	user := User{Username: "testuser", Password: "testpass"}
	body, _ := json.Marshal(user)

	req := httptest.NewRequest("POST", "/users", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	userResource.CreateOne(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status OK; got %v", resp.Status)
	}
}

func TestCreateUserDuplicate(t *testing.T) {
	storage := NewStorage()
	userResource := UserResource{s: storage}

	user := User{Username: "testuser", Password: "testpass"}
	body, _ := json.Marshal(user)

	req := httptest.NewRequest("POST", "/users", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	userResource.CreateOne(w, req)

	req = httptest.NewRequest("POST", "/users", bytes.NewBuffer(body))
	w = httptest.NewRecorder()

	userResource.CreateOne(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status BadRequest; got %v", resp.Status)
	}
}
