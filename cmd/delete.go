package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/xylecara/gee-lister/handler"
	"github.com/xylecara/gee-lister/formatter"
)

func Delete(taskList handler.TaskList, tasksJSON string) error {
	hasTask := false
	if len(os.Args) > 2 {
			for index, value := range taskList.Tasks {
				if os.Args[2] == strconv.Itoa(value.ID) || os.Args[2] == value.Task {
					taskList.Tasks = append(taskList.Tasks[:index], taskList.Tasks[index+1:]... )
					hasTask = true
					break
				}
			}

			if !hasTask {
				fmt.Printf("no task with name or id %v found\n", os.Args[2])
			} else {
				taskList = formatter.FixIds(taskList)

				err := handler.Write(tasksJSON, &taskList)
				if err != nil {
					return err
				}
			}
			
		} else {
			fmt.Println("Not enough arguments, proper usage:\nglister delete <task-name-or-id>")
		}

		return nil
}