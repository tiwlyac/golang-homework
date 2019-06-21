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
	checkError(c, err)
	t.ID = len(todos) + 1
	todos = append(todos, t)
	c.JSON(http.StatusCreated, t)
}

func getTodoByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	checkError(c, err)

	for i, elem := range todos {
		if id == elem.ID {
			c.JSON(http.StatusOK, todos[i])
			return
		}
	}

	c.JSON(http.StatusOK, "Could not find this todo.")
}

func getTodosHandler(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func putTodoByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	checkError(c, err)
	t := Todo{}
	err = c.ShouldBindJSON(&t)
	checkError(c, err)

	for i, elem := range todos {
		if id == elem.ID {
			todos[i].Status = t.Status 
			todos[i].Title = t.Title
			c.JSON(http.StatusOK, todos[i])
			return
		}
	}

	c.JSON(http.StatusOK, "Could not find this todo.")
}

func deleteTodoByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	checkError(c, err)

	for i, elem := range todos {
		if id == elem.ID {
			todos = remove(todos,i)
			c.JSON(http.StatusOK, gin.H{ "status": "success" })
			return
		}
	}

	c.JSON(http.StatusOK, "Could not find this todo.")
}

func checkError(c *gin.Context, err error) {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
}

func remove(slice []Todo, s int) []Todo {
    return append(slice[:s], slice[s+1:]...)
}

func main() {
	r := gin.Default()
	r.POST("/api/todos", postTodoHandler)
	r.GET("/api/todos/:id", getTodoByIDHandler)
	r.GET("/api/todos", getTodosHandler)
	r.PUT("/api/todos/:id", putTodoByIDHandler)
	r.DELETE("/api/todos/:id", deleteTodoByIDHandler)
	r.Run(":1234")
}