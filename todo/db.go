package todo

import (
	"fmt"
	"log"

	//importing postgres for side effects
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var db *gorm.DB

//make sure dbname exists
const connectionString = "user=postgres password=2022 host=127.0.0.1 port=5432 dbname=learningtodoapp sslmode=disable"

//the init function is always called when the script is initialized
func init(){
	connectToDb()
}

func connectToDb(){
	var err error
	//connect to db
	db,err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	
	if err != nil{
		panic(err)
	}
	fmt.Println("Connected to db.")
}

func getDb() *gorm.DB {
	return db
}

func CreateTodoTable(){
	err := db.AutoMigrate(&Todo{})
	if err != nil {
		log.Fatal("Error creating table : ", err)
	}else {
		log.Println("Created todo db")
	}
}

func CheckTodoExists(todo *Todo) (bool, Todo) {

	var found *Todo = nil
	db.First(found, "title = ?", todo.Title)
	
	if found == nil {
		log.Fatalf("Not found by %s", todo.Title)
		return false, nil
	}else {
		log.Printf("Found by %s", todo.Title)
		return true, *found
	}
	
}

func InsertSingleTodo(todo *Todo){
	db.Create(todo)
	log.Println("Inserted Todo")
}

func UpdateTodoStatus(todo *Todo){
	if exists, _ := CheckTodoExists(todo); !exists {
		log.Fatalf("No todo found. not updating")
	}else{
		db.Update("is_done", todo.IsDone)
		log.Println("Updated todo")
	}
}

func UpdateTodoTitle(todo *Todo, newTitle string){
	if exists, _ := CheckTodoExists(todo); !exists {
		log.Fatalf("No todo found. not updating")
	}else{
		db.Update("title", newTitle)
		log.Println("Updated todo")
	}
}

func UpdateTodoMessage(todo *Todo, newMessage string){
	if exists, _ := CheckTodoExists(todo); !exists {
		log.Fatalf("No todo found. not updating")
	}else{
		db.Update("message", newMessage)
		log.Println("Updated todo")
	}
}

func FetchTodos(todos *[]Todo){
	db.Find(todos)
}

func FetchDoneTodos(todos *[]Todo){
	db.Find(todos, "is_done = ?", true)
}

func DeleteTodoByTitle(title string)  {
	if isFound, todo := CheckTodoExists(&Todo{Title: title}); !isFound {
		log.Fatalf("Missing todo with title : %s\n",title)
	}else {
		db.Delete(todo, "title = ?", title)
		log.Println("Deleted todo")
	}
}

// func DeleteAllTodos(){
// 	db.Delete("")
// }

// func ClearDoneTodos()  {
	
// }

// func CloseDb()  {
// 	err := db. ()
// 	if err != nil {
// 		log.Fatal("Failed to close connection : ", err)
// 	}
// }