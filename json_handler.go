package main

import (
	"encoding/json"
	"os"
)

type TaskList struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	ID          int    `json:"id"`
	Task        string `json:"task"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	IsFinished  bool   `json:"isFinished"`
}

// From json to go struct
func Read(jsonFilePath string) (TaskList, error) {
	file, err := os.Open(jsonFilePath)
	if err != nil {
		return TaskList{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	taskList := TaskList{}
	err = decoder.Decode(&taskList)
	if err != nil {
		return TaskList{}, err
	}

	return taskList, nil
}

// From go struct to json
func Write(jsonFilePath string, taskList *TaskList) error {
	file, err := os.Create(jsonFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")
	err = encoder.Encode(&taskList)
	if err != nil {
		return err
	}

	return nil
}
