package api

import (
	"net/http"

	"todo-app/db"
)

const tasksLimit = 50

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func TasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.Tasks(tasksLimit)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	if tasks == nil {
		tasks = []*db.Task{}
	}

	writeJSONOK(w, TasksResp{Tasks: tasks})
}
