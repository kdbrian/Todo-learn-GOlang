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
