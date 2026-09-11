package cmd

import (
	"fmt"
	"os"

	"github.com/xylecara/gee-lister/handler"
	"github.com/xylecara/gee-lister/formatter"
)

func List(taskList handler.TaskList) {
	stopLoop := false
	for _, value := range taskList.Tasks {
		if len(os.Args) > 2 {
			switch os.Args[2]{
			case "done":
				if value.Status == "done" {
					formatter.LayoutTasks(taskList, os.Args[2])
					stopLoop = true
				}

			case "todo":
				if value.Status == "todo" {
					formatter.LayoutTasks(taskList, os.Args[2])
					stopLoop = true
				}

			case "ip":
				if value.Status == "in progress" {
					formatter.LayoutTasks(taskList, os.Args[2])
					stopLoop = true
				}

			default:
				fmt.Println("Unknown command, only commands available for list are:\ndone, todo, ip\nUsage:\nglister list\nWith commands:\ngliister list done")
				stopLoop = true
			}

			if stopLoop {
				break
			}
		}
	}
	
	if len(os.Args) < 3 {
		formatter.LayoutTasks(taskList, "")
	}
	
}