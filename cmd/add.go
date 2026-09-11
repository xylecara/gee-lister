package cmd

import (
	"fmt"
	"os"

	"github.com/xylecara/gee-lister/handler"
)

func Add(tasklist handler.TaskList, timeNow, tasksJSON string) error {
	lastTaskID := len(tasklist.Tasks)

	if len(os.Args) > 3 {
			tasklist.Tasks = append(tasklist.Tasks, handler.Task{
				ID:lastTaskID + 1, 
				Task: os.Args[2], 
				Description: os.Args[3], 
				CreatedAt: timeNow, 
				UpdatedAt: timeNow, 
				Status: "todo"})

			err := handler.Write(tasksJSON, &tasklist)
			if err != nil {
				return err
			}

		} else {
			fmt.Println("Not enough arguments, proper usage:\nglister add <task-name> <task-description>")
		}

		return nil
}