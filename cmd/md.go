package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/xylecara/gee-lister/handler"
)

//Stands for mark done
func MD(taskList handler.TaskList, tasksJSON string) error {
	hasTask := false
	if len(os.Args) > 2 {
		for index, value := range taskList.Tasks {
			if os.Args[2] == strconv.Itoa(value.ID) || os.Args[2] == value.Task {
				taskList.Tasks[index].Status = "done"
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
		fmt.Println("Not enough arguments, use it like: glister md <task-name-or-id>")
	}

	return nil
}