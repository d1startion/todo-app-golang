package api

import (
	"net/http"

	"todo-app/db"
)

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		AddTaskHandler(w, r)
	case http.MethodGet:
		getTaskHandler(w, r)
	case http.MethodPut:
		updateTaskHandler(w, r)
	case http.MethodDelete:
		id := r.URL.Query().Get("id")
		if id == "" {
			writeJSON(w, map[string]string{"error": "Не указан идентификатор"})
			return
		}
		err := db.DeleteTask(id)
		if err != nil {
			writeJSON(w, map[string]string{"error": "Задача не найдена"})
			return
		}
		writeJSON(w, map[string]any{})
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
