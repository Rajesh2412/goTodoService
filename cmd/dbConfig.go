package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234567890"
	dbname   = "todos"
	sslmode  = "disable"
)

var count int64

func (app *App) connectDB() *sql.DB {
	// conString := "postgres://postgres:1234567890@localhost:5432/todo?sslmode=disable"
	//connString := os.Env("DSN")
	connString := "user='koyeb-adm' password=YGkD3V6OaXqm host=ep-noisy-snow-a281berd.eu-central-1.pg.koyeb.app dbname='todos'"
	// connString := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

	for {
		db, err := OpenDB(connString)
		if err != nil {
			log.Println("Postgres not yet ready...")
			count++
		} else {
			fmt.Println("Connected to Postgres")
			return db
		}

		if count > 10 {
			log.Println(err)
			return nil

		}

		log.Println("backing off two seconds")
		time.Sleep(2 * time.Second)
		continue

	}
}

func OpenDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		fmt.Println("Failed to connect to DB")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db, nil
}
