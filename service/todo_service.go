package service

import (
	"log"

	repository "day.happy365/gotodo.repository"
)

func AddTodo(todo repository.Todo) int64 {
	log.SetPrefix("service#AddTodo: ")

	todoId, err := repository.AddTodo(todo)
	if err != nil {
		log.Fatal(err)
		return -1
	}

	return todoId
}

func DeleteTodo(todoId int64) int64 {
	rowsAffected, err := repository.DeleteTodo(todoId)
	if err != nil {
		log.Fatal(err)
		return -1
	}

	return rowsAffected
}

func UpdateTodo(todo repository.Todo) int64 {
	rowsAffected, err := repository.UpdateTodo(todo)
	if err != nil {
		log.Fatal(err)
		return -1
	}

	return rowsAffected
}

func SelectTodosByCategory(category string) []repository.Todo {
	todos, err := repository.SelectTodosByCategory(category)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return todos
}

func SelectTodosByIsComplete(isComplete string) []repository.Todo {
	todos, err := repository.SelectTodosByCategory(isComplete)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return todos
}

func SelectTodosByContent(content string) []repository.Todo {
	todos, err := repository.SelectTodosByCategory(content)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return todos
}
