package main

import (
	"flag"
	"fmt"
	"homeWork_12_/internal/storage"
	"os"
)

func main() {
	s := storage.NewStorage("passwords.txt")

	if len(os.Args) < 2 {
		fmt.Println("Expected 'save', 'get', or 'list' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "save":
		saveCmd := flag.NewFlagSet("save", flag.ExitOnError)
		name := saveCmd.String("name", "", "Name of the service")
		password := saveCmd.String("password", "", "Password for the service")
		saveCmd.Parse(os.Args[2:])

		if *name == "" || *password == "" {
			fmt.Println("Both name and password are required")
			saveCmd.Usage()
			os.Exit(1)
		}

		s.Save(*name, *password)
		fmt.Println("Password saved successfully")

	case "get":
		getCmd := flag.NewFlagSet("get", flag.ExitOnError)
		name := getCmd.String("name", "", "Name of the service")
		getCmd.Parse(os.Args[2:])

		if *name == "" {
			fmt.Println("Name is required")
			getCmd.Usage()
			os.Exit(1)
		}

		password, found := s.Get(*name)
		if found {
			fmt.Printf("Password for %s: %s\n", *name, password)
		} else {
			fmt.Printf("No password found for %s\n", *name)
		}

	case "list":
		names := s.List()
		if len(names) == 0 {
			fmt.Println("No passwords stored")
		} else {
			fmt.Println("Saved passwords:")
			for _, name := range names {
				fmt.Println(name)
			}
		}

	default:
		fmt.Println("Expected 'save', 'get', or 'list' subcommands")
		os.Exit(1)
	}
}
