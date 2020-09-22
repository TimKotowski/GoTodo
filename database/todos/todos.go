package todos

import (
	"database/sql"
	"fmt"
	"time"
)

// Database struct defines the todos database.
type Database struct {
	db *sql.DB
}

// New returns a new Todos struct for the database, which will
// have methods attached to it so we can create, read, update,
// and delete todos from the Postgres database.
func New(db *sql.DB) *Database {
	return &Database{
		db: db,
	}
}

// Todo defines a todo.
type Todo struct {
	ID          int       `json:"id"`
	Created     time.Time `json:"created"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
}

// define a set of Todos

type Todos struct {
	Todos []*GetParams `json:"todos"`
}

const (
	// stmGetAll deines the SQL statment to
	// get all todos from the database
	stmGetAllTodos = `SELECT id, description, completed FROM todos`
	// stmtInsert defines the SQL statement to
	// get a specific todo
	stmGetOneTodo = `SELECT id, description, completed WHERE id = $1`
	// stmtInsert defines the SQL statement to
	// insert a new todo into the database.
	stmtInsert = `INSERT INTO todos (created, description, completed) VALUES ($1, $2, $3) RETURNING id`
	// stmtInsert defines the SQL statement to
	// update a  todo from the database.
	stmtUpdate = `UPDATE todos SET description=$2 WHERE id=$1`
	// stmtInsert defines the SQL statement to
	// delete a  todo from the database.
	stmtDeleted = `DELETE FROM todos WHERE ID=$1`
	// stmtInsert defines the SQL statement to
	// delete a  todo from the database.
	stmtDeleteTodoStatus = `IF (SELECT completed="true" FROM todos WHERE id=$1) THEN
		DELETE FROM todos WHERE id=$1
		ELSE
		UPDATE todo SET completed="true" WHERE id=$1
		END IF`
	// stmtInsert defines the SQL statement to
	// update completed status to true in database
	stmCompletedStatus = `UPDATE todos SET completed="true" where id=$1`
)

// NewParams defines the parameters for the New method.
type NewParams struct {
	Description string `json:"description"`
}


// New creates a new todo.
func (db *Database) New(params *NewParams) (*Todo, error) {
	// Create a new Todo.
	todo := &Todo{
		Created:     time.Now(),
		Description: params.Description,
	}
	// put the the created todo in the database
	// the todo saraible will hold the json request body info
	// get the info from the todo and insert it in the the Exec
	id := 0
	err := db.db.QueryRow(stmtInsert, todo.Created, todo.Description, todo.Completed).Scan(&id)
	if err != nil {
		fmt.Println("inside db.Exec error")
		fmt.Printf("error is: %v\n", err.Error())
		return nil, err
	}

	fmt.Printf("called exec with: %v %v %v\n", todo.Created, todo.Description, todo.Completed)

	todo.ID = id
	return todo, nil
}

// GetParams defines the parameters for the Get method.
type GetParams struct {
	ID        int       `json:"id"`
	Description string    `json:"description"`
	Completed bool      `json:"completed"`
}
func (db *Database) GetAll() (*Todos, error) {

		todos := &Todos{
			Todos: []*GetParams{},
		}

		rows, err := db.db.Query(stmGetAllTodos)
		if err != nil {
			fmt.Println("sinide db.Query error")
			fmt.Printf("eorr is in %v\n", err.Error())
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			todo := &GetParams{}

			err = rows.Scan(&todo.ID,  &todo.Description, &todo.Completed); if err != nil {
				fmt.Println("sinide db.Query error")
				fmt.Printf("eorr is in %v\n", err.Error())
				return nil, err
			}

			todos.Todos = append(todos.Todos, todo)
		}
		return todos, nil
}

// type Parameter struct {
// 	ID int `json:"id"`
// }

func (db *Database) GetSpecificTodo(params *int) (*GetParams, error) {
		todo := &GetParams{}

			var id int
			var descripton string
			err := db.db.QueryRow(stmGetOneTodo, params).Scan(&id, &descripton)
			switch err {
			case sql.ErrNoRows:
				fmt.Println("No rows were returned!")
			case nil:
				fmt.Println(id, descripton)
			default:
				fmt.Println("error")
			}

			return todo, nil
}
