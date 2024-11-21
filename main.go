package main

import (
	"fmt"
	"todoapp/todo"
)

func main() {

	//close db before script close
	defer func ()  {
		todo.CloseDb()
	}()


	todo.CreateTodoTable()
	todo.InsertSingleTodo(
		& todo.Todo{
		Id: 1,
		Title: "Hello todo 1",
	 	Message: "Hello I was created. first",
		IsDone: false,
	})

	todo.InsertSingleTodo(
		& todo.Todo{
		Id: 2,
		Title: "Hello todo 2",
	 	Message: "Hello I was created. second",
		IsDone: false,
	})

	//fetching all todos
	var todos []todo.Todo
	todo.FetchTodos(&todos)
	fmt.Println("ALL : ",todos)

	//fetching done todos
	var doneTodos []todo.Todo
	todo.FetchDoneTodos(&doneTodos)
	fmt.Println("DONE : ",doneTodos)
	
	//updating todos
	forUpdates := &todo.Todo{Title: "Hello todo 1", IsDone: true}
	todo.UpdateTodoStatus(forUpdates)
	todo.UpdateTodoMessage(forUpdates, "Fuck lorem ispum.")
	todo.UpdateTodoTitle(forUpdates, "I am for lorem ipsum.")

	//more updates
	var updatedTodos []todo.Todo
	todo.FetchDoneTodos(&updatedTodos)
	fmt.Println("UPDATED : ",updatedTodos)
	fmt.Println(todo.CheckTodoExists(&todo.Todo{Title: "Hello todo 1"}))

	todo.InsertSingleTodo(&todo.Todo{
		Title: "I will be deleted.", 
		Message: "But why",
		IsDone: true,
	})
	todo.DeleteTodoByTitle("I will be deleted.")
	
	var afterDelete []todo.Todo
	todo.FetchTodos(&afterDelete)
	fmt.Println("AFTER DELETING : ",updatedTodos)

	//clearing db
	todo.ClearDoneTodos()
	var notDoneTodos []todo.Todo
	todo.FetchTodos(&notDoneTodos)
	fmt.Println("NOT DONE : ",notDoneTodos)

	// todo.DeleteAllTodos()
}

