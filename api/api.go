package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	common "day.happy365/gotodo.common"
	repository "day.happy365/gotodo.repository"
	service "day.happy365/gotodo.service"
	"github.com/gin-gonic/gin"
)

func pong(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, common.Success("pong"))
}

func addTodo(c *gin.Context) {
	var todo repository.Todo
	if err := c.BindJSON(&todo); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, common.Fail[any]("请求失败"))
		return
	}

	todoId, err := service.AddTodo(todo)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, service.ErrInvalidInput) {
			status = http.StatusBadRequest
		}
		c.IndentedJSON(status, common.Fail[any](err.Error()))
		return
	}
	c.IndentedJSON(http.StatusOK, common.Success(todoId))
}

func deleteTodo(c *gin.Context) {
	todoIdParam := c.Param("todoId")
	todoId, err := strconv.ParseInt(todoIdParam, 10, 64)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, common.Fail[any]("参数错误"))
		return
	}

	affectedRows, err := service.DeleteTodo(todoId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, common.Fail[any](err.Error()))
		return
	}
	c.IndentedJSON(http.StatusOK, common.Success[any](affectedRows))
}

func updateTodo(c *gin.Context) {
	todoIdParam := c.Param("todoId")
	todoId, err := strconv.ParseInt(todoIdParam, 10, 64)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, common.Fail[any]("参数错误"))
		return
	}

	var todo repository.Todo
	if err := c.BindJSON(&todo); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, common.Fail[any]("请求失败"))
		return
	}

	todo.TodoId = int(todoId)

	rowsAffected, err := service.UpdateTodo(todo)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, common.Fail[any]("请求失败"))
		return
	}

	c.IndentedJSON(http.StatusOK, common.Success(rowsAffected))
}

func selectTodoCategory(c *gin.Context) {
	categories, err := service.SelectTodoCategory()
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, common.Fail[any]("请求失败"))
		return
	}

	c.IndentedJSON(http.StatusOK, common.Success(categories))
}

func selectTodoByTodoId(c *gin.Context) {
	todoIdParam := c.Param("todoId")
	todoId, err := strconv.ParseInt(todoIdParam, 10, 64)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, common.Fail[any]("参数错误"))
		return
	}

	todo, err := service.SelectTodoByTodoId(todoId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, common.Fail[any](err.Error()))
		return
	}
	c.IndentedJSON(http.StatusOK, common.Success(todo))
}

func selectTodos(c *gin.Context) {
	category := c.Query("category")
	content := c.Query("content")
	isComplete := c.Query("isComplete")

	if category == "" && content == "" && isComplete == "" {
		c.IndentedJSON(http.StatusBadRequest, common.Fail[any]("请求参数不能为空"))
		return
	}

	var todos []repository.Todo
	var err error

	if category != "" {
		todos, err = service.SelectTodosByCategory(category)
	} else if content != "" {
		todos, err = service.SelectTodosByContent(content)
	} else if isComplete != "" {
		todos, err = service.SelectTodosByIsComplete(isComplete)
	}

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, common.Fail[any](err.Error()))
		return
	}

	c.IndentedJSON(http.StatusOK, common.Success(todos))
}

func main() {
	engine := gin.Default()

	// todo api
	engine.GET("/ping", pong)
	engine.POST("/todos", addTodo)
	engine.DELETE("/todos/:todoId", deleteTodo)
	engine.PUT("/todos/:todoId", updateTodo)
	engine.GET("/todos/:todoId", selectTodoByTodoId)
	engine.GET("/todos", selectTodos)

	// category api
	engine.GET("/todos/categories", selectTodoCategory)

	addr := "127.0.0.1:8000"
	engine.Run(addr)

	fmt.Printf("server listening at %s\n", addr)
}
