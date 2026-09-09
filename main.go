package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/xylecara/gee-lister/commands"
	"github.com/xylecara/gee-lister/formatter"
	"github.com/xylecara/gee-lister/handler"
)

func main() {
	timeNow := time.Now()
	timeNowStr := timeNow.Format("2006-Jan-02 15:04:05")
	tasksJSON := filepath.Join("json", "tasks.json")
	taskList, err := handler.Read(tasksJSON)
	if err != nil {
		log.Fatal("could not read tasks json, check if it is missing or renamed(must be named tasks.json)")
	}

	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments, try again")
		return
	}

	switch os.Args[1] {
	case "add":
		err := commands.Add(taskList, timeNowStr, tasksJSON)
		if err != nil {
			log.Fatal()
		}

	case "update":
		if len(os.Args) > 4 {
			switch os.Args[3] {
			case "name":
				for index, value := range taskList.Tasks {
					if os.Args[2] == strconv.Itoa(value.ID) || os.Args[2] == value.Task {
						taskList.Tasks[index].Task = os.Args[4]
						break
					}
				}
			case "description":
				for index, value := range taskList.Tasks {
					if os.Args[2] == strconv.Itoa(value.ID) || os.Args[2] == value.Task {
						taskList.Tasks[index].Description = os.Args[4]
						break
					}
				}
			default:
				fmt.Println("Unknown command")
			}

			for index, value := range taskList.Tasks {
					if os.Args[2] == strconv.Itoa(value.ID) || os.Args[2] == value.Task {
						taskList.Tasks[index].UpdatedAt = timeNowStr
						break
					}
				}

			err := handler.Write(tasksJSON, &taskList)
			if err != nil {
				log.Fatal()
			}

		} else {
			fmt.Println("Not enough arguments, use it like: \nglister update <task-name-or-id> name <new-name>\nglister update <task-name-or-id> description <new-desc>")
		} 
		
	case "delete":
		if len(os.Args) > 2 {
			for index, value := range taskList.Tasks {
				if os.Args[2] == strconv.Itoa(value.ID) || os.Args[2] == value.Task {
					taskList.Tasks = append(taskList.Tasks[:index], taskList.Tasks[index+1:]... )
					break
				}
			}
			taskList = formatter.FixIds(taskList)
			
			err := handler.Write(tasksJSON, &taskList)
			if err != nil {
				log.Fatal()
			}
		} else {
			fmt.Println("Not enough arguments, us it like: glister delete <task-name-or-id>")
		}

	case "mark-in-progress":
		fmt.Println("Mark-in-progress worked successfully!")
	case "mark-done":
		fmt.Println("Mark-done worked successfully!")
	case "list":
		commands.List(taskList)		

	default:
		fmt.Println("Unknown command")
	}
}
