package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/xylecara/gee-lister/handler"
)

//Stands for mark in progress
func MIP(taskList handler.TaskList, tasksJSON string) error {
	hasTask := false
	if len(os.Args) > 2 {
		for index, value := range taskList.Tasks {
			if os.Args[2] == strconv.Itoa(value.ID) || os.Args[2] == value.Task {
				taskList.Tasks[index].Status = "in progress"
				hasTask = true
				break
			} 
		}

		if os.Args[2] == "help" {
		fmt.Println(`NAME:
	mip - Marks status as in progress

USAGE:
	glister mip <task-name-or-id>

EXAMPLES:
	$ glister mip Cycling
	$ glister mip "Ride a Bicycle"
	$ glister mip 1`)
	} else if !hasTask && os.Args[2] != "help" {
			fmt.Printf("no task with name or id %q found\n", os.Args[2])
		} else {
			err := handler.Write(tasksJSON, &taskList)
			if err != nil {
				return err
			}
		}

	} else {
		fmt.Println("Not enough arguments, proper usage:\nglister md <task-name-or-id>")
	}

	return nil
}