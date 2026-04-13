package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"todo-app/db"
)

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "invalid json",
		})
		return
	}

	if strings.TrimSpace(task.ID) == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "Не указан идентификатор",
		})
		return
	}

	if _, err := strconv.Atoi(task.ID); err != nil {
		writeJSON(w, http.StatusNotFound, map[string]string{
			"error": "Задача не найдена",
		})
		return
	}

	if _, err := db.GetTask(task.ID); err != nil {
		writeJSON(w, http.StatusNotFound, map[string]string{
			"error": "Задача не найдена",
		})
		return
	}

	if strings.TrimSpace(task.Title) == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "title is required",
		})
		return
	}

	if err := checkDate(&task); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	if err := db.UpdateTask(&task); err != nil {
		writeJSON(w, http.StatusNotFound, map[string]string{
			"error": "Задача не найдена",
		})
		return
	}

	writeJSONOK(w, map[string]any{})
}
