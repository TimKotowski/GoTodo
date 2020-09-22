package main

import (
	"database/sql"
	"fmt"
	"gotodo/api"
	"net/http"
	"time"

	"gotodo/database"

	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "timkotowski"
	password = "butter333"
	dbname   = "gotodo"
)

func main() {
	r := chi.NewRouter()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// Create a new gotodo database.
	gdb := database.New(db)

	// Create a new API using our router.
	api.New(gdb, r)

	// Create a new HTTP server.
	server := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("Running server...")

	// Start the HTTP server.
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}
