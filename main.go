package main

func main() {

	todo := Todo{}
	todoPtr := &todo

	//updating with pointers
	todoPtr.Title = "Hello world"
	todoPtr.Message = "Hello world this is a message."
	todoPtr.IsDone = false

	printTodo(todoPtr)

	//updating status
	updateTodoStatus(todoPtr)
	printTodo(todoPtr)

	updateTitle(todoPtr, "Some change title")
	printTodo(todoPtr)

	updateMessage(todoPtr,"I changed the title.")
	printTodo(todoPtr)

}

