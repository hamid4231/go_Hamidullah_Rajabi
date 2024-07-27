package services

import (
	"fmt"
	"practice/models"
)

var database []models.Todo = []models.Todo{}

func AddTodo(newTodo models.Todo) {
	database = append(database, newTodo)
}

func GetTodoList() []models.Todo {
	return database
}

// getTodoByID displays todo by given ID
func GetTodoByID(id uint) {
	for _, t := range database {
		if t.ID == id {
			fmt.Printf("ID: %d, Title: %s, Content: %s\n", t.ID, t.Title, t.IsDone)
			return
		}
	}
	fmt.Printf("Todo with ID %d not found\n", id)
}

// deleteTodoByID deletes todo by given ID
func DeleteTodoByID(id uint) {
	for index, t := range database {
		if t.ID == id {
			database = append(database[:index], database[index+1:]...)
			return
		}
	}
	fmt.Printf("Todo with ID %d not found\n", id)
}
