package main

import (
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

	todoId := service.AddTodo(todo)
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

	affectedRows := service.DeleteTodo(todoId)
	c.IndentedJSON(http.StatusOK, common.Success[any](affectedRows))
}

func selectTodoByTodoId(c *gin.Context) {
	todoIdParam := c.Param("todoId")
	todoId, err := strconv.ParseInt(todoIdParam, 10, 64)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, common.Fail[any]("参数错误"))
		return
	}

	todo := service.SelectTodoByTodoId(todoId)
	c.IndentedJSON(http.StatusOK, common.Success(todo))
}

func selectTodosByCategoryOrContentOrIsComplete(c *gin.Context) {
	category := c.Query("category")
	content := c.Query("content")
	isComplete := c.Query("isComplete")
	log.Printf("api#selectTodosByCategoryOrContentOrIsComplete, category: %s, content: %s, isComplete: %s\n", category, content, isComplete)
	if category == "" && content == "" && isComplete == "" {
		log.Println("api#selectTodosByCategoryOrContentOrIsComplete: 请求参数为空")
		c.IndentedJSON(http.StatusBadRequest, common.Fail[any]("请求参数不能为空"))
		return
	}

	var todos []repository.Todo
	if category != "" {
		log.Println("api#selectTodosByCategoryOrContentOrIsComplete: 根据 category 查询")
		todos = service.SelectTodosByCategory(category)
	} else if content != "" {
		log.Println("api#selectTodosByCategoryOrContentOrIsComplete: 根据 content 查询")
		todos = service.SelectTodosByContent(content)
	} else if isComplete != "" {
		log.Println("api#selectTodosByCategoryOrContentOrIsComplete: 根据 isComplete 查询")
		todos = service.SelectTodosByIsComplete(isComplete)
	}

	c.IndentedJSON(http.StatusOK, common.Success(todos))
}

func main() {
	engine := gin.Default()

	engine.GET("/ping", pong)
	engine.POST("/todos", addTodo)
	engine.DELETE("/todos", deleteTodo)
	engine.GET("/todos/:todoId", selectTodoByTodoId)
	engine.GET("/todos", selectTodosByCategoryOrContentOrIsComplete)

	addr := "127.0.0.1:8000"
	engine.Run(addr)

	fmt.Printf("server listening at %s\n", addr)
}
