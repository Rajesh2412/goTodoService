package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

//to send the api response to the front end
// we need to make the json object and add data to that object
// and send that json data to the frontend
// to do  so we need to create a struct and add tags which makes the
//regular structs keys to work as json structs

type TodoList struct {
	ID              int8
	TodoName        string
	TodoDescription string
}
type TodoItem struct {
	TodoName        string
	TodoDescription string
}

func (app *App) GetTodos(w http.ResponseWriter, r *http.Request) {

	data, err := app.Models.TodoList.GetAllTodos()
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
	}

	payload := jsonPayload{
		Error:   false,
		Message: "Success",
		Data:    data,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *App) AddTodos(w http.ResponseWriter, r *http.Request) {

	newTodo := TodoItem{}
	err := app.readJSON(w, r, &newTodo)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
	}

	postTodoStatement := `INSERT INTO todos_list("todoName", "todoDescription") VALUES ($1, $2)`

	_, err = app.DBInstance.Exec(postTodoStatement, newTodo.TodoName, newTodo.TodoDescription)

	if err != nil {
		panic(err)
	}

	output := jsonPayload{
		Error:   false,
		Message: "Added to the TODO List",
	}

	app.writeJSON(w, http.StatusAccepted, output)

}

func (app *App) UpdateTodo(w http.ResponseWriter, r *http.Request) {

	todoID := chi.URLParam(r, "todoID")

	// decoding the json object from the request body
	todoItem := TodoItem{}
	err := json.NewDecoder(r.Body).Decode(&todoItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updateTodoStatement := `UPDATE public.todos_list SET  "todoName"=$2, "todoDescription"= $3 WHERE id=$1;`

	_, err = app.DBInstance.Exec(updateTodoStatement, todoID, todoItem.TodoName, todoItem.TodoDescription)

	if err != nil {
		panic(err)
	}

	output := jsonPayload{
		Error:   false,
		Message: "updated",
	}
	out, _ := json.Marshal(output)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(out)

}

func (app *App) DeleteTodo(w http.ResponseWriter, r *http.Request) {

	todoID := chi.URLParam(r, "todoID")

	deleteTodoStatement := `DELETE FROM public.todos_list WHERE id=$1;`

	_, err := app.DBInstance.Exec(deleteTodoStatement, todoID)
	if err != nil {
		panic(err)
	}

	output := jsonPayload{
		Error:   false,
		Message: "Deleted",
	}
	out, _ := json.Marshal(output)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(out)

}
