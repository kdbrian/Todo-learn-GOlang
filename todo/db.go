package todo

import (
	"database/sql"
	"fmt"
	"log"

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

func GetDb() *sql.DB {
	return db
}

/*
type TodoTbl interface{
	//database operations
	CreateTodoTable()
	InsertSingleTodo(*Todo)
	UpdateTodoStatus(*Todo)
	UpdateTodoTitle(*Todo,string)
	UpdateTodoMessage(*Todo,string)
	FetchTodos()[]Todo
	FetchDoneTodos()[]Todo
	DeleteAllTodos()
}
*/


const createTodoTable =`
CREATE TABLE IF NOT EXISTS public.todos (
	id integer,
	title varchar(100),
	message varchar(100),
 	is_done BOOLEAN NOT NULL DEFAULT FALSE
)

WITH(OIDS=FALSE)
`
const insertSingleTodo=`
INSERT INTO todos VALUES ($1, $2, $3, $4)
`
const getTodoByTitle=`
SELECT message FROM todos WHERE title = $1
`
const updateTodoStatus =`
UPDATE todos SET
is_done = $1
WHERE title = $2
`
const updateTodoTitle =`
UPDATE todos SET
title = $1
WHERE title = $2
`
const updateTodoMessage =`
UPDATE todos SET
message = $1
WHERE title = $2
`
const fetchAllTodos = `
SELECT * FROM todos
`

const fetchDoneTodos = `
SELECT * FROM todos WHERE is_done = TRUE
`

func CreateTodoTable(){
	// db := GetDb()
	if db.Ping() != nil{
		panic("Error contacting db")
	}

	_, err := db.Exec(createTodoTable)
	
	if err != nil {
		log.Fatal("Error creating table : ", err)
	}else {
		log.Println("Created todo db")
	}
}

func CheckTodoExists(todo *Todo) bool {
	statement, err := db.Prepare(getTodoByTitle)
	if err != nil {
		log.Fatal("Failed to prepare getByTitle : ", err)
		return false
	}

	res, err := statement.Exec(todo.Title)
	if err != nil {
		log.Fatalf("Failed to insert %s : %v", todo.Title,err)
		return false
	}

	if rows,err := res.RowsAffected(); err != nil{
		log.Fatalf("Failed to insert %s : %v", todo.Title,err)
		return false
	}else {
		if rows == 0 {
			log.Printf("Sorry not found %s.\n", todo.Title)
			return false
		}else {
			log.Printf("Got %s, %d affected",todo.Title, rows)
			return true
		}
	}
}

func InsertSingleTodo(todo *Todo){
	//prepare statement
	statement, err := db.Prepare(insertSingleTodo)
	if err!=nil {
		log.Fatal("Failed to prepare statement")
	}

	//execute
	res, err := statement.Exec(todo.Id, todo.Title, todo.Message, todo.IsDone)
	if err != nil {
		log.Fatal("Failed to insert record : ", err)
	}

	//check for result
	rows, err := res.RowsAffected()

	if err != nil {
		log.Fatal("Failed to insert record : ", err)
	}
	log.Printf("Inserted %d successfully.", rows)
}

func UpdateTodoStatus(todo *Todo){
	if exists := CheckTodoExists(todo); !exists {
		log.Fatalf("No todo found. not updating")
	}else{

		statement, err := db.Prepare(updateTodoStatus)
		if err != nil {
			log.Fatal("Failed to prepare update", err)
		}

		_, err = statement.Exec(todo.IsDone, todo.Title)
		if err != nil {
			log.Fatal("Failed to update", err)
		}

		log.Println("Updated todo")
	}
}

func UpdateTodoTitle(todo *Todo, newTitle string){
	if exists := CheckTodoExists(todo); !exists {
		log.Fatalf("No todo found. not updating")
	}else{

		statement, err := db.Prepare(updateTodoTitle)
		if err != nil {
			log.Fatal("Failed to prepare update", err)
		}

		_, err = statement.Exec(newTitle,todo.Title)
		if err != nil {
			log.Fatal("Failed to update", err)
		}

		log.Println("Updated todo")
	}
}

func UpdateTodoMessage(todo *Todo, newMessage string){
	if exists := CheckTodoExists(todo); !exists {
		log.Fatalf("No todo found. not updating")
	}else{

		statement, err := db.Prepare(updateTodoMessage)
		if err != nil {
			log.Fatal("Failed to prepare update", err)
		}

		_, err = statement.Exec(newMessage,todo.Title)
		if err != nil {
			log.Fatal("Failed to update", err)
		}

		log.Println("Updated todo")
	}
}

func FetchTodos(todos *[]Todo){
	// todos =nil; todos = &(make([]Todo,1))
	result, err := db.Prepare(fetchAllTodos)

	if err != nil {
		log.Fatal("Failed to load todos : ", err)
	}

	if rows,err := result.Query(); err != nil {
		log.Fatal("Failed to load todos : ", err)
	}else {
		for rows.Next() {
			var title, message string
			var isDone bool
			var id int
			err := rows.Scan(&id, &title,&message,&isDone)
			// fmt.Printf("Got (%d, %s, %s, , %v)", id, title, message, isDone)
			if err != nil {
				log.Fatal("Failed to offload row : ", err)
			}
			(*todos) = append((*todos), Todo{Id: id, Title: title, IsDone: isDone})
		}
	}
}

func FetchDoneTodos(todos *[]Todo){
	result, err := db.Prepare(fetchDoneTodos)

	if err != nil {
		log.Fatal("Failed to load todos : ", err)
	}

	if rows,err := result.Query(); err != nil {
		log.Fatal("Failed to load todos : ", err)
	}else {
		for rows.Next() {
			var title, message string
			var isDone bool
			var id int
			err := rows.Scan(&id, &title,&message,&isDone)
			if err != nil {
				log.Fatal("Failed to offload row : ", err)
			}
			(*todos) = append((*todos), Todo{Id: id, Title: title, Message: message, IsDone: isDone})
		}
	}
	
	// return todos
}

func DeleteAllTodos(){
	const trancuteTodoTbl = `
	TRUNCATE TABLE todos
	`
	_, err := db.Exec(trancuteTodoTbl)

	if err != nil {
		log.Fatal("Failed to drop table : ", err)
	}
}