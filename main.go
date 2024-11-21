package main

import "fmt"

type Todo struct {
	Title, Message string
	IsDone         bool
}

func main() {

	todo := Todo{}
	todo.Title = "Hello world"
	todo.Message = "Hello world this is a message."
	todo.IsDone = false
	fmt.Println(todo)

	todoPtr := &todo
	//updating status
	updateTodoStatus(todoPtr)
	fmt.Println(todo)

	updateTitle(todoPtr, "Some change title")
	fmt.Println(todo)

	updateMessage(todoPtr,"I changed the title.")
	fmt.Println(todo)

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