package todo

import "fmt"

type Todo struct {
	Title, Message string
	IsDone         bool
}

func PrintTodo(todo *Todo){
	fmt.Println("-------------------------")
	fmt.Println("Title:", todo.Title)
	fmt.Println("Message:", todo.Message)
	fmt.Println("isDone:", todo.IsDone)
	fmt.Println("-------------------------")
}

func UpdateTitle(todo *Todo, newTitle string){
	todo.Title = newTitle
}
func UpdateMessage(todo *Todo, message string){
	todo.Message = message
}

func UpdateTodoStatus(todo *Todo){
	todo.IsDone = !todo.IsDone
}