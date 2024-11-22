# <code>Todo { Title : <span style="color: rgb(0, 172, 215);">"Learning golang"</span>, Message :<span style="color: rgb(126, 79, 38)">"Learning golang by building a full stack todo app"</span> }</code>

## All about coding for fun
This code is divided by branches. starting from master where I use plain variables to save the values upto complex structs that are read into the relational database(postgress)

## Branches
- ### Master
    contains all the plain code, saving a single todo using a straight forward approach of having the fields declared separately
    ```go
        todo := "Go home"
	    isDone := false
    ```
    these are updated and printed by direct mutation.
- ### Structs
    I introduce the concepts of using structs to simplify how related data is sparse over. With this I push all <em>Todo</em> related code to the <code>/todo</code> package for easier maintainance.

    The todo package introduces a new <code>interface</code> which contains some functionalities listed below :- 
    ```go
        type TodoBehavior interface{
            PrintTodo()
            UpdateTitle(newTitle string)
            UpdateMessage(message string)
            UpdateTodoStatus(status bool)
        }
    ```
    As of Go any type can implement this interface and have access to the methods so as to achieve the <code>abstraction</code> feature.
    So I had to write all the code to implement the methods in the TodoBehavior type.

    For simplicty and to avoid a lot of copying by the compiler, I passed a reference of todo to the methods as a pointer.
    ```go
        func (todo *Todo)PrintTodo()
    ```
    With this i still get access to all methods while just pointing to the same adress.
- ### Postgres
    Postgres provides an API over golang that can be used to access sql databases, with raw sql and prepared statements.
    - starts by creating a database connection specifying all the required parameters
    ```go
        //params : drivername, connectionString
        db, err := sql.Open("postgres", "user=postgres password=2022 host=127.0.0.1 port=5432 dbname=learningtodoapp sslmode=disable")
    ```

    - Check connection is established.
    ```go
        if db.Ping() != nil{
		    panic("Error contacting db")
	    }
    ```
    - write down the raw sql and prepared statements. Raw sql contains raw sql and used mostly when performing <code>unfiltered selections, drop table</code> or any command without the where clause. <b>Prepared statements</b> are used mostly when <code>inserting, updating and deleting</code>
        - An example if <em>raw sql</em> is : 
            ```go
            const fetchAllTodos = `
            SELECT * FROM todos
            `
            ```
            <b>Note:</b> the use of string literals to make sure the string is passed as is.

        - Example of <em>Prepared statement</em> is : 
            ```go
                const insertSingleTodo=`
                    INSERT INTO todos VALUES ($1, $2, $3, $4)
                `
            ```

            <code>Prepare</code> prepares the statement by validating it. It returns an <strong>alias</strong> on which you can execute queries by passing the <code>positional arguments</code> into it.
            ```go
                
                //preparing 
                statement, err := db.Prepare(insertSingleTodo)

                //making sure everything worked fine
                if err != nil {
		            log.Fatal("Failed to prepare getByTitle : ", err)
	            }

                //executing
                //passing argument in same positional order
                res, err := statement.Exec(1, "Supa mario", "Watch supa mario at noon", false)
            ```

        - statement.<code>Exec</code> has a different return based on the prepared statement. For the most part a <code>sql.Rows</code> is returned giving more control access over the flow.
        
        ```go
            rows, err := res.RowsAffected()
            //if 0 no changes/unsuccessful exec 
            //otherwise changes/successful exec
        ```

    ## From Postgres 
    It got too tideous as I had to define all the functions to <code>CREATE, READ, UPDATE & DELETE</code> the todo entities. Also I has to ensure the arguments were <strong style="color:teal;">passed</strong> correctly and the code was becoming <strong style="color:goldenrod;">unmaintainable</strong>.
    <br>
    So moving to an <strong>ORM(Object Relational Mapper)</strong> was not a option or a choice. I had to do it😂.

- ### Gorm (Go Object Relational Mapper)
    - Gorm provides ORM functions over the go language and an underlying database, given a driver. So it is a good match.

    - Download dependencies
        ```bash
            #add dependencies
            $ go get gorm.io/gorm gorm.io/driver/postgres
        ```

    - Like postgres Gorm provides a easy API over the database. 

    - First I modify the <code>Todo struct</code> to make it a gorm entity.
        ```go
            type Todo struct {
                gorm.Model
                Title, Message string
                IsDone         bool
            }
        ```
    - I define the <em>gorm.Model</em> field which adds some esential fields that might be needed. these are:
        - ID (An Autoincremented self assigned <span style="color: rgb(0, 172, 215);">Primary Key</span>)
        - date_created (As it is said)
        - date_updated (As it is said)
        - date_deleted (As it is said)

        - Among these are some functions to efficiently update the values while interacting with the db.

    - By default the fields are <code>snake_cased</code> in the database column names.

    - Then create a db reference Just like in postgres.
        ```go
            
            var db *gorm.DB//global

            const connectionString = "user=postgres password=2022 host=127.0.0.1 port=5432 dbname=learningtodoapp sslmode=disable"

            db,err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
            //db is an instance of *gorm.DB

        ```

    - To create a table pass the reference to the struct into the <code>AutoMigrate</code> function. Which works by creating the table for us if doesnt already exist, and then ensure all new records contains all the fields in the schema; hence <span style="color:red;">AutoMigrate</span>.
        ```go
            func CreateTodoTable(){
                err := db.AutoMigrate(&Todo{})
                if err != nil {
                    log.Fatal("Error creating table : ", err)
                }else {
                    log.Println("Created todo db")
                }
            }
        ```

    - To insert a value into a table use the <code>db.Create</code> function passing the reference to the object being saved
        ```go
            func InsertSingleTodo(todo *Todo){
                db.Create(todo)
                log.Println("Inserted Todo")
            }
        ```
    
    - Gorm provides an easier interface over the underlying database and cuts off a lot of <em>boilerplate</em> code for doing database operations.


## Updates Loading 🚀
### correct this by opening a pull request; otherwise enjoy the fun of being a <span style="font-size:60px; color: rgb(0, 172, 215);"> GO <span style="color:rgb(195, 47, 38); ">Getter</span></span>

