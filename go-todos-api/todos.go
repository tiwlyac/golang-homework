package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status string `json:"status"`
}

var todos = []Todo{}

func postTodoHandler(c *gin.Context) {
	t := Todo{}
	err := c.ShouldBindJSON(&t)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		t.ID = len(todos) + 1
		todos = append(todos, t)
		c.JSON(http.StatusCreated, t)
	}
}

func main() {
	r := gin.Default()
	r.POST("/api/todos", postTodoHandler)
	r.Run(":1234")
}