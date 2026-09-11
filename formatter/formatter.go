package formatter

import (
	"fmt"
	"os"
	"strings"

	"github.com/xylecara/gee-lister/handler"
)

//fixes the ids when deleting
func FixIds(taskList handler.TaskList) handler.TaskList {
	for index := range taskList.Tasks {
		taskList.Tasks[index].ID = index + 1
	}

	return taskList
}

func LayoutTasks(tasklist handler.TaskList, status string) {
	if status == "ip" {
		status = "in progress"
	}

	for index, value := range tasklist.Tasks {
		if len(os.Args) < 3 {
			layoutPrint(tasklist.Tasks[index])
		}
		
		if value.Status == status {
			layoutPrint(tasklist.Tasks[index])
		}
	}


}

func layoutPrint(task handler.Task) {
	fmt.Printf("%v. %s\n", task.ID, task.Task)
	fmt.Printf("\nCreated At: %s\n", task.CreatedAt)
	fmt.Printf("Updated At: %s\n", task.UpdatedAt)
	fmt.Printf("\nStatus: %s\n", task.Status)
	fmt.Printf("\nDescription:\n%s\n", task.Description)
			
	seperator := strings.Repeat("-", 20)

	fmt.Printf("\n%s\n", seperator)
}