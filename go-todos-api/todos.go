package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
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

func getTodoByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for i, elem := range todos {
		if id == elem.ID {
			c.JSON(http.StatusOK, todos[i])
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Could not find this todo",
	})
}

func getTodosHandler(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func putTodoByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	t := Todo{}
	err = c.ShouldBindJSON(&t)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for i, elem := range todos {
		if id == elem.ID {
			todos[i].Status = t.Status 
			todos[i].Title = t.Title
			c.JSON(http.StatusOK, todos[i])
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Could not find this todo",
	})
}

func main() {
	r := gin.Default()
	r.POST("/api/todos", postTodoHandler)
	r.GET("/api/todos/:id", getTodoByIDHandler)
	r.GET("/api/todos", getTodosHandler)
	r.PUT("/api/todos/:id", putTodoByIDHandler)
	r.Run(":1234")
}