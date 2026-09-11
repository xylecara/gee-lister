package cmd

import (
	"fmt"
	"strconv"
	"os"

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
			
			if !hasTask {
			fmt.Printf("no task with name or id %q found\n", os.Args[2])
			} else {
				err := handler.Write(tasksJSON, &taskList)
				if err != nil {
					return err
				}
			}
				
		} else {
			fmt.Println("Not enough arguments, proper usage: \nglister update <task-name-or-id> name <new-name>\nglister update <task-name-or-id> description <new-desc>")
		}

		return nil
}