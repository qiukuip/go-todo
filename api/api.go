package main

import (
	"fmt"
	"log"
	"net/http"

	common "day.happy365/gotodo.common"
	repository "day.happy365/gotodo.repository"
	service "day.happy365/gotodo.service"
	"github.com/gin-gonic/gin"
)

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
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

func main() {
	engine := gin.Default()

	engine.GET("/ping", pong)
	engine.POST("/todo", addTodo)

	addr := "127.0.0.1:8000"
	engine.Run(addr)

	fmt.Printf("server listening at %s\n", addr)
}
