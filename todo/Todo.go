package todo

import (
	"fmt"

	"gorm.io/gorm"
)

func init(){
	fmt.Println("Initialized")
}

type Todo struct {
	gorm.Model
	Title, Message string
	IsDone         bool
}


type TodoBehavior interface{	
	//normal operations
	PrintTodo()
	UpdateTitle(newTitle string)
	UpdateMessage(message string)
	UpdateTodoStatus(status bool)
}

// implementing the TodoBehavior interface
func (todo *Todo)PrintTodo(){
	fmt.Println("-------------------------")
	fmt.Println("Title:", todo.Title)
	fmt.Println("Message:", todo.Message)
	fmt.Println("isDone:", todo.IsDone)
	fmt.Println("-------------------------")
}

func (todo *Todo) UpdateTitle(newTitle string){
	todo.Title = newTitle
}
func (todo *Todo) UpdateMessage(message string){
	todo.Message = message
}

func (todo *Todo) UpdateTodoStatus(status bool){
	todo.IsDone = status
}

