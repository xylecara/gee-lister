package commands

import (
	"fmt"
	"os"
	"strconv"

	"github.com/xylecara/gee-lister/handler"
	"github.com/xylecara/gee-lister/formatter"
)

var hasTask = false

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
			fmt.Println("Not enough arguments, use it like: glister add <task-name> <task-description>")
		}

		return nil
}

func Update(taskList handler.TaskList, timeNow, tasksJSON string) error {
	if len(os.Args) > 4 {
			switch os.Args[3] {
			case "name":
				for index, value := range taskList.Tasks {
					if os.Args[2] == strconv.Itoa(value.ID) || os.Args[2] == value.Task {
						taskList.Tasks[index].Task = os.Args[4]
						break
					}
				}
			case "description":
				for index, value := range taskList.Tasks {
					if os.Args[2] == strconv.Itoa(value.ID) || os.Args[2] == value.Task {
						taskList.Tasks[index].Description = os.Args[4]
						break
					}
				}
			default:
				fmt.Println("Unknown command")
			}

			for index, value := range taskList.Tasks {
					if os.Args[2] == strconv.Itoa(value.ID) || os.Args[2] == value.Task {
						taskList.Tasks[index].UpdatedAt = timeNow
						hasTask = true
						break
					}
				}
			
			if !hasTask {
			fmt.Printf("no task with name or id %v found\n", os.Args[2])
			} else {
				err := handler.Write(tasksJSON, &taskList)
				if err != nil {
					return err
				}
			}
				
		} else {
			fmt.Println("Not enough arguments, use it like: \nglister update <task-name-or-id> name <new-name>\nglister update <task-name-or-id> description <new-desc>")
		}

		return nil
}

func Delete(taskList handler.TaskList, tasksJSON string) error {
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
			fmt.Println("Not enough arguments, use it like: glister delete <task-name-or-id>")
		}

		return nil
}

//Stands for mark in progress
func MIP(taskList handler.TaskList, tasksJSON string) error {
	if len(os.Args) > 2 {
		for index, value := range taskList.Tasks {
			if os.Args[2] == strconv.Itoa(value.ID) || os.Args[2] == value.Task {
				taskList.Tasks[index].Status = "in-progress"
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
		fmt.Println("Not enough arguments, use it like: glister mip <task-name-or-id>")
	}

	return nil
}

//Stands for mark done
func MD(taskList handler.TaskList, tasksJSON string) error {
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

//Stands for mark stop
func MS(taskList handler.TaskList, tasksJSON string) error {
	if len(os.Args) > 2 {
		for index, value := range taskList.Tasks {
			if os.Args[2] == strconv.Itoa(value.ID) || os.Args[2] == value.Task {
				taskList.Tasks[index].Status = "todo"
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

func List(taskList handler.TaskList) {
	for _, value := range taskList.Tasks {
			fmt.Println(value)
		}
}