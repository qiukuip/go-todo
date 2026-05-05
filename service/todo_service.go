package service

import (
	"errors"
	"log"

	repository "day.happy365/gotodo.repository"
)

var ErrInvalidInput = errors.New("invalid input")
var ErrTodoNotFound = errors.New("todo not found")

func AddTodo(todo repository.Todo) (int64, error) {
	log.SetPrefix("service#AddTodo: ")

	if todo.Content == "" {
		return -1, ErrInvalidInput
	}

	todoId, err := repository.AddTodo(todo)
	if err != nil {
		return -1, err
	}

	return todoId, nil
}

func DeleteTodo(todoId int64) (int64, error) {
	rowsAffected, err := repository.DeleteTodo(todoId)
	if err != nil {
		return -1, err
	}

	return rowsAffected, nil
}

func UpdateTodo(todo repository.Todo) (int64, error) {
	if todo.TodoId <= 0 {
		return -1, ErrInvalidInput
	}

	rowsAffected, err := repository.UpdateTodo(todo)
	if err != nil {
		return -1, err
	}

	return rowsAffected, nil
}

func SelectTodoCategory() ([]string, error) {
	return repository.SelectTodoCategory()
}

func SelectTodoByTodoId(todoId int64) (repository.Todo, error) {
	return repository.SelectTodoByTodoId(todoId)
}

func SelectTodosByCategory(category string) ([]repository.Todo, error) {
	return repository.SelectTodosByCategory(category)
}

func SelectTodosByIsComplete(isComplete string) ([]repository.Todo, error) {
	return repository.SelectTodosByIsComplete(isComplete)
}

func SelectTodosByContent(content string) ([]repository.Todo, error) {
	return repository.SelectTodosByContent(content)
}
