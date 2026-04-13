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
			writeJSON(w, http.StatusBadRequest, map[string]string{
				"error": "Не указан идентификатор",
			})
			return
		}

		if err := db.DeleteTask(id); err != nil {
			writeJSON(w, http.StatusNotFound, map[string]string{
				"error": "Задача не найдена",
			})
			return
		}

		writeJSONOK(w, map[string]any{})
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
