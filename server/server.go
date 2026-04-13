package server

import (
	"log"
	"net/http"
	"os"

	"todo-app/api"
)

var (
	webDir = "web"
	port   = getServPort()
)

func getServPort() string {
	if port := os.Getenv("TODO_PORT"); port != "" {
		return port
	}
	return "7540"
}

func Start() {
	if _, err := os.Stat(webDir); os.IsNotExist(err) {
		log.Fatal(err)
	}

	routSet()
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func routSet() {
	http.Handle("/", http.FileServer(http.Dir(webDir)))
	http.HandleFunc("/api/task", api.TaskHandler)
	http.HandleFunc("/api/task/done", api.TaskDoneHandler)
	http.HandleFunc("/api/nextdate", api.NextDateHandler)
	http.HandleFunc("/api/tasks", api.TasksHandler)
}
