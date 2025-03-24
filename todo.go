package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
)

func addTodo(db *sql.DB) {

	reader := bufio.NewReader(os.Stdin) // Create a reader for user input
	fmt.Print("Enter your task: ")

	// Read full line input, including spaces
	task, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	task = strings.TrimSpace(task) // Remove trailing newline
	if task == "" {
		fmt.Println("Task cannot be empty!")
		return
	}
	InsertTask(db, task)
}

func listTodo(db *sql.DB) {
	ListTasks(db)
}

func markAsDone(db *sql.DB) {
	var id int
	fmt.Print("Enter task ID to mark as done: ")
	fmt.Scan(&id)
	MarkTaskDone(db, id)
}

func main() {
	fmt.Println("Welcome to Todo App")

	db := OpenDatabase()
	defer db.Close()

	CreateTable(db)

	for {

		fmt.Println("\n1. Add Todo")
		fmt.Println("2. List Todo")
		fmt.Println("3. Mark as done")
		fmt.Println("4. Exit")
		fmt.Println("Enter your choice : ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addTodo(db)
		case 2:
			listTodo(db)
		case 3:
			markAsDone(db)
		case 4:
			fmt.Println("Thanks for using Todo App")
			return
		default:
			fmt.Println("Invalid choice")
		}

	}
}
