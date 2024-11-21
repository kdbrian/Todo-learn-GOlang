package main

import "fmt"

type Todo struct {
	Title, Message string
	IsDone         bool
}

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


func printTodo(todo *Todo){
	fmt.Println("-------------------------")
	fmt.Println("Title:", todo.Title)
	fmt.Println("Message:", todo.Message)
	fmt.Println("isDone:", todo.IsDone)
	fmt.Println("-------------------------")
}

func updateTitle(todo *Todo, newTitle string){
	todo.Title = newTitle
}
func updateMessage(todo *Todo, message string){
	todo.Message = message
}

func updateTodoStatus(todo *Todo){
	todo.IsDone = !todo.IsDone
}