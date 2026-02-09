package api

import "net/http"

func Init() {
	http.HandleFunc("/api/task", TaskHandler)
	http.HandleFunc("/api/task/done", TaskDoneHandler)
	http.HandleFunc("/api/tasks", TasksHandler)
	http.HandleFunc("/api/nextdate", NextDateHandler)
}
