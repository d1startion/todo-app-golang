package main

import (
	"log"
	"os"

	"todo-app/db"
	"todo-app/server"
)

func main() {
	dbFile := "scheduler.db"
	if envDBFile := os.Getenv("TODO_DBFILE"); envDBFile != "" {
		dbFile = envDBFile
	}

	err := db.Init(dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	server.Start()

}
