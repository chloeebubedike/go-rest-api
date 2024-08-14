package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool `json:"completed"`
}


var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false },
	{ID: "2", Item: "Read Book", Completed: false },
	{ID: "3", Item: "Write Code", Completed: false },
}

// helper functions
func getTodos(context *gin.Context){
	context.IndentedJSON(http.StatusAccepted, todos)
}

func getTodosById(id string) (*todo, error){
	for i, t := range todos{
		if t.ID == id{
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

// post request
func addTodo(context *gin.Context){
	var newTodo todo 

	if err := context.BindJSON(&newTodo); err != nil {
		return 
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}


// patch request
func toggleTodoStatus(context *gin.Context){
	id := context.Param("id")
	todo, err := getTodosById(id)

	if err !=nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
	}

	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusAccepted, todo)

}

// get request
func getTodo(context *gin.Context){
	id := context.Param("id")

	todo, err := getTodosById(id)

	if err !=nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
	}

	context.IndentedJSON(http.StatusAccepted, todo)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todos", addTodo)
	router.Run("localhost:9090")
}