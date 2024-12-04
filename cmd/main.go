package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	data "github.com/Rajesh2412/go-todo-app/models"
)

type App struct {
	DBInstance *sql.DB
	Models     data.Models
}

func main() {

	port := os.Getenv("PORT")
	if port != "" {
		port = "8080"
	}
	log.Printf("Starting the Todo app on port %s", port)

	app := App{}

	//try connecitng to the DB
	connection := app.connectDB()

	app = App{
		DBInstance: connection,
		Models:     data.New(connection),
	}

	if connection == nil {
		log.Panic("can't connect to postgres")
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: app.Routes(),
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Panic("unable to start the server", err)
	}

}
