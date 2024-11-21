package main

import "todoapp/todo"

func main() {

	todo := todo.Todo{}
	todoPtr := &todo

	//updating with pointers
	todoPtr.Title = "Hello world"
	todoPtr.Message = "Hello world this is a message."
	todoPtr.IsDone = false

	//methods not visible
	todoPtr.PrintTodo()

	//updating status
	todoPtr.UpdateTodoStatus(true)
	todoPtr.PrintTodo()

	todoPtr.UpdateTitle("Some change title")
	todoPtr.PrintTodo()

	todoPtr.UpdateMessage("I changed the title.")
	todoPtr.PrintTodo()

}

