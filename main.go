package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	tasksJSON := filepath.Join("json", "tasks.json")
	tasks, err := Read(tasksJSON)
	if err != nil {
		log.Fatal("could not read tasks json, check if it is missing or renamed(must be named tasks.json)")
	}

	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments, try again")
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) > 3 {
			tasks.Tasks = append(tasks.Tasks, Task{os.Args[2], false})

			err := Write(tasksJSON, &tasks)
			if err != nil {
				log.Fatal()
			}
		} else {
			fmt.Println("Not enough arguments, try again")
		}

	case "update":
		fmt.Println("Update worked successfully!")
	case "delete":
		fmt.Println("Delete worked successfully!")
	case "mark-in-progress":
		fmt.Println("Mark-in-progress worked successfully!")
	case "mark-done":
		fmt.Println("Mark-done worked successfully!")
	case "list":
		for _, value := range tasks.Tasks {
			fmt.Println(value)
		}

	default:
		fmt.Println("Unknown command")
	}
}
