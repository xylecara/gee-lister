package formatter

import "github.com/xylecara/gee-lister/handler"

//fixes the ids when deleting
func FixIds(taskList handler.TaskList) handler.TaskList {
	for index := range taskList.Tasks {
		taskList.Tasks[index].ID = index + 1
	}

	return taskList
}