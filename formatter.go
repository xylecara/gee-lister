package main

//fixes the ids when deleting
func fixIds(taskList TaskList) TaskList {
	for index := range taskList.Tasks {
		taskList.Tasks[index].ID = index + 1
	}

	return taskList
}