package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

func OpenDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "todo.db")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database opened successfully!")
	return db
}

func CreateTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		task TEXT NOT NULL,
		done BOOLEAN DEFAULT FALSE
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table created successfully!")
}

func InsertTask(db *sql.DB, task string) {
	query := `INSERT INTO todos (task, done) VALUES (?, ?)`
	_, err := db.Exec(query, task, false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Task added successfully!")
}

func ListTasks(db *sql.DB) {
	rows, err := db.Query("SELECT id, task, done FROM todos")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Todo List:")
	for rows.Next() {
		var id int
		var task string
		var done bool
		err := rows.Scan(&id, &task, &done)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("[%d] %s - %t\n", id, task, done)
	}
}

func MarkTaskDone(db *sql.DB, id int) {
	query := `UPDATE todos SET done = TRUE WHERE id = ?`
	_, err := db.Exec(query, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Task marked as done!")
}
