package data

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type Models struct {
	TodoList TodoList
}

var db *sql.DB

const dbTimeout = time.Second * 3

type TodoList struct {
	ID              int8   `json:"id"`
	TodoName        string `json:"todoName"`
	TodoDescription string `json:"todoDescription"`
}

func New(connString *sql.DB) Models {
	db = connString
	return Models{
		TodoList: TodoList{},
	}
}

func (todoList *TodoList) GetAllTodos() ([]*TodoList, error) {

	ctxt, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `SELECT * FROM todos_list`

	result, err := db.QueryContext(ctxt, query)
	if err != nil {
		return nil, err
	}

	defer result.Close()

	var todoItems []*TodoList

	for result.Next() {
		var todoItem TodoList
		err := result.Scan(
			&todoItem.ID,
			&todoItem.TodoName,
			&todoItem.TodoDescription,
		)
		if err != nil {
			log.Println("Error Scanning", err)
			return nil, err
		}

		todoItems = append(todoItems, &todoItem)
	}

	return todoItems, nil

}
