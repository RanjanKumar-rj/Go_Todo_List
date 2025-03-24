package main

import "fmt"

func addTodo() {
	fmt.Println("Adding a new todo...")
	// Implement add todo functionality
}

func listTodo() {
	fmt.Println("Listing all todos...")
	// Implement list todo functionality
}

func markAsDone() {
	fmt.Println("Marking todo as done...")
	// Implement mark as done functionality
}

func main() {
	fmt.Println("Welcome to Todo App")
	for {
		fmt.Println("\n1. App Todo")
		fmt.Println("2. List Todo")
		fmt.Println("3. Mark as done")
		fmt.Println("4. Exit")
		fmt.Println("Enter your choice : ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addTodo()
		case 2:
			listTodo()
		case 3:
			markAsDone()
		case 4:
			fmt.Println("Thanks for using Todo App")
			return
		default:
			fmt.Println("Invalid choice")
		}

	}
}
