package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/xylecara/gee-lister/handler"
)

func Update(taskList handler.TaskList, timeNow, tasksJSON string) error {
	hasTask := false
	if len(os.Args) > 4 {
		switch os.Args[3] {
		case "name":
			for index, value := range taskList.Tasks {
				if os.Args[2] == strconv.Itoa(value.ID) || os.Args[2] == value.Task {
					taskList.Tasks[index].Task = os.Args[4]
					hasTask = true
					break
				}
			}
		case "description":
			for index, value := range taskList.Tasks {
				if os.Args[2] == strconv.Itoa(value.ID) || os.Args[2] == value.Task {
					taskList.Tasks[index].Description = os.Args[4]
					hasTask = true
					break
				}
			}
		default:
			fmt.Println("Unknown command, only available commands for update is:\nname\ndescription")
		}

		for index, value := range taskList.Tasks {
			if os.Args[2] == strconv.Itoa(value.ID) || os.Args[2] == value.Task {
				taskList.Tasks[index].UpdatedAt = timeNow
				hasTask = true
				break
			}
		}

		if !hasTask && os.Args[2] != "help" {
			fmt.Printf("no task with name or id %q found\n", os.Args[2])
		} else {
			err := handler.Write(tasksJSON, &taskList)
			if err != nil {
				return err
			}
		}

	} else if len(os.Args) == 3 {
		if os.Args[2] == "help" {
			fmt.Println(`NAME:
	update - Updates a task's name or description

USAGE:
	glister update <task-name-or-id> [command] <new-task-name-or-description>

COMMANDS:
	name - updates name
	description - updates description

EXAMPLES:
	$ glister update Cycling name "Ride a Bicycle"
	$ glister update "Ride a Bicycle" description "Go ride a bike"
	$ glister update 1 name Cycling`)
		}
	} else {
		fmt.Println("Not enough arguments, proper usage: \nglister update <task-name-or-id> name <new-name>\nglister update <task-name-or-id> description <new-desc>")
	}

	return nil
}
