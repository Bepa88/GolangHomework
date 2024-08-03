package main

import (
	"fmt"
	"homeWork_12_/internal/storage"
	"strings"
)

func main() {
	s := storage.NewStorage("passwords.txt")

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. List saved passwords")
		fmt.Println("2. Save a new password")
		fmt.Println("3. Retrieve a saved password")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")

		var choice string
		fmt.Scan(&choice)
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			listPasswords(s)
		case "2":
			savePassword(s)
		case "3":
			retrievePassword(s)
		case "4":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func listPasswords(s *storage.Storage) {
	fmt.Println("\nSaved passwords:")
	names := s.GetAllNames()
	for _, name := range names {
		fmt.Println(name)
	}
}

func savePassword(s *storage.Storage) {
	fmt.Print("Enter name: ")
	var name string
	fmt.Scan(&name)
	name = strings.TrimSpace(name)

	fmt.Print("Enter password: ")
	var password string
	fmt.Scan(&password)
	password = strings.TrimSpace(password)

	s.Save(name, password)
	fmt.Println("Password saved.")
}

func retrievePassword(s *storage.Storage) {
	fmt.Print("Enter name: ")
	var name string
	fmt.Scan(&name)
	name = strings.TrimSpace(name)

	password, found := s.Get(name)
	if found {
		fmt.Printf("Password for %s: %s\n", name, password)
	} else {
		fmt.Println("Password not found.")
	}
}
