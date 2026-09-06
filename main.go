package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments, try again")
		return
	}

	switch os.Args[1] {
	case "add":
		fmt.Println("Add worked successfully!")
	case "update":
		fmt.Println("Update worked successfully!")
	case "delete":
		fmt.Println("Delete worked successfully!")
	case "mark-in-progress":
		fmt.Println("Mark-in-progress worked successfully!")
	case "mark-done":
		fmt.Println("Mark-done worked successfully!")
	case "list":
		fmt.Println("List worked successfully!")
	default:
		fmt.Println("Unknown command")
	}
}
