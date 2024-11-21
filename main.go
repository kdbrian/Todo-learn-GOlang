package main

import (
	"fmt"
	"todoapp/todo"
)

func main() {

	//for databases
	// newTodo := todo.Todo{
	// 	Id: 1,
	// 	Title: "Hello todo",
	//  	Message: "Hello I was created.",
	// 	IsDone: false,
	// }

	//close db after script close
	defer func(){
		todo.GetDb().Close()
	}()

	// todo.CreateTodoTable()
	// todo.InsertSingleTodo(&newTodo)
	var todos []todo.Todo
	todo.FetchTodos(&todos)
	fmt.Println(todos)
	
	todo.DeleteAllTodos()
}

