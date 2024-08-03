package storage

import (
	"fmt"
	"os"
	"strings"
)

type Storage struct {
	filePath string
	data     map[string]string
}

func NewStorage(filePath string) *Storage {
	s := &Storage{
		filePath: filePath,
		data:     make(map[string]string),
	}
	s.loadFromFile()
	return s
}

func (s *Storage) Get(name string) (string, bool) {
	password, found := s.data[name]
	return password, found
}

func (s *Storage) Set(name, password string) {
	s.data[name] = strings.TrimSpace(password)
	s.saveToFile()
}

func (s *Storage) List() []string {
	names := make([]string, 0, len(s.data))
	for name := range s.data {
		names = append(names, name)
	}
	return names
}

func (s *Storage) loadFromFile() {
	file, err := os.Open(s.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var fileInfo os.FileInfo
	fileInfo, err = file.Stat()
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}

	if fileInfo.Size() == 0 {
		return
	}

	content := make([]byte, fileInfo.Size())
	_, err = file.Read(content)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		name := parts[0]
		password := parts[1]
		s.data[name] = password
	}
}

func (s *Storage) saveToFile() {
	file, err := os.Create(s.filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	for name, password := range s.data {
		_, err := file.WriteString(fmt.Sprintf("%s:%s\n", name, password))
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
}

func (s *Storage) Save(name, password string) {
	s.data[name] = password
	s.saveToFile()
}

func (s *Storage) GetAllNames() []string {
	var names []string
	for name := range s.data {
		names = append(names, name)
	}
	return names
}
