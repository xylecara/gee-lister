package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/xylecara/gee-lister/cmd"
	"github.com/xylecara/gee-lister/handler"
)

func main() {
	timeNow := time.Now()
	timeNowStr := timeNow.Format("2006-Jan-02 15:04:05")
	tasksJSON := filepath.Join("json", "tasks.json")
	taskList, err := handler.Read(tasksJSON)
	if err != nil {
		log.Fatal("could not read tasks json, check if it is missing or renamed(must be named tasks.json)")
	}

	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments, try again")
		return
	}

	switch os.Args[1] {
	case "add":
		err := cmd.Add(taskList, timeNowStr, tasksJSON)
		if err != nil {
			log.Fatal(err.Error())
		}

	case "update":
		err := cmd.Update(taskList, timeNowStr, tasksJSON)
		if err != nil {
			log.Fatal(err.Error())
		} 
		
	case "delete":
		err := cmd.Delete(taskList, tasksJSON)
		if err != nil {
			log.Fatal(err.Error())
		}

	//stands for mark in progress
	case "mip":
		err := cmd.MIP(taskList, tasksJSON)
		if err != nil {
			log.Fatal(err.Error())
		}
	
	//stands for mark done
	case "md":
		err := cmd.MD(taskList, tasksJSON)
		if err != nil {
			log.Fatal(err.Error())
		}

	case "ms":
		err :=  cmd.MS(taskList, tasksJSON)
		if err != nil {
			log.Fatal(err.Error())
		}
		
	case "list":
		cmd.List(taskList)		
	
	default:
		fmt.Println("Unknown command")
	}
}
