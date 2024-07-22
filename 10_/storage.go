package main

import (
	"fmt"
	"sort"
	"sync"
)

type Storage struct {
	m        sync.Mutex
	lastID   int
	allTasks map[int]Task
	allUsers map[string]User
}

func NewStorage() *Storage {
	return &Storage{
		allTasks: make(map[int]Task),
		allUsers: make(map[string]User),
	}
}

func (s *Storage) GetAllTasks() []Task {
	s.m.Lock()
	defer s.m.Unlock()

	var tasks = make([]Task, 0, len(s.allTasks))

	for _, t := range s.allTasks {
		tasks = append(tasks, t)
	}

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].ID < tasks[j].ID
	})

	return tasks
}

func (s *Storage) CreateOneTask(t Task) int {
	s.m.Lock()
	defer s.m.Unlock()

	fmt.Println("Trying to create task")
	t.ID = s.lastID + 1
	s.allTasks[t.ID] = t
	s.lastID++
	fmt.Printf("Created task. Last ID: %v\n", s.lastID)
	return t.ID
}

func (s *Storage) GetTaskByID(id int) (Task, bool) {
	s.m.Lock()
	defer s.m.Unlock()

	t, ok := s.allTasks[id]
	return t, ok
}

func (s *Storage) DeleteTaskByID(id int) bool {
	s.m.Lock()
	defer s.m.Unlock()

	_, ok := s.allTasks[id]
	if !ok {
		return false
	}

	delete(s.allTasks, id)
	return true
}

func (s *Storage) GetUserByUsername(username string) (User, bool) {
	s.m.Lock()
	defer s.m.Unlock()

	u, ok := s.allUsers[username]
	return u, ok
}

func (s *Storage) CreateOneUser(u User) bool {
	s.m.Lock()
	defer s.m.Unlock()

	_, ok := s.allUsers[u.Username]
	if ok {
		return false
	}

	s.allUsers[u.Username] = u
	return true
}
