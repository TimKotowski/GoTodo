package api

import (
	"encoding/json"
	"fmt"
	"gotodo/database"
	"gotodo/database/todos"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

// New exports a new API.
func New(db *database.Database, r *chi.Mux) {
	r.Get("/api/v1/todos", HandleGetAllTodos(db))
	r.Get("/api/v1/todos/single/{id}", HandleGetTodo(db))
	r.Post("/api/v1/todos", HandleCreate(db))
	r.Post("/api/v1/todos/update/description/{id}", HandleUpdate(db))
	// r.Post("/api/v1/update/completed/:id", HandleCompletionStatus(db))
}

func HandleCreate(db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		var params todos.NewParams
		if err := json.NewDecoder(r.Body).Decode((&params)); err != nil {
			log.Println(w, err, http.StatusBadRequest)
			return
		}

		fmt.Printf("params are %+v\n", params)

		todo, err := db.Todos.New(&params)
		if err != nil {
			fmt.Printf("err on .New is: %v\n", err.Error())
		}
		fmt.Printf("todo is: %+v\n", todo)
		jsonData, err := json.Marshal(todo)
		fmt.Println("jsonData", jsonData)

		if err != nil {
			log.Println(err)
			w.Write([]byte("We got an error"))
			return
		}
		fmt.Println("json", string(jsonData))
		w.Write(jsonData)
	}
}

func HandleGetAllTodos(db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		todo, err := db.Todos.GetAll()
		if err != nil {
			fmt.Printf("err on .New is: %v\n", err.Error())
		}
		fmt.Printf("todo is: %+v\n", todo)
		jsonData, err := json.Marshal(todo)

		if err != nil {
			log.Println(err)
			w.Write([]byte("We got an error"))
			return
		}
		fmt.Println("json", string(jsonData))
		w.Write(jsonData)
	}
}

func HandleGetTodo(db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		params := chi.URLParam(r, "id")
		id, _ := strconv.Atoi(params)

		todo, err := db.Todos.GetSpecificTodo(id)
		if err != nil {
			fmt.Printf("err on .New is: %v\n", err.Error())
		}
		fmt.Printf("todo is: %+v\n", todo)

		jsonData, err := json.Marshal(todo)
		fmt.Println("jsonData", jsonData)

		if err != nil {
			log.Println(err)
			w.Write([]byte("We got an error"))
			return
		}
		fmt.Println("json", string(jsonData))
		w.Write(jsonData)
	}
}


func HandleUpdate(db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		var params todos.UpdateParams
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			log.Println(w, err, http.StatusBadRequest)
			return
		}
		fmt.Printf("params are %+v\n", params)

		ID := chi.URLParam(r, "id")
		if ID == ""  {
			log.Println(w, http.StatusInternalServerError)
			log.Println("error in id")
			return
		}

		id, _ := strconv.Atoi(ID)

		todo, err := db.Todos.UpdateTodo(id, &params)
		if err != nil {
			fmt.Printf("err on .New is: %v\n", err.Error())
			return
		}
		fmt.Printf("done")

		jsonData, err := json.Marshal(todo)
		fmt.Println("jsonData", jsonData)

		if err != nil {
			log.Println(err)
			w.Write([]byte("We got an error"))
			return
		}
		fmt.Println("json", string(jsonData))
		w.Write([]byte(jsonData))
	}

}
