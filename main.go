package main

import (
	"database/sql"
	"fmt"
	"log"
	"todoapp/todo"

	//importing postgres for side effects
	_ "github.com/lib/pq"
)

var db *sql.DB

//make sure dbname exists
const connectionString = "user=postgres password=2022 host=127.0.0.1 port=5432 dbname=learningtodoapp sslmode=disable"

//the init function is always called when the script is initialized
func init(){
	connectToDb()
}

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

func connectToDb(){
	var err error
	//connect to db
	db,err = sql.Open("postgres", connectionString)
	
	if err != nil{
		panic(err)
	}

	//check if connection established
	if err = db.Ping(); err != nil {
		log.Fatal("Failed to connect to db:", err)
	}else {
		fmt.Println("Connected to db.")
	}
}
