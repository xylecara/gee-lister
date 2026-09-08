package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func main() {
	timeNow := time.Now()
	timeNowStr := timeNow.Format("2006-Jan-02 15:04:05")
	tasksJSON := filepath.Join("json", "tasks.json")
	tasks, err := Read(tasksJSON)
	if err != nil {
		log.Fatal("could not read tasks json, check if it is missing or renamed(must be named tasks.json)")
	}
	lastTaskID := len(tasks.Tasks)

	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments, try again")
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) > 3 {
			tasks.Tasks = append(tasks.Tasks, Task{lastTaskID + 1, os.Args[2], os.Args[3], timeNowStr, timeNowStr, false})

			err := Write(tasksJSON, &tasks)
			if err != nil {
				log.Fatal()
			}
		} else {
			fmt.Println("Not enough arguments, use it like glister add <task-name> <task-description>")
		}

	case "update":
		fmt.Println("Update worked successfully!")
	case "delete":
		if len(os.Args) > 2 {
			for index, value := range tasks.Tasks {
				if os.Args[2] == strconv.Itoa(value.ID) || os.Args[2] == value.Task {
					tasks.Tasks = append(tasks.Tasks[:index], tasks.Tasks[index+1:]... )
					break
				}
			}
			tasks = fixIds(tasks)
			
			err := Write(tasksJSON, &tasks)
			if err != nil {
				log.Fatal()
			}
		} else {
			fmt.Println("Not enough arguments, us it like glister delete <task-name-or-id>")
		}

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
