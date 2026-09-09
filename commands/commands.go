package commands

import (
	"fmt"
	"github.com/xylecara/gee-lister/handler"
)

func List(tasklist handler.TaskList) {
	for _, value := range tasklist.Tasks {
			fmt.Println(value)
		}
}